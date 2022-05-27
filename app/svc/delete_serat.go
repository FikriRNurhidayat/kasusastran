package svc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	api "github.com/fikrirnurhidayat/kasusastran/api"
)

func (s *SeratsService) DeleteSerat(ctx context.Context, req *api.DeleteSeratRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, errors.ErrInvalidUUID
	}

	err = s.DeleteSeratUseCase.Exec(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "serat not found: %v", req.GetId())
	}

	return &emptypb.Empty{}, nil
}
