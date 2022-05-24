package seratssvc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/errors"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/usecase"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/fikrirnurhidayat/kasusastran/api"
)

func (s *SeratsService) UpdateSerat(ctx context.Context, req *api.UpdateSeratRequest) (*api.Serat, error) {
	id, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, errors.ErrInvalidUUID
	}

	serat, err := s.UpdateSeratUseCase.Exec(ctx, id, &usecase.UpdateSeratParams{
		Title:             req.GetTitle(),
		Description:       req.GetDescription(),
		Body:              req.GetBody(),
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
		Body:              serat.Body,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	return res, nil
}
