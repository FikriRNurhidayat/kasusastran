package srv

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/proto"
	"google.golang.org/grpc/grpclog"
)

type AuthenticationServer struct {
	proto.UnimplementedAuthenticationServer

	logger          grpclog.LoggerV2
	registerService svc.RegisterService
	loginService    svc.LoginService
}

func (s *AuthenticationServer) setLogger(logger grpclog.LoggerV2) {
	s.logger = logger
}

func NewAuthenticationsServer(setters ...AuthenticationServerSetter) *AuthenticationServer {
	server := &AuthenticationServer{}

	for _, set := range setters {
		set(server)
	}

	return server
}

type AuthenticationServerSetter func(*AuthenticationServer)

func WithRegisterService(registerService svc.RegisterService) AuthenticationServerSetter {
	return func(s *AuthenticationServer) {
		s.registerService = registerService
	}
}

func WithLoginService(loginService svc.LoginService) AuthenticationServerSetter {
	return func(s *AuthenticationServer) {
		s.loginService = loginService
	}
}
