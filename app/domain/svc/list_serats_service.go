package svc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
)

type ListSeratsService interface {
	Call(context.Context, *ListSeratsParams) (*ListSeratsResult, error)
}

type ListSeratsResult struct {
	Serats     []*entity.Serat
	Pagination *manager.Pagination
}

type ListSeratsParams struct {
	Pagination *manager.Pagination
}

type listSeratsService struct {
	seratRepository repository.SeratRepository
}

func NewListSeratsService(SeratRepository repository.SeratRepository) ListSeratsService {
	return &listSeratsService{
		seratRepository: SeratRepository,
	}
}

func (u *listSeratsService) Call(ctx context.Context, params *ListSeratsParams) (*ListSeratsResult, error) {
	serats, count, err := u.seratRepository.List(ctx, params.Pagination.ToListQuery())

	if err != nil {
		return nil, err
	}

	res := &ListSeratsResult{
		Serats:     serats,
		Pagination: params.Pagination.WithTotal(count),
	}

	return res, nil
}
