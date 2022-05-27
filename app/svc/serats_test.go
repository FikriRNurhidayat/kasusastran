package svc_test

import mockUseCase "github.com/fikrirnurhidayat/kasusastran/mocks/domain/usecase"

type MockSeratService struct {
	CreateSeratUseCase *mockUseCase.CreateSeratUseCaser
	UpdateSeratUseCase *mockUseCase.UpdateSeratUseCaser
	GetSeratUseCase    *mockUseCase.GetSeratUseCaser
	DeleteSeratUseCase *mockUseCase.DeleteSeratUseCaser
	ListSeratsUseCase  *mockUseCase.ListSeratsUseCaser
}
