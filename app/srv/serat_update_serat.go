package srv

import (
	"context"

	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/errors"
	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/svc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/fikrirnurhidayat/api.kasusastran.io/api"
)

func (s *SeratsServer) UpdateSerat(ctx context.Context, req *api.UpdateSeratRequest) (*api.Serat, error) {
	id, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, errors.ErrInvalidUUID
	}

	serat, err := s.updateSeratService.Exec(ctx, id, &svc.UpdateSeratParams{
		Title:             req.GetTitle(),
		Description:       req.GetDescription(),
		Content:           req.GetContent(),
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
		Content:           serat.Content,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	return res, nil
}
