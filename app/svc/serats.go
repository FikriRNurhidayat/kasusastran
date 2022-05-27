package svc

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/usecase"

	api "github.com/fikrirnurhidayat/kasusastran/api"
)

type SeratsService struct {
	api.UnimplementedSeratsServer

	CreateSeratUseCase usecase.CreateSeratUseCaser
	UpdateSeratUseCase usecase.UpdateSeratUseCaser
	GetSeratUseCase    usecase.GetSeratUseCaser
	DeleteSeratUseCase usecase.DeleteSeratUseCaser
	ListSeratsUseCase  usecase.ListSeratsUseCaser
}

func NewSeratsService() *SeratsService {
	return &SeratsService{}
}

func (s *SeratsService) SetCreateSeratUseCase(CreateSeratUseCase usecase.CreateSeratUseCaser) *SeratsService {
	s.CreateSeratUseCase = CreateSeratUseCase
	return s
}

func (s *SeratsService) SetUpdateSeratUseCase(UpdateSeratUseCase usecase.UpdateSeratUseCaser) *SeratsService {
	s.UpdateSeratUseCase = UpdateSeratUseCase
	return s
}

func (s *SeratsService) SetDeleteSeratUseCase(DeleteSeratUseCase usecase.DeleteSeratUseCaser) *SeratsService {
	s.DeleteSeratUseCase = DeleteSeratUseCase
	return s
}

func (s *SeratsService) SetGetSeratUseCase(GetSeratUseCase usecase.GetSeratUseCaser) *SeratsService {
	s.GetSeratUseCase = GetSeratUseCase
	return s
}

func (s *SeratsService) SetListSeratsUseCase(ListSeratsUseCase usecase.ListSeratsUseCaser) *SeratsService {
	s.ListSeratsUseCase = ListSeratsUseCase
	return s
}
