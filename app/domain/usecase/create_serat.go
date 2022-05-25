package usecase

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
)

type CreateSeratParams struct {
	Title             string
	Description       string
	Content           string
	CoverImageUrl     string
	ThumbnailImageUrl string
}

type CreateSeratUseCaser interface {
	Exec(ctx context.Context, params *CreateSeratParams) (*entity.Serat, error)
}

type CreateSeratUseCase struct {
	SeratRepository postgres.SeratRepositorier
}

func NewCreateSeratUseCase(seratRepository postgres.SeratRepositorier) CreateSeratUseCaser {
	return &CreateSeratUseCase{
		SeratRepository: seratRepository,
	}
}

func (u *CreateSeratUseCase) Exec(ctx context.Context, params *CreateSeratParams) (*entity.Serat, error) {
	return u.SeratRepository.CreateSerat(ctx, &entity.Serat{
		Title:             params.Title,
		Description:       params.Description,
		Content:           params.Content,
		CoverImageUrl:     params.CoverImageUrl,
		ThumbnailImageUrl: params.ThumbnailImageUrl,
	})
}
