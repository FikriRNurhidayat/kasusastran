package usecase

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	"github.com/google/uuid"
)

type GetSeratUseCaser interface {
	Exec(ctx context.Context, id uuid.UUID) (*entity.Serat, error)
}

type GetSeratUseCase struct {
	SeratRepository postgres.SeratRepositorier
}

func NewGetSeratUseCase(seratRepository postgres.SeratRepositorier) GetSeratUseCaser {
	return &GetSeratUseCase{
		SeratRepository: seratRepository,
	}
}

func (u *GetSeratUseCase) Exec(ctx context.Context, id uuid.UUID) (*entity.Serat, error) {
	return u.SeratRepository.GetSerat(ctx, id)
}
