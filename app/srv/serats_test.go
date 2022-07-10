package srv_test

import mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"

type MockSeratsServer struct {
	createSeratService *mocks.CreateSeratService
	updateSeratService *mocks.UpdateSeratService
	getSeratService    *mocks.GetSeratService
	deleteSeratService *mocks.DeleteSeratService
	listSeratsService  *mocks.ListSeratsService
}
