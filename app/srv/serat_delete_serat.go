package srv

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/fikrirnurhidayat/kasusastran/proto"
)

func (s *SeratsServer) DeleteSerat(ctx context.Context, req *proto.DeleteSeratRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, err
	}

	err = s.deleteSeratService.Call(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "serat not found: %v", req.GetId())
	}

	return &emptypb.Empty{}, nil
}
