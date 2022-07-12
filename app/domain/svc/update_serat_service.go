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

type UpdateSeratParams struct {
	Title             string
	Description       string
	Content           string
	CoverImageUrl     string
	ThumbnailImageUrl string
}

type UpdateSeratResult entity.Serat

type UpdateSeratService interface {
	Call(ctx context.Context, id uuid.UUID, params *UpdateSeratParams) (*UpdateSeratResult, error)
}

type updateSeratService struct {
	seratRepository   repository.SeratRepository
	seratEventEmitter event.SeratEventEmitter
}

func NewUpdateSeratService(seratRepository repository.SeratRepository, seratEventEmitter event.SeratEventEmitter) UpdateSeratService {
	return &updateSeratService{
		seratRepository:   seratRepository,
		seratEventEmitter: seratEventEmitter,
	}
}

func (s *updateSeratService) Call(ctx context.Context, id uuid.UUID, params *UpdateSeratParams) (*UpdateSeratResult, error) {
	serat, err := s.seratRepository.Update(ctx, id, &entity.Serat{
		Title:             params.Title,
		Description:       params.Description,
		Content:           params.Content,
		CoverImageUrl:     params.CoverImageUrl,
		ThumbnailImageUrl: params.ThumbnailImageUrl,
	})

	if err != nil {
		return nil, err
	}

	res := &UpdateSeratResult{
		ID:                serat.ID,
		Title:             serat.Title,
		Description:       serat.Description,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	err = s.seratEventEmitter.EmitUpdatedEvent(&message.Serat{
		ID:        uuid.New(),
		Kind:      event.SERAT_UPDATED_TOPIC,
		CreatedAt: time.Now(),
		Actor:     &message.Actor{},
		Payload:   serat,
	})

	return res, err
}
