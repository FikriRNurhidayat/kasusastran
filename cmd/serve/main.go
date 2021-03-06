package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	nsq "github.com/fikrirnurhidayat/kasusastran/app/driver/nsq"
	pg "github.com/fikrirnurhidayat/kasusastran/app/driver/postgres"
	rds "github.com/fikrirnurhidayat/kasusastran/app/driver/redis"
	"github.com/fikrirnurhidayat/kasusastran/proto"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/fikrirnurhidayat/kasusastran/app/config"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/redis"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/encoding/protojson"
)

// Drivers
var tcp net.Listener
var server *grpc.Server
var mux *runtime.ServeMux
var opts []grpc.DialOption
var db query.Querier

// Services
var authenticationServer *srv.AuthenticationServer
var seratsServer *srv.SeratsServer
var wulangansServer *srv.WulangansServer

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	log                = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
)

func init() {
	// Initialize driver
	dbURL := config.GetDatabaseConnectionString()
	dbConn, err := pg.Connect(dbURL)

	if err != nil {
		log.Fatalf("pg.Connect: cannot connect to database: %v", err)
	}

	tcp, err = net.Listen("tcp", *grpcServerEndpoint)

	if err != nil {
		log.Fatalf("net.Listen: cannot initialize to tcp connection: %v", err)
	}

	rc, err := rds.ConnectRedis(config.GetRedisConnectionString())
	if err != nil {
		log.Fatalf("rds.ConnectRedis: cannot initialize redis connection: %v", err)
	}

	mux = runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		runtime.WithErrorHandler(srv.CustomErrorHandler),
	)
	opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	server = grpc.NewServer()
	db = query.New(dbConn)

	nsqc := nsq.NewConnection(config.GetNSQLookupAddress(), config.GetNSQAddress())
	producer, err := nsqc.NewEventProducer()

	if err != nil {
		log.Fatalf("nsqConn.NewEventProducer: cannot initialize to nsq connection: %v", err)
	}

	// Initialize Manager
	paginationManager := manager.NewPaginationManager(
		manager.WithDefaultPage(1),
		manager.WithDefaultPageSize(10),
	)
	passwordManager := manager.NewPasswordManager(bcrypt.DefaultCost)
	tokenManager := manager.NewTokenManager(
		manager.WithSigningMethod(jwt.SigningMethodHS256),
		manager.WithAccessTokenExpiresIn(config.GetAccessTokenExpirationTime()),
		manager.WithAccessTokenSecretKey(config.GetAccessTokenSecretKey()),
		manager.WithRefreshTokenExpiresIn(config.GetRefreshTokenExpirationTime()),
		manager.WithRefreshTokenSecretKey(config.GetRefreshTokenSecretKey()),
		manager.WithLogger(log),
	)

	// Initialize Repository
	postgresSeratRepository := postgres.NewPostgresSeratRepository(db)
	postgresWulanganRepository := postgres.NewPostgresWulanganRepository(db)
	postgresUserRepository := postgres.NewPostgresUserRepository(db)
	redisRepository := redis.NewRedisRepository(rc, log)
	redisSeratRepository := redis.NewRedisSeratRepository(redisRepository, postgresSeratRepository, log)

	// Initialize event emitter
	eventEmitter := event.NewEventEmitter(producer)
	seratEventEmitter := event.NewSeratEventEmitter(eventEmitter)
	wulanganEventEmitter := event.NewWulanganEventEmitter(eventEmitter)
	sessionEventEmitter := event.NewSessionEventEmitter(eventEmitter)
	userEventEmitter := event.NewUserEventEmitter(eventEmitter)

	// Initialize Service
	createSeratService := svc.NewCreateSeratService(redisSeratRepository, seratEventEmitter)
	updateSeratService := svc.NewUpdateSeratService(redisSeratRepository, seratEventEmitter)
	getSeratService := svc.NewGetSeratService(redisSeratRepository, seratEventEmitter)
	listSeratService := svc.NewListSeratsService(log, redisSeratRepository, seratEventEmitter)
	deleteSeratService := svc.NewDeleteSeratService(redisSeratRepository, seratEventEmitter)

	createWulanganService := svc.NewCreateWulanganService(postgresWulanganRepository, wulanganEventEmitter)
	getWulanganService := svc.NewGetWulanganService(postgresWulanganRepository, wulanganEventEmitter)
	registerService := svc.NewRegisterService(log, postgresUserRepository, userEventEmitter, passwordManager, tokenManager)
	loginService := svc.NewLoginService(log, postgresUserRepository, sessionEventEmitter, passwordManager, tokenManager)

	// Initialize Service
	seratsServer = srv.NewSeratsServer(
		srv.WithCreateSeratService(createSeratService),
		srv.WithGetSeratService(getSeratService),
		srv.WithListSeratsService(listSeratService),
		srv.WithUpdateSeratService(updateSeratService),
		srv.WithDeleteSeratService(deleteSeratService),
		srv.WithPaginationManager(paginationManager),
		srv.WithLogger[*srv.SeratsServer](log),
	)

	wulangansServer = srv.NewWulangansServer(
		srv.WithCreateWulanganService(createWulanganService),
		srv.WithGetWulanganService(getWulanganService),
		srv.WithLogger[*srv.WulangansServer](log),
	)

	authenticationServer = srv.NewAuthenticationsServer(
		srv.WithRegisterService(registerService),
		srv.WithLoginService(loginService),
		srv.WithLogger[*srv.AuthenticationServer](log),
	)
}

func runGRPCServer() error {
	proto.RegisterSeratsServer(server, seratsServer)
	proto.RegisterWulangansServer(server, wulangansServer)
	proto.RegisterAuthenticationServer(server, authenticationServer)

	return server.Serve(tcp)
}

func runGatewayServer(ctx context.Context) (err error) {
	err = proto.RegisterSeratsHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	err = proto.RegisterWulangansHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	err = proto.RegisterAuthenticationHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: srv.Cors(mux),
	}

	return srv.ListenAndServe()
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Run grpc server as go routine
	go runGRPCServer()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return runGatewayServer(ctx)
}

func main() {
	flag.Parse()

	grpclog.SetLoggerV2(log)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
