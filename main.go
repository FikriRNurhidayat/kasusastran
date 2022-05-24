package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	pg "github.com/fikrirnurhidayat/kasusastran/app/driver/postgres"

	"github.com/fikrirnurhidayat/kasusastran/app/config"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/usecase"
	"github.com/fikrirnurhidayat/kasusastran/app/svc/seratssvc"
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
var Query query.Querier

// Services
var SeratsService *seratssvc.SeratsService

// Usecases
var CreateSeratUseCase usecase.CreateSeratUseCaser
var UpdateSeratUseCaes usecase.UpdateSeratUseCaser
var DeleteSeratUseCase usecase.DeleteSeratUseCaser
var GetSeratUseCase usecase.GetSeratUseCaser
var ListSeratUseCase usecase.ListSeratsUseCaser

// Repositories
var SeratRepository postgres.SeratRepositorier

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
	Query = query.New(db)

	// Initialize Repository
	SeratRepository = postgres.NewSeratRepository(Query)

	// Initialize Usecase
	CreateSeratUseCase = usecase.NewCreateSeratUseCase(SeratRepository)
	UpdateSeratUseCaes = usecase.NewUpdateSeratUseCase(SeratRepository)
	GetSeratUseCase = usecase.NewGetSeratUseCase(SeratRepository)
	ListSeratUseCase = usecase.NewListSeratsUseCase(SeratRepository)
	DeleteSeratUseCase = usecase.NewDeleteSeratUseCase(SeratRepository)

	// Initialize Service
	SeratsService = seratssvc.
		New().
		SetGetSeratUseCase(GetSeratUseCase).
		SetListSeratsUseCase(ListSeratUseCase).
		SetUpdateSeratUseCase(UpdateSeratUseCaes).
		SetDeleteSeratUseCase(DeleteSeratUseCase).
		SetCreateSeratUseCase(CreateSeratUseCase)
}

func runGRPCServer() error {
	api.RegisterSeratsServer(server, SeratsService)

	return server.Serve(tcp)
}

func runGatewayServer(ctx context.Context) error {
	err := api.RegisterSeratsHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
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
