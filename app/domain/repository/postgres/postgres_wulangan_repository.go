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
func (r *PostgresWulanganRepository) Get(ctx context.Context, id uuid.UUID) (wulangan *entity.Wulangan, err error) {
	row, err := r.db.GetWulangan(ctx, id)

	if err != nil {
		return nil, err
	}

	wulangan = &entity.Wulangan{
		ID:                row.ID,
		Title:             row.Title,
		Description:       row.Description,
		CoverImageUrl:     row.CoverImageUrl,
		ThumbnailImageUrl: row.ThumbnailImageUrl,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
	}

	return wulangan, nil
}

// List implements repository.WulanganRepository
func (*PostgresWulanganRepository) List(ctx context.Context) ([]*entity.Wulangan, error) {
	panic("unimplemented")
}

// Update implements repository.WulanganRepository
func (*PostgresWulanganRepository) Update(ctx context.Context, id uuid.UUID, w *entity.Wulangan) (*entity.Wulangan, error) {
	panic("unimplemented")
}

func (r *PostgresWulanganRepository) Create(ctx context.Context, wulangan *entity.Wulangan) (*entity.Wulangan, error) {
	row, err := r.db.CreateWulangan(ctx, &query.CreateWulanganParams{
		Title:             wulangan.Title,
		Description:       wulangan.Description,
		CoverImageUrl:     wulangan.CoverImageUrl,
		ThumbnailImageUrl: wulangan.ThumbnailImageUrl,
	})

	if err != nil {
		return nil, err
	}

	wulangan.ID = row.ID
	wulangan.Title = row.Title
	wulangan.Description = row.Description
	wulangan.CoverImageUrl = row.CoverImageUrl
	wulangan.ThumbnailImageUrl = row.ThumbnailImageUrl
	wulangan.CreatedAt = row.CreatedAt
	wulangan.UpdatedAt = row.UpdatedAt

	return wulangan, nil
}

func NewPostgresWulanganRepository(db query.Querier) repository.WulanganRepository {
	return &PostgresWulanganRepository{
		db: db,
	}
}
