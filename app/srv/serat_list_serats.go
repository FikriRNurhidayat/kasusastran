package srv

import (
	"context"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SeratsServer) ListSerats(ctx context.Context, req *api.ListSeratsRequest) (res *api.ListSeratsResponse, err error) {
	result, err := s.listSeratsService.Call(ctx, &svc.ListSeratsParams{
		Pagination: s.paginationManager.FromIncomingRequest(req.GetPagination()),
	})

	if err != nil {
		return res, status.Errorf(codes.Internal, "failed to retrieve list of svc: %v", err)
	}

	res = &api.ListSeratsResponse{
		Meta: &api.ListSeratsResponse_MetaListSerats{
			Pagination: s.paginationManager.NewOutgoingResponse(result.Pagination),
		},
		Serats: []*api.Serat{},
	}

	for _, serat := range result.Serats {
		pack := &api.Serat{
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
