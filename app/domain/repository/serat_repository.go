package repository

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/google/uuid"
)

type SeratRepository interface {
	Create(ctx context.Context, iserat *entity.Serat) (serat *entity.Serat, err error)
	Get(ctx context.Context, id uuid.UUID) (serat *entity.Serat, err error)
	Update(ctx context.Context, id uuid.UUID, userat *entity.Serat) (serat *entity.Serat, err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
	List(ctx context.Context, query *ListQuery) (serats []*entity.Serat, count uint32, err error)
}
