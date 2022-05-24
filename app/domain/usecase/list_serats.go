package usecase

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
)

type ListSeratsUseCaser interface {
	Exec(context.Context, *entity.Pagination) ([]entity.Serat, *entity.Pagination, error)
}

type ListSeratsUseCase struct {
	SeratRepository postgres.SeratRepositorier
}

func NewListSeratsUseCase(SeratRepository postgres.SeratRepositorier) ListSeratsUseCaser {
	return &ListSeratsUseCase{
		SeratRepository: SeratRepository,
	}
}

func (u *ListSeratsUseCase) Exec(ctx context.Context, ipg *entity.Pagination) (serats []entity.Serat, opg *entity.Pagination, err error) {
	serats, count, err := u.SeratRepository.ListSerats(ctx, ipg)

	if err != nil {
		return
	}

	opg = entity.NewPagination(ipg.GetLimit(), ipg.GetOffset(), count)

	return
}
