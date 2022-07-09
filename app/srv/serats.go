package srv

import (
	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/svc"

	api "github.com/fikrirnurhidayat/api.kasusastran.io/api"
)

type SeratsServer struct {
	api.UnimplementedSeratsServer

	createSeratService svc.CreateSeratService
	updateSeratService svc.UpdateSeratService
	getSeratService    svc.GetSeratService
	deleteSeratService svc.DeleteSeratService
	listSeratsService  svc.ListSeratsService
}

func NewSeratsServer() *SeratsServer {
	return &SeratsServer{}
}

func (s *SeratsServer) SetCreateSeratUseCase(CreateSeratUseCase svc.CreateSeratService) *SeratsServer {
	s.createSeratService = CreateSeratUseCase
	return s
}

func (s *SeratsServer) SetUpdateSeratUseCase(UpdateSeratUseCase svc.UpdateSeratService) *SeratsServer {
	s.updateSeratService = UpdateSeratUseCase
	return s
}

func (s *SeratsServer) SetDeleteSeratUseCase(DeleteSeratUseCase svc.DeleteSeratService) *SeratsServer {
	s.deleteSeratService = DeleteSeratUseCase
	return s
}

func (s *SeratsServer) SetGetSeratUseCase(GetSeratUseCase svc.GetSeratService) *SeratsServer {
	s.getSeratService = GetSeratUseCase
	return s
}

func (s *SeratsServer) SetListSeratsUseCase(ListSeratsUseCase svc.ListSeratsService) *SeratsServer {
	s.listSeratsService = ListSeratsUseCase
	return s
}
