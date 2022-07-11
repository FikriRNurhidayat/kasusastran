package srv

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	man "github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
)

type SeratsServer struct {
	api.UnimplementedSeratsServer

	createSeratService svc.CreateSeratService
	updateSeratService svc.UpdateSeratService
	getSeratService    svc.GetSeratService
	deleteSeratService svc.DeleteSeratService
	listSeratsService  svc.ListSeratsService

	paginationManager man.PaginationManager
}

func NewSeratsServer(setters ...SeratsServerSetter) *SeratsServer {
	server := &SeratsServer{}

	for _, set := range setters {
		set(server)
	}

	return server
}

type SeratsServerSetter func(*SeratsServer)

func WithPaginationManager(paginationManager man.PaginationManager) SeratsServerSetter {
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
