package svc

import (
	"context"

	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/repository"
	"github.com/google/uuid"
)

type DeleteSeratService interface {
	Exec(ctx context.Context, id uuid.UUID) error
}

type deleteSeratService struct {
	seratRepository repository.SeratRepository
}

func NewDeleteSeratService(seratRepository repository.SeratRepository) DeleteSeratService {
	return &deleteSeratService{
		seratRepository: seratRepository,
	}
}

func (u *deleteSeratService) Exec(ctx context.Context, id uuid.UUID) error {
	return u.seratRepository.Delete(ctx, id)
}
