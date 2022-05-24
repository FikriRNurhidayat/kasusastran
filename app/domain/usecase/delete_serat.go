package usecase

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	"github.com/google/uuid"
)

type DeleteSeratUseCaser interface {
	Exec(ctx context.Context, id uuid.UUID) error
}

type DeleteSeratUseCase struct {
	SeratRepository postgres.SeratRepositorier
}

func NewDeleteSeratUseCase(seratRepository postgres.SeratRepositorier) DeleteSeratUseCaser {
	return &DeleteSeratUseCase{
		SeratRepository: seratRepository,
	}
}

func (u *DeleteSeratUseCase) Exec(ctx context.Context, id uuid.UUID) error {
	return u.SeratRepository.DeleteSerat(ctx, id)
}
