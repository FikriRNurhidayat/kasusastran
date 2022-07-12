package entity

import (
	"github.com/google/uuid"
)

type Serat struct {
	ID                uuid.UUID `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Content           string    `json:"content"`
	CoverImageUrl     string    `json:"cover_image_url"`
	ThumbnailImageUrl string    `json:"thumbnail_image_url"`
}
