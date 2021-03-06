package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/fikrirnurhidayat/kasusastran/proto"
)

func (s *SeratsServer) UpdateSerat(ctx context.Context, req *proto.UpdateSeratRequest) (*proto.Serat, error) {
	id, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.OutOfRange, err.Error())
	}

	serat, err := s.updateSeratService.Call(ctx, id, &svc.UpdateSeratParams{
		Title:             req.GetTitle(),
		Description:       req.GetDescription(),
		CoverImageUrl:     req.GetCoverImageUrl(),
		ThumbnailImageUrl: req.GetThumbnailImageUrl(),
	})

	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "s.UpdateSeratUseCase.Call: %v", err.Error())
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
