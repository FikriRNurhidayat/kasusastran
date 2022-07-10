package svc

import (
	"context"
	"errors"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type GetWulanganService interface {
	Call(ctx context.Context, id uuid.UUID) (*GetWulanganResult, error)
}

type GetWulanganResult entity.Wulangan

type getWulanganService struct {
	wulanganRepository repository.WulanganRepository
}

func (s *getWulanganService) Call(ctx context.Context, id uuid.UUID) (*GetWulanganResult, error) {
	wulangan, err := s.wulanganRepository.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if wulangan == nil {
		return nil, errors.New("wulangan not found")
	}

	res := &GetWulanganResult{
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

func NewGetWulanganService(wulanganRepository repository.WulanganRepository) GetWulanganService {
	return &getWulanganService{
		wulanganRepository: wulanganRepository,
	}
}
