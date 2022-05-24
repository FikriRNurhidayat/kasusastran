package postgres

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/google/uuid"
)

type SeratRepositorier interface {
	CreateSerat(ctx context.Context, iserat *entity.Serat) (serat *entity.Serat, err error)
	GetSerat(ctx context.Context, id uuid.UUID) (serat *entity.Serat, err error)
	UpdateSerat(ctx context.Context, id uuid.UUID, userat *entity.Serat) (serat *entity.Serat, err error)
	DeleteSerat(ctx context.Context, id uuid.UUID) (err error)
	ListSerats(ctx context.Context, pagination *entity.Pagination) (serats []entity.Serat, count uint32, err error)
}

type SeratRepository struct {
	db query.Querier
}

func NewSeratRepository(db query.Querier) SeratRepositorier {
	return &SeratRepository{
		db: db,
	}
}

func (r *SeratRepository) UpdateSerat(ctx context.Context, id uuid.UUID, userat *entity.Serat) (serat *entity.Serat, err error) {
	row, err := r.db.UpdateSerat(ctx, &query.UpdateSeratParams{
		ID:                id,
		Title:             userat.Title,
		Description:       userat.Description,
		Body:              userat.Body,
		CoverImageUrl:     userat.CoverImageUrl,
		ThumbnailImageUrl: userat.ThumbnailImageUrl,
	})

	if err != nil {
		return
	}

	serat = &entity.Serat{
		ID:                row.ID,
		Title:             row.Title,
		Body:              row.Body,
		Description:       row.Description,
		CoverImageUrl:     row.CoverImageUrl,
		ThumbnailImageUrl: row.ThumbnailImageUrl,
	}

	return
}

func (r *SeratRepository) CreateSerat(ctx context.Context, iserat *entity.Serat) (serat *entity.Serat, err error) {
	row, err := r.db.CreateSerat(ctx, &query.CreateSeratParams{
		Title:             iserat.Title,
		Description:       iserat.Description,
		Body:              iserat.Body,
		CoverImageUrl:     iserat.CoverImageUrl,
		ThumbnailImageUrl: iserat.ThumbnailImageUrl,
	})

	if err != nil {
		return
	}

	serat = &entity.Serat{
		ID:                row.ID,
		Title:             row.Title,
		Body:              row.Body,
		Description:       row.Description,
		CoverImageUrl:     row.CoverImageUrl,
		ThumbnailImageUrl: row.ThumbnailImageUrl,
	}

	return
}

func (r *SeratRepository) GetSerat(ctx context.Context, id uuid.UUID) (serat *entity.Serat, err error) {
	row, err := r.db.GetSerat(ctx, id)

	if err != nil {
		return
	}

	serat = &entity.Serat{
		ID:                row.ID,
		Title:             row.Title,
		Body:              row.Body,
		Description:       row.Description,
		CoverImageUrl:     row.CoverImageUrl,
		ThumbnailImageUrl: row.ThumbnailImageUrl,
	}

	return
}

func (r *SeratRepository) DeleteSerat(ctx context.Context, id uuid.UUID) error {
	return r.db.DeleteSerat(ctx, id)
}

func (r *SeratRepository) ListSerats(ctx context.Context, pagination *entity.Pagination) (serats []entity.Serat, count uint32, err error) {
	count = 0

	params := &query.ListSeratsParams{
		Limit:  int32(pagination.GetLimit()),
		Offset: int32(pagination.GetOffset()),
	}

	rows, err := r.db.ListSerats(ctx, params)

	if err != nil {
		return
	}

	for _, row := range rows {
		serat := entity.Serat{
			ID:                row.ID,
			Title:             row.Title,
			Description:       row.Description,
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
