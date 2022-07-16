package svc

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	"github.com/google/uuid"
)

type GetWulanganService interface {
	Call(ctx context.Context, id uuid.UUID) (*GetWulanganResult, error)
}

type GetWulanganResult entity.Wulangan

type getWulanganService struct {
	wulanganRepository   repository.WulanganRepository
	wulanganEventEmitter event.WulanganEventEmitter
}

func (s *getWulanganService) Call(ctx context.Context, id uuid.UUID) (*GetWulanganResult, error) {
	wulangan, err := s.wulanganRepository.Get(ctx, id)

	if err != nil {
		return nil, trouble.WULANGAN_NOT_FOUND
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

	if err := s.wulanganEventEmitter.EmitRetrievedEvent(&message.Wulangan{
		ID:        uuid.New(),
		Kind:      event.WULANGAN_RETRIEVED_TOPIC,
		CreatedAt: time.Now(),
		Actor:     &message.Actor{},
		Payload:   wulangan,
	}); err != nil {
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	return res, nil
}

func NewGetWulanganService(wulanganRepository repository.WulanganRepository, wulanganEventEmitter event.WulanganEventEmitter) GetWulanganService {
	return &getWulanganService{
		wulanganRepository:   wulanganRepository,
		wulanganEventEmitter: wulanganEventEmitter,
	}
}
