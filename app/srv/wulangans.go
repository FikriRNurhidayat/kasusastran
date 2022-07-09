package srv

import (
	api "github.com/fikrirnurhidayat/api.kasusastran.io/api"
)

type WulangansServer struct {
	api.UnimplementedWulangansServer
}

func NewWulangansServer() *WulangansServer {
	return &WulangansServer{}
}
