package srv

import (
	"context"

	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/svc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/fikrirnurhidayat/api.kasusastran.io/api"
)

func (s *SeratsServer) CreateSerat(ctx context.Context, req *api.CreateSeratRequest) (*api.Serat, error) {
	serat, err := s.createSeratService.Exec(ctx, &svc.CreateSeratParams{
		Title:             req.GetTitle(),
		Description:       req.GetDescription(),
		Content:           req.GetContent(),
		CoverImageUrl:     req.GetCoverImageUrl(),
		ThumbnailImageUrl: req.GetThumbnailImageUrl(),
	})

	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "s.CreateSeratUseCase.Exec: %v", err.Error())
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
