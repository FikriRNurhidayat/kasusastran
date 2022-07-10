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

type CreateWulanganService interface {
	Call(context.Context, *CreateWulanganParams) (*entity.Wulangan, error)
}

type createWulanganService struct {
	wulanganRepository repository.WulanganRepository
}

func NewCreateWulanganService(wulanganRepository repository.WulanganRepository) CreateWulanganService {
	return &createWulanganService{
		wulanganRepository: wulanganRepository,
	}
}

func (s *createWulanganService) Call(ctx context.Context, params *CreateWulanganParams) (*entity.Wulangan, error) {
	w := &entity.Wulangan{
		Title:             params.Title,
		Description:       params.Description,
		CoverImageUrl:     params.CoverImageUrl,
		ThumbnailImageUrl: params.ThumbnailImageUrl,
	}

	w, err := s.wulanganRepository.Create(ctx, w)

	if err != nil {
		return nil, err
	}

	return w, nil
}
