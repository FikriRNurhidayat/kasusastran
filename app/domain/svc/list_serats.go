package svc

import (
	"context"

	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/entity"
	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/repository"
)

type ListSeratsService interface {
	Exec(context.Context, *entity.Pagination) ([]entity.Serat, *entity.Pagination, error)
}

type listSeratsService struct {
	seratRepository repository.SeratRepository
}

func NewListSeratsService(SeratRepository repository.SeratRepository) ListSeratsService {
	return &listSeratsService{
		seratRepository: SeratRepository,
	}
}

func (u *listSeratsService) Exec(ctx context.Context, ipg *entity.Pagination) (serats []entity.Serat, opg *entity.Pagination, err error) {
	serats, count, err := u.seratRepository.List(ctx, ipg)

	if err != nil {
		return
	}

	opg = entity.NewPagination(ipg.GetLimit(), ipg.GetOffset(), count)

	return
}
