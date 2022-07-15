package srv

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/fikrirnurhidayat/kasusastran/proto"
)

func (s *SeratsServer) GetSerat(ctx context.Context, req *proto.GetSeratRequest) (*proto.Serat, error) {
	id, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, err
	}

	serat, err := s.getSeratService.Call(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "serat not found: %v", req.GetId())
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
