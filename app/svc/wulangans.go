package svc

import (
	api "github.com/fikrirnurhidayat/kasusastran/api"
)

type WulangansService struct {
	api.UnimplementedWulangansServer
}

func NewWulangansService() *WulangansService {
	return &WulangansService{}
}
