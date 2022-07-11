package svc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type GetSeratService interface {
	Call(ctx context.Context, id uuid.UUID) (*GetSeratResult, error)
}

type GetSeratResult entity.Serat

type getSeratService struct {
	seratRepository   repository.SeratRepository
	seratEventEmitter event.SeratEventEmitter
}

func NewGetSeratService(seratRepository repository.SeratRepository, seratEventEmitter event.SeratEventEmitter) GetSeratService {
	return &getSeratService{
		seratRepository:   seratRepository,
		seratEventEmitter: seratEventEmitter,
	}
}

func (s *getSeratService) Call(ctx context.Context, id uuid.UUID) (*GetSeratResult, error) {
	serat, err := s.seratRepository.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	res := &GetSeratResult{
		ID:                serat.ID,
		Title:             serat.Title,
		Description:       serat.Description,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	err = s.seratEventEmitter.EmitRetrievedEvent(&event.Message{
		Actor:   &event.Actor{},
		Payload: res,
	})

	return res, err
}
