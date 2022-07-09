package svc

import (
	"context"

	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/entity"
	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/repository"
	"github.com/google/uuid"
)

type GetSeratService interface {
	Exec(ctx context.Context, id uuid.UUID) (*entity.Serat, error)
}

type getSeratService struct {
	seratRepository repository.SeratRepository
}

func NewGetSeratService(seratRepository repository.SeratRepository) GetSeratService {
	return &getSeratService{
		seratRepository: seratRepository,
	}
}

func (u *getSeratService) Exec(ctx context.Context, id uuid.UUID) (*entity.Serat, error) {
	return u.seratRepository.Get(ctx, id)
}
