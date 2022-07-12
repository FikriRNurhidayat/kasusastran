package srv

import (
	api "github.com/fikrirnurhidayat/kasusastran/api"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
)

type WulangansServer struct {
	api.UnimplementedWulangansServer

	createWulanganService svc.CreateWulanganService
	getWulanganService    svc.GetWulanganService
}

func NewWulangansServer(setters ...WulangansServerSetter) *WulangansServer {
	s := new(WulangansServer)

	for _, set := range setters {
		set(s)
	}

	return s
}

type WulangansServerSetter func(*WulangansServer)

func WithCreateWulanganService(createWulanganService svc.CreateWulanganService) WulangansServerSetter {
	return func(server *WulangansServer) {
		server.createWulanganService = createWulanganService
	}
}

func WithGetWulanganService(getWulanganService svc.GetWulanganService) WulangansServerSetter {
	return func(server *WulangansServer) {
		server.getWulanganService = getWulanganService
	}
}

func WithListWulangansService() WulangansServerSetter {
	return func(server *WulangansServer) {
	}
}

func WithUpdateWulanganService() WulangansServerSetter {
	return func(server *WulangansServer) {
	}
}

func WithDeleteWulanganService() WulangansServerSetter {
	return func(server *WulangansServer) {
	}
}
