package postgres

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type PostgresSeratRepository struct {
	db query.Querier
}

func NewPostgresSeratRepository(db query.Querier) repository.SeratRepository {
	return &PostgresSeratRepository{
		db: db,
	}
}

func (r *PostgresSeratRepository) Update(ctx context.Context, id uuid.UUID, userat *entity.Serat) (serat *entity.Serat, err error) {
	row, err := r.db.UpdateSerat(ctx, &query.UpdateSeratParams{
		ID:                id,
		Title:             userat.Title,
		Description:       userat.Description,
		Content:           userat.Content,
		CoverImageUrl:     userat.CoverImageUrl,
		ThumbnailImageUrl: userat.ThumbnailImageUrl,
	})

	if err != nil {
		return
	}

	serat = &entity.Serat{
		ID:                row.ID,
		Title:             row.Title,
		Content:           row.Content,
		Description:       row.Description,
		CoverImageUrl:     row.CoverImageUrl,
		ThumbnailImageUrl: row.ThumbnailImageUrl,
	}

	return
}

func (r *PostgresSeratRepository) Create(ctx context.Context, iserat *entity.Serat) (serat *entity.Serat, err error) {
	row, err := r.db.CreateSerat(ctx, &query.CreateSeratParams{
		Title:             iserat.Title,
		Description:       iserat.Description,
		Content:           iserat.Content,
		CoverImageUrl:     iserat.CoverImageUrl,
		ThumbnailImageUrl: iserat.ThumbnailImageUrl,
	})

	if err != nil {
		return
	}

	serat = &entity.Serat{
		ID:                row.ID,
		Title:             row.Title,
		Content:           row.Content,
		Description:       row.Description,
		CoverImageUrl:     row.CoverImageUrl,
		ThumbnailImageUrl: row.ThumbnailImageUrl,
	}

	return
}

func (r *PostgresSeratRepository) Get(ctx context.Context, id uuid.UUID) (serat *entity.Serat, err error) {
	row, err := r.db.GetSerat(ctx, id)

	if err != nil {
		return
	}

	serat = &entity.Serat{
		ID:                row.ID,
		Title:             row.Title,
		Content:           row.Content,
		Description:       row.Description,
		CoverImageUrl:     row.CoverImageUrl,
		ThumbnailImageUrl: row.ThumbnailImageUrl,
	}

	return
}

func (r *PostgresSeratRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.DeleteSerat(ctx, id)
}

func (r *PostgresSeratRepository) List(ctx context.Context, listQuery *repository.ListQuery) (serats []*entity.Serat, count uint32, err error) {
	count = 0

	params := &query.ListSeratsParams{
		Limit:  int32(listQuery.Limit),
		Offset: int32(listQuery.Offset),
	}

	rows, err := r.db.ListSerats(ctx, params)

	if err != nil {
		return
	}

	for _, row := range rows {
		serat := &entity.Serat{
			ID:                row.ID,
			Title:             row.Title,
			Description:       row.Description,
			Content:           row.Content,
			CoverImageUrl:     row.CoverImageUrl,
			ThumbnailImageUrl: row.ThumbnailImageUrl,
		}
		serats = append(serats, serat)
	}

	c, err := r.db.CountSerats(ctx)

	if err != nil {
		return
	}

	count = uint32(c)

	return
}
