package svc

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
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

	err = s.seratEventEmitter.EmitListedEvent(&event.Message{
		Actor:   &event.Actor{},
		Payload: serats,
	})

	return res, err
}
