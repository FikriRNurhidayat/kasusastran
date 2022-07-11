package svc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
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
	err := s.seratRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	err = s.seratEventEmitter.EmitDeletedEvent(&event.Message{
		Actor:   &event.Actor{},
		Payload: id.String(),
	})

	return err
}
