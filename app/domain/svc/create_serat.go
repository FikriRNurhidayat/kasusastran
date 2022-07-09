package svc

import (
	"context"

	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/entity"
	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/repository"
)

type CreateSeratParams struct {
	Title             string
	Description       string
	Content           string
	CoverImageUrl     string
	ThumbnailImageUrl string
}

type CreateSeratService interface {
	Exec(ctx context.Context, params *CreateSeratParams) (*entity.Serat, error)
}

type createSeratService struct {
	seratRepository repository.SeratRepository
}

func NewCreateSeratService(seratRepository repository.SeratRepository) CreateSeratService {
	return &createSeratService{
		seratRepository: seratRepository,
	}
}

func (s *createSeratService) Exec(ctx context.Context, params *CreateSeratParams) (*entity.Serat, error) {
	return s.seratRepository.Create(ctx, &entity.Serat{
		Title:             params.Title,
		Description:       params.Description,
		Content:           params.Content,
		CoverImageUrl:     params.CoverImageUrl,
		ThumbnailImageUrl: params.ThumbnailImageUrl,
	})
}
