package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	pg "github.com/fikrirnurhidayat/kasusastran/app/driver/postgres"

	"github.com/fikrirnurhidayat/kasusastran/app/config"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
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
var dbQuery query.Querier

// Services
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
	db, err := pg.Connect(dbURL)

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
	dbQuery = query.New(db)

	// Initialize Repository
	seratRepository := postgres.NewPostgresSeratRepository(dbQuery)
	wulanganRepository := postgres.NewPostgresWulanganRepository(dbQuery)

	// Initialize Service
	createSeratService := svc.NewCreateSeratService(seratRepository)
	updateSeratService := svc.NewUpdateSeratService(seratRepository)
	getSeratService := svc.NewGetSeratService(seratRepository)
	listSeratService := svc.NewListSeratsService(seratRepository)
	deleteSeratService := svc.NewDeleteSeratService(seratRepository)
	createWulanganService := svc.NewCreateWulanganService(wulanganRepository)

	// Initialize Service
	seratsServer = srv.NewSeratsServer(
		srv.WithCreateSeratService(createSeratService),
		srv.WithGetSeratService(getSeratService),
		srv.WithListSeratsService(listSeratService),
		srv.WithUpdateSeratService(updateSeratService),
		srv.WithDeleteSeratService(deleteSeratService),
	)

	wulangansServer = srv.NewWulangansServer(
		srv.WithCreateWulanganService(createWulanganService),
	)
}

func runGRPCServer() error {
	api.RegisterSeratsServer(server, seratsServer)
	api.RegisterWulangansServer(server, wulangansServer)

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

	srv := &http.Server{
		Addr:    ":8081",
		Handler: cors(mux),
	}

	return srv.ListenAndServe()
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
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
