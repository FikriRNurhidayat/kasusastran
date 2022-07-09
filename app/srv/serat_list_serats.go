package srv

import (
	"context"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SeratsServer) ListSerats(ctx context.Context, req *api.ListSeratsRequest) (res *api.ListSeratsResponse, err error) {
	pagination := &entity.Pagination{}

	if req.GetPagination() != nil {
		if req.GetPagination().GetPage() != 0 {
			pagination.Page = req.GetPagination().GetPage()
		}

		if req.GetPagination().GetPageSize() != 0 {
			pagination.PageSize = req.GetPagination().GetPageSize()
		}
	}

	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}

	svc, pagination, err := s.listSeratsService.Exec(ctx, pagination)

	if err != nil {
		return res, status.Errorf(codes.Internal, "failed to retrieve list of svc: %v", err)
	}

	res = &api.ListSeratsResponse{
		Meta: &api.ListSeratsResponse_MetaListSerats{
			Pagination: &api.ResponsePagination{
				Page:      pagination.Page,
				PageSize:  pagination.PageSize,
				PageCount: pagination.PageCount,
				Total:     pagination.Total,
			},
		},
		Serats: []*api.Serat{},
	}

	for _, serat := range svc {
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
