-- name: GetWulangan :one
SELECT
  wulangans.id,
  wulangans.title,
  wulangans.description,
  wulangans.cover_image_url,
  wulangans.thumbnail_image_url,
  wulangans.created_at,
  wulangans.updated_at
FROM wulangans
WHERE wulangans.id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: ListWulangans :many
SELECT
  wulangans.id,
  wulangans.title,
  wulangans.description,
  wulangans.cover_image_url,
  wulangans.thumbnail_image_url,
  wulangans.created_at,
  wulangans.updated_at
FROM wulangans
WHERE deleted_at IS NULL
LIMIT $1
OFFSET $2;

-- name: CountWulangans :one
SELECT COUNT(wulangans.id) FROM wulangans WHERE deleted_at IS NULL;

-- name: DeleteWulangan :exec
UPDATE wulangans SET deleted_at = NOW() WHERE id = $1;

-- name: CreateWulangan :one
INSERT INTO wulangans (
  title,
  description,
  cover_image_url,
  thumbnail_image_url,
  created_at,
  updated_at
)
VALUES (
  @title,
  @description,
  @cover_image_url,
  @thumbnail_image_url,
  NOW(),
  NOW()
)
RETURNING
  id,
  title,
  description,
  cover_image_url,
  thumbnail_image_url,
  created_at,
  updated_at;

-- name: UpdateWulangan :one
UPDATE wulangans
SET
  title = @title,
  description = @description,
  cover_image_url = @cover_image_url,
  thumbnail_image_url = @thumbnail_image_url,
  updated_at = NOW()
WHERE wulangans.id = @id
RETURNING
  id,
  title,
  description,
  cover_image_url,
  thumbnail_image_url,
  created_at,
  updated_at;
