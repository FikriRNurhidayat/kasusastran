package entity

import (
	"github.com/google/uuid"
)

type Serat struct {
	ID                uuid.UUID
	Title             string
	Description       string
	Content           string
	CoverImageUrl     string
	ThumbnailImageUrl string
}
