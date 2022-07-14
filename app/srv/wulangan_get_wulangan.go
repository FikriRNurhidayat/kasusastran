package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *WulangansServer) GetWulangan(ctx context.Context, req *proto.GetWulanganRequest) (*proto.Wulangan, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id := uuid.MustParse(req.GetId())
	pack, err := s.getWulanganService.Call(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	res := &proto.Wulangan{
		Id:                pack.ID.String(),
		Title:             pack.Title,
		Description:       pack.Description,
		CoverImageUrl:     pack.CoverImageUrl,
		ThumbnailImageUrl: pack.ThumbnailImageUrl,
		CreatedAt:         timestamppb.New(pack.CreatedAt),
		UpdatedAt:         timestamppb.New(pack.UpdatedAt),
	}

	return res, nil
}
