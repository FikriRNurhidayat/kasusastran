package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/lib/troublemaker"

	"github.com/fikrirnurhidayat/kasusastran/proto"
)

func (s *SeratsServer) CreateSerat(ctx context.Context, req *proto.CreateSeratRequest) (res *proto.Serat, err error) {
	if err := req.Validate(); err != nil {
		return nil, troublemaker.FromValidationError(err)
	}

	var serat *svc.CreateSeratResult
	if serat, err = s.createSeratService.Call(ctx, &svc.CreateSeratParams{
		Title:             req.GetTitle(),
		Description:       req.GetDescription(),
		CoverImageUrl:     req.GetCoverImageUrl(),
		ThumbnailImageUrl: req.GetThumbnailImageUrl(),
	}); err != nil {
		return nil, err
	}

	res = &proto.Serat{
		Id:                serat.ID.String(),
		Title:             serat.Title,
		Description:       serat.Description,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	return res, nil
}
