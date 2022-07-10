package entity

import (
	"time"

	"github.com/google/uuid"
)

type Wulangan struct {
	ID                uuid.UUID
	Title             string
	Description       string
	CoverImageUrl     string
	ThumbnailImageUrl string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
