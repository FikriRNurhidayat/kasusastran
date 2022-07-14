package srv

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"google.golang.org/grpc/grpclog"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/proto"
)

type SeratsServer struct {
	proto.UnimplementedSeratsServer

	logger             grpclog.LoggerV2
	createSeratService svc.CreateSeratService
	updateSeratService svc.UpdateSeratService
	getSeratService    svc.GetSeratService
	deleteSeratService svc.DeleteSeratService
	listSeratsService  svc.ListSeratsService
	paginationManager  manager.PaginationManager
}

func (s *SeratsServer) setLogger(logger grpclog.LoggerV2) {
	s.logger = logger
}

func NewSeratsServer(setters ...SeratsServerSetter) *SeratsServer {
	server := &SeratsServer{}

	for _, set := range setters {
		set(server)
	}

	return server
}

type SeratsServerSetter func(*SeratsServer)

func WithPaginationManager(paginationManager manager.PaginationManager) SeratsServerSetter {
	return func(server *SeratsServer) {
		server.paginationManager = paginationManager
	}
}

func WithGetSeratService(getSeratService svc.GetSeratService) SeratsServerSetter {
	return func(server *SeratsServer) {
		server.getSeratService = getSeratService
	}
}

func WithListSeratsService(listSeratsService svc.ListSeratsService) SeratsServerSetter {
	return func(server *SeratsServer) {
		server.listSeratsService = listSeratsService
	}
}

func WithCreateSeratService(createSeratService svc.CreateSeratService) SeratsServerSetter {
	return func(server *SeratsServer) {
		server.createSeratService = createSeratService
	}
}

func WithUpdateSeratService(updateSeratService svc.UpdateSeratService) SeratsServerSetter {
	return func(server *SeratsServer) {
		server.updateSeratService = updateSeratService
	}
}

func WithDeleteSeratService(deleteSeratService svc.DeleteSeratService) SeratsServerSetter {
	return func(server *SeratsServer) {
		server.deleteSeratService = deleteSeratService
	}
}
