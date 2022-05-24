package seratssvc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/fikrirnurhidayat/kasusastran/api"
)

func (s *SeratsService) GetSerat(ctx context.Context, req *api.GetSeratRequest) (*api.Serat, error) {
	id, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, errors.ErrInvalidUUID
	}

	serat, err := s.GetSeratUseCase.Exec(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "serat not found: %v", req.GetId())
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
