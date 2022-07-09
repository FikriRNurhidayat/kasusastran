package srv

import (
	api "github.com/fikrirnurhidayat/kasusastran/api"
)

type WulangansServer struct {
	api.UnimplementedWulangansServer
}

func NewWulangansServer() *WulangansServer {
	return &WulangansServer{}
}

type WulangansServerSetter func(*WulangansServer)

func WithCreateWulanganService() WulangansServerSetter {
	return func(server *WulangansServer) {
	}
}

func WithGetWulanganService() WulangansServerSetter {
	return func(server *WulangansServer) {
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
