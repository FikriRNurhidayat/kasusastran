package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *AuthenticationServer) Register(ctx context.Context, req *proto.RegisterRequest) (res *proto.Session, err error) {
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

	res = &proto.Session{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiredAt:    timestamppb.New(result.ExpiredAt),
	}

	return res, nil
}
