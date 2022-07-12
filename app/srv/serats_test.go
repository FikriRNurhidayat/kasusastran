package srv_test

import (
	mockManager "github.com/fikrirnurhidayat/kasusastran/mocks/domain/manager"
	mockService "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"
)

type MockSeratsServer struct {
	createSeratService *mockService.CreateSeratService
	updateSeratService *mockService.UpdateSeratService
	getSeratService    *mockService.GetSeratService
	deleteSeratService *mockService.DeleteSeratService
	listSeratsService  *mockService.ListSeratsService
	paginationManager  *mockManager.PaginationManager
}
