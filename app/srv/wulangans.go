package srv

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/proto"
	"google.golang.org/grpc/grpclog"
)

type WulangansServer struct {
	proto.UnimplementedWulangansServer

	logger                grpclog.LoggerV2
	createWulanganService svc.CreateWulanganService
	getWulanganService    svc.GetWulanganService
}

func (s *WulangansServer) setLogger(logger grpclog.LoggerV2) {
	s.logger = logger
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
