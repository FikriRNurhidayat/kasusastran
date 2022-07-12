package svc

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
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
	wulanganRepository   repository.WulanganRepository
	wulanganEventEmitter event.WulanganEventEmitter
}

func NewCreateWulanganService(wulanganRepository repository.WulanganRepository, wulanganEventEmitter event.WulanganEventEmitter) CreateWulanganService {
	return &createWulanganService{
		wulanganRepository:   wulanganRepository,
		wulanganEventEmitter: wulanganEventEmitter,
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

	err = s.wulanganEventEmitter.EmitCreatedEvent(&message.Wulangan{
		ID:        uuid.New(),
		Kind:      event.WULANGAN_CREATED_TOPIC,
		CreatedAt: time.Now(),
		Actor:     &message.Actor{},
		Payload:   wulangan,
	})

	return res, err
}
