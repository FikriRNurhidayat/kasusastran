package usecase

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	"github.com/google/uuid"
)

type UpdateSeratParams struct {
	Title             string
	Description       string
	Body              string
	CoverImageUrl     string
	ThumbnailImageUrl string
}

type UpdateSeratUseCaser interface {
	Exec(ctx context.Context, id uuid.UUID, params *UpdateSeratParams) (*entity.Serat, error)
}

type UpdateSeratUseCase struct {
	SeratRepository postgres.SeratRepositorier
}

func NewUpdateSeratUseCase(seratRepository postgres.SeratRepositorier) UpdateSeratUseCaser {
	return &UpdateSeratUseCase{
		SeratRepository: seratRepository,
	}
}

func (u *UpdateSeratUseCase) Exec(ctx context.Context, id uuid.UUID, params *UpdateSeratParams) (*entity.Serat, error) {
	return u.SeratRepository.UpdateSerat(ctx, id, &entity.Serat{
		Title:             params.Title,
		Description:       params.Description,
		Body:              params.Body,
		CoverImageUrl:     params.CoverImageUrl,
		ThumbnailImageUrl: params.ThumbnailImageUrl,
	})
}
