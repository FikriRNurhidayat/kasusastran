package svc

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type ListSeratsService interface {
	Call(context.Context, *ListSeratsParams) (*ListSeratsResult, error)
}

type ListSeratsResult struct {
	Serats     []*entity.Serat
	Pagination *manager.Pagination
}

type ListSeratsParams struct {
	Pagination *manager.Pagination
}

type listSeratsService struct {
	seratRepository   repository.SeratRepository
	seratEventEmitter event.SeratEventEmitter
}

func NewListSeratsService(seratRepository repository.SeratRepository, seratEventEmitter event.SeratEventEmitter) ListSeratsService {
	return &listSeratsService{
		seratRepository:   seratRepository,
		seratEventEmitter: seratEventEmitter,
	}
}

func (s *listSeratsService) Call(ctx context.Context, params *ListSeratsParams) (*ListSeratsResult, error) {
	serats, count, err := s.seratRepository.List(ctx, params.Pagination.ToListQuery())

	if err != nil {
		return nil, err
	}

	res := &ListSeratsResult{
		Serats:     serats,
		Pagination: params.Pagination.WithTotal(count),
	}

	err = s.seratEventEmitter.EmitListedEvent(&message.Serats{
		ID:        uuid.New(),
		Kind:      event.SERAT_LISTED_TOPIC,
		CreatedAt: time.Now(),
		Actor:     &message.Actor{},
		Payload:   serats,
	})

	return res, err
}
