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
	seratRepository   repository.SeratRepository
	seratEventEmitter event.SeratEventEmitter
}

func NewCreateSeratService(seratRepository repository.SeratRepository, seratEventEmitter event.SeratEventEmitter) CreateSeratService {
	return &createSeratService{
		seratRepository:   seratRepository,
		seratEventEmitter: seratEventEmitter,
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
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	res := &CreateSeratResult{
		ID:                serat.ID,
		Title:             serat.Title,
		Description:       serat.Description,
		Content:           serat.Content,
		CoverImageUrl:     serat.CoverImageUrl,
		ThumbnailImageUrl: serat.ThumbnailImageUrl,
	}

	if err := s.seratEventEmitter.EmitCreatedEvent(&message.Serat{
		ID:        uuid.New(),
		Kind:      event.SERAT_CREATED_TOPIC,
		CreatedAt: time.Now(),
		Actor:     &message.Actor{},
		Payload:   serat,
	}); err != nil {
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	return res, nil
}
