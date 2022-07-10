package svc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type GetSeratService interface {
	Call(ctx context.Context, id uuid.UUID) (*GetSeratResult, error)
}

type GetSeratResult entity.Serat

type getSeratService struct {
	seratRepository repository.SeratRepository
}

func NewGetSeratService(seratRepository repository.SeratRepository) GetSeratService {
	return &getSeratService{
		seratRepository: seratRepository,
	}
}

func (u *getSeratService) Call(ctx context.Context, id uuid.UUID) (*GetSeratResult, error) {
	serat, err := u.seratRepository.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	res := &GetSeratResult{
		ID:                serat.ID,
		Title:             serat.Title,
		Description:       serat.Description,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	return res, nil
}
