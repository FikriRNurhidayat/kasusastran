package svc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
)

type CreateWulanganParams struct {
	Title             string
	Description       string
	CoverImageUrl     string
	ThumbnailImageUrl string
}

type CreateWulanganResult entity.Wulangan

type CreateWulanganService interface {
	Call(context.Context, *CreateWulanganParams) (*CreateWulanganResult, error)
}

type createWulanganService struct {
	wulanganRepository repository.WulanganRepository
}

func NewCreateWulanganService(wulanganRepository repository.WulanganRepository) CreateWulanganService {
	return &createWulanganService{
		wulanganRepository: wulanganRepository,
	}
}

func (s *createWulanganService) Call(ctx context.Context, params *CreateWulanganParams) (*CreateWulanganResult, error) {
	wulangan := &entity.Wulangan{
		Title:             params.Title,
		Description:       params.Description,
		CoverImageUrl:     params.CoverImageUrl,
		ThumbnailImageUrl: params.ThumbnailImageUrl,
	}

	wulangan, err := s.wulanganRepository.Create(ctx, wulangan)

	if err != nil {
		return nil, err
	}

	res := &CreateWulanganResult{
		ID:                wulangan.ID,
		Title:             wulangan.Title,
		Description:       wulangan.Description,
		CoverImageUrl:     wulangan.CoverImageUrl,
		ThumbnailImageUrl: wulangan.ThumbnailImageUrl,
		CreatedAt:         wulangan.CreatedAt,
		UpdatedAt:         wulangan.UpdatedAt,
	}

	return res, nil
}
