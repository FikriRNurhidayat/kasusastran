package svc

import (
	"context"
	"errors"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type GetWulanganService interface {
	Call(ctx context.Context, id uuid.UUID) (*entity.Wulangan, error)
}

type getWulanganService struct {
	wulanganRepository repository.WulanganRepository
}

// Call implements GetWulanganService
func (s *getWulanganService) Call(ctx context.Context, id uuid.UUID) (*entity.Wulangan, error) {
	w, err := s.wulanganRepository.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if w == nil {
		return nil, errors.New("wulangan not found")
	}

	return w, nil
}

func NewGetWulanganService(wulanganRepository repository.WulanganRepository) GetWulanganService {
	return &getWulanganService{
		wulanganRepository: wulanganRepository,
	}
}
