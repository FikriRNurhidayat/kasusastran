package svc

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type DeleteSeratService interface {
	Call(ctx context.Context, id uuid.UUID) error
}

type deleteSeratService struct {
	seratRepository   repository.SeratRepository
	seratEventEmitter event.SeratEventEmitter
}

func NewDeleteSeratService(seratRepository repository.SeratRepository, seratEventEmitter event.SeratEventEmitter) DeleteSeratService {
	return &deleteSeratService{
		seratRepository:   seratRepository,
		seratEventEmitter: seratEventEmitter,
	}
}

func (s *deleteSeratService) Call(ctx context.Context, id uuid.UUID) error {
	serat, err := s.seratRepository.Get(ctx, id)

	if err != nil {
		return err
	}

	err = s.seratRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	err = s.seratEventEmitter.EmitDeletedEvent(&message.Serat{
		ID:        uuid.New(),
		Kind:      event.SERAT_DELETED_TOPIC,
		CreatedAt: time.Now(),
		Actor:     &message.Actor{},
		Payload:   serat,
	})

	return err
}
