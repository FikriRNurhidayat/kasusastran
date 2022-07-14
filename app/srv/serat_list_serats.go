package srv

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SeratsServer) ListSerats(ctx context.Context, req *proto.ListSeratsRequest) (res *proto.ListSeratsResponse, err error) {
	result, err := s.listSeratsService.Call(ctx, &svc.ListSeratsParams{
		Pagination: s.paginationManager.FromIncomingRequest(req.GetPagination()),
	})

	if err != nil {
		return res, status.Errorf(codes.Internal, "failed to retrieve list of svc: %v", err)
	}

	res = &proto.ListSeratsResponse{
		Meta: &proto.ListSeratsResponse_MetaListSerats{
			Pagination: s.paginationManager.NewOutgoingResponse(result.Pagination),
		},
		Serats: []*proto.Serat{},
	}

	for _, serat := range result.Serats {
		pack := &proto.Serat{
			Id:                serat.ID.String(),
			Title:             serat.Title,
			Description:       serat.Description,
			CoverImageUrl:     serat.CoverImageUrl,
			ThumbnailImageUrl: serat.ThumbnailImageUrl,
		}

		res.Serats = append(res.Serats, pack)
	}

	return
}
