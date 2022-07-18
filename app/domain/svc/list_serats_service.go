package svc

import (
	"context"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	"github.com/google/uuid"
	"google.golang.org/grpc/grpclog"
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
	logger            grpclog.LoggerV2
	seratRepository   repository.SeratRepository
	seratEventEmitter event.SeratEventEmitter
}

func NewListSeratsService(logger grpclog.LoggerV2, seratRepository repository.SeratRepository, seratEventEmitter event.SeratEventEmitter) ListSeratsService {
	return &listSeratsService{
		logger:            logger,
		seratRepository:   seratRepository,
		seratEventEmitter: seratEventEmitter,
	}
}

func (s *listSeratsService) Call(ctx context.Context, params *ListSeratsParams) (*ListSeratsResult, error) {
	serats, count, err := s.seratRepository.List(ctx, params.Pagination.ToListQuery())

	if err != nil {
		s.logger.Error(err)
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	res := &ListSeratsResult{
		Serats:     serats,
		Pagination: params.Pagination.WithTotal(count),
	}

	if err := s.seratEventEmitter.EmitListedEvent(&message.Serats{
		ID:        uuid.New(),
		Kind:      event.SERAT_LISTED_TOPIC,
		CreatedAt: time.Now(),
		Actor:     &message.Actor{},
		Payload:   serats,
	}); err != nil {
		s.logger.Error(err)
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	return res, nil
}
