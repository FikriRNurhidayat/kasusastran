package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	nsq "github.com/fikrirnurhidayat/kasusastran/app/driver/nsq"
	pg "github.com/fikrirnurhidayat/kasusastran/app/driver/postgres"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/fikrirnurhidayat/kasusastran/app/config"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
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
	seratRepository := postgres.NewPostgresSeratRepository(db)
	wulanganRepository := postgres.NewPostgresWulanganRepository(db)
	userRepository := postgres.NewPostgresUserRepository(db)

	// Initialize event emitter
	eventEmitter := event.NewEventEmitter(producer)
	seratEventEmitter := event.NewSeratEventEmitter(eventEmitter)
	wulanganEventEmitter := event.NewWulanganEventEmitter(eventEmitter)
	sessionEventEmitter := event.NewSessionEventEmitter(eventEmitter)
	userEventEmitter := event.NewUserEventEmitter(eventEmitter)

	// Initialize Service
	createSeratService := svc.NewCreateSeratService(seratRepository, seratEventEmitter)
	updateSeratService := svc.NewUpdateSeratService(seratRepository, seratEventEmitter)
	getSeratService := svc.NewGetSeratService(seratRepository, seratEventEmitter)
	listSeratService := svc.NewListSeratsService(seratRepository, seratEventEmitter)
	deleteSeratService := svc.NewDeleteSeratService(seratRepository, seratEventEmitter)
	createWulanganService := svc.NewCreateWulanganService(wulanganRepository, wulanganEventEmitter)
	getWulanganService := svc.NewGetWulanganService(wulanganRepository, wulanganEventEmitter)
	registerService := svc.NewRegisterService(userRepository, userEventEmitter, passwordManager, tokenManager)
	loginService := svc.NewLoginService(userRepository, sessionEventEmitter, passwordManager, tokenManager)

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
	api.RegisterSeratsServer(server, seratsServer)
	api.RegisterWulangansServer(server, wulangansServer)
	api.RegisterAuthenticationServer(server, authenticationServer)

	return server.Serve(tcp)
}

func runGatewayServer(ctx context.Context) (err error) {
	err = api.RegisterSeratsHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	err = api.RegisterWulangansHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	err = api.RegisterAuthenticationHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: cors(mux),
	}

	return srv.ListenAndServe()
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
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
