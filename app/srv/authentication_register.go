package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/api"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *AuthenticationServer) Register(ctx context.Context, req *api.RegisterRequest) (res *api.Session, err error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	result, err := s.registerService.Call(ctx, &svc.RegisterParams{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res = &api.Session{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiredAt:    timestamppb.New(result.ExpiredAt),
	}

	return res, nil
}
