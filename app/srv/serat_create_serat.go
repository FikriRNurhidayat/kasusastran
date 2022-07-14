package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/fikrirnurhidayat/kasusastran/proto"
)

func (s *SeratsServer) CreateSerat(ctx context.Context, req *proto.CreateSeratRequest) (*proto.Serat, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.OutOfRange, err.Error())
	}

	serat, err := s.createSeratService.Call(ctx, &svc.CreateSeratParams{
		Title:             req.GetTitle(),
		Description:       req.GetDescription(),
		CoverImageUrl:     req.GetCoverImageUrl(),
		ThumbnailImageUrl: req.GetThumbnailImageUrl(),
	})

	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "s.CreateSeratUseCase.Call: %v", err.Error())
	}

	res := &proto.Serat{
		Id:                serat.ID.String(),
		Title:             serat.Title,
		Description:       serat.Description,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	return res, nil
}
