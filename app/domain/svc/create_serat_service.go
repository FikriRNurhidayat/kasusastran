package svc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
)

type CreateSeratParams struct {
	Title             string
	Description       string
	Content           string
	CoverImageUrl     string
	ThumbnailImageUrl string
}

type CreateSeratResult entity.Serat

type CreateSeratService interface {
	Call(ctx context.Context, params *CreateSeratParams) (*CreateSeratResult, error)
}

type createSeratService struct {
	seratRepository repository.SeratRepository
}

func NewCreateSeratService(seratRepository repository.SeratRepository) CreateSeratService {
	return &createSeratService{
		seratRepository: seratRepository,
	}
}

func (s *createSeratService) Call(ctx context.Context, params *CreateSeratParams) (*CreateSeratResult, error) {
	serat, err := s.seratRepository.Create(ctx, &entity.Serat{
		Title:             params.Title,
		Description:       params.Description,
		Content:           params.Content,
		CoverImageUrl:     params.CoverImageUrl,
		ThumbnailImageUrl: params.ThumbnailImageUrl,
	})

	if err != nil {
		return nil, err
	}

	res := &CreateSeratResult{
		ID:                serat.ID,
		Title:             serat.Title,
		Description:       serat.Description,
		Content:           serat.Content,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	return res, nil
}
