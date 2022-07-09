package svc

import (
	"context"

	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/entity"
	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/repository"
	"github.com/google/uuid"
)

type UpdateSeratParams struct {
	Title             string
	Description       string
	Content           string
	CoverImageUrl     string
	ThumbnailImageUrl string
}

type UpdateSeratService interface {
	Exec(ctx context.Context, id uuid.UUID, params *UpdateSeratParams) (*entity.Serat, error)
}

type updateSeratService struct {
	SeratRepository repository.SeratRepository
}

func NewUpdateSeratService(seratRepository repository.SeratRepository) UpdateSeratService {
	return &updateSeratService{
		SeratRepository: seratRepository,
	}
}

func (u *updateSeratService) Exec(ctx context.Context, id uuid.UUID, params *UpdateSeratParams) (*entity.Serat, error) {
	return u.SeratRepository.Update(ctx, id, &entity.Serat{
		Title:             params.Title,
		Description:       params.Description,
		Content:           params.Content,
		CoverImageUrl:     params.CoverImageUrl,
		ThumbnailImageUrl: params.ThumbnailImageUrl,
	})
}
