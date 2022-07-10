package postgres

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type PostgresWulanganRepository struct {
	db query.Querier
}

// Delete implements repository.WulanganRepository
func (*PostgresWulanganRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// Get implements repository.WulanganRepository
func (*PostgresWulanganRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Wulangan, error) {
	panic("unimplemented")
}

// List implements repository.WulanganRepository
func (*PostgresWulanganRepository) List(ctx context.Context) ([]*entity.Wulangan, error) {
	panic("unimplemented")
}

// Update implements repository.WulanganRepository
func (*PostgresWulanganRepository) Update(ctx context.Context, id uuid.UUID, w *entity.Wulangan) (*entity.Wulangan, error) {
	panic("unimplemented")
}

func (r *PostgresWulanganRepository) Create(ctx context.Context, w *entity.Wulangan) (*entity.Wulangan, error) {
	row, err := r.db.CreateWulangan(ctx, &query.CreateWulanganParams{
		Title:             w.Title,
		Description:       w.Description,
		CoverImageUrl:     w.CoverImageUrl,
		ThumbnailImageUrl: w.ThumbnailImageUrl,
	})

	if err != nil {
		return nil, err
	}

	w.ID = row.ID
	w.Title = row.Title
	w.Description = row.Description
	w.CoverImageUrl = row.CoverImageUrl
	w.ThumbnailImageUrl = row.ThumbnailImageUrl
	w.CreatedAt = row.CreatedAt
	w.UpdatedAt = row.UpdatedAt

	return w, nil
}

func NewPostgresWulanganRepository(db query.Querier) repository.WulanganRepository {
	return &PostgresWulanganRepository{
		db: db,
	}
}
