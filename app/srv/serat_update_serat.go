package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/errors"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/fikrirnurhidayat/kasusastran/api"
)

func (s *SeratsServer) UpdateSerat(ctx context.Context, req *api.UpdateSeratRequest) (*api.Serat, error) {
	id, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, errors.ErrInvalidUUID
	}

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.OutOfRange, err.Error())
	}

	serat, err := s.updateSeratService.Exec(ctx, id, &svc.UpdateSeratParams{
		Title:             req.GetTitle(),
		Description:       req.GetDescription(),
		CoverImageUrl:     req.GetCoverImageUrl(),
		ThumbnailImageUrl: req.GetThumbnailImageUrl(),
	})

	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "s.UpdateSeratUseCase.Exec: %v", err.Error())
	}

	res := &api.Serat{
		Id:                serat.ID.String(),
		Title:             serat.Title,
		Description:       serat.Description,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	return res, nil
}
