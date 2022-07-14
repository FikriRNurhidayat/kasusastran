package srv_test

import mockService "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"

type MockAuthenticationServer struct {
	registerService *mockService.RegisterService
	loginService    *mockService.LoginService
}
