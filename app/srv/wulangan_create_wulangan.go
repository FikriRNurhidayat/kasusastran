package srv

import (
	"context"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *WulangansServer) CreateWulangan(ctx context.Context, req *api.CreateWulanganRequest) (*api.Wulangan, error) {
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

	res := &api.Wulangan{
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
