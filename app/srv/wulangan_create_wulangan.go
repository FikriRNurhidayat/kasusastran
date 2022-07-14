package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *WulangansServer) CreateWulangan(ctx context.Context, req *proto.CreateWulanganRequest) (*proto.Wulangan, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.OutOfRange, err.Error())
	}

	pack, err := s.createWulanganService.Call(ctx, &svc.CreateWulanganParams{
		Title:             req.GetTitle(),
		Description:       req.GetDescription(),
		CoverImageUrl:     req.GetCoverImageUrl(),
		ThumbnailImageUrl: req.GetThumbnailImageUrl(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
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
