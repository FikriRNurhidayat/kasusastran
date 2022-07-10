package repository

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/google/uuid"
)

type WulanganRepository interface {
	Create(ctx context.Context, w *entity.Wulangan) (*entity.Wulangan, error)
	Get(ctx context.Context, id uuid.UUID) (*entity.Wulangan, error)
	List(ctx context.Context) ([]*entity.Wulangan, error)
	Update(ctx context.Context, id uuid.UUID, w *entity.Wulangan) (*entity.Wulangan, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
