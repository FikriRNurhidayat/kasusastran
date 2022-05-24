-- name: GetSerat :one 
SELECT
  serats.id,
  serats.title,
  serats.description,
  serats.body,
  serats.cover_image_url,
  serats.thumbnail_image_url
FROM serats
WHERE serats.id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: ListSerats :many 
SELECT
  serats.id,
  serats.title,
  serats.description,
  serats.cover_image_url,
  serats.thumbnail_image_url
FROM serats
WHERE deleted_at IS NULL
LIMIT $1
OFFSET $2;

-- name: CountSerats :one 
SELECT COUNT(serats.id) FROM serats WHERE deleted_at IS NULL;

-- name: DeleteSerat :exec
UPDATE serats SET deleted_at = NOW() WHERE id = $1;

-- name: CreateSerat :one
INSERT INTO serats (
  title,
  description,
  body,
  cover_image_url,
  thumbnail_image_url,
  created_at,
  updated_at
)
VALUES (
  @title,
  @description,
  @body,
  @cover_image_url,
  @thumbnail_image_url,
  NOW(),
  NOW()
)
RETURNING
  id,
  title,
  description,
  body,
  cover_image_url,
  thumbnail_image_url,
  created_at,
  updated_at;

-- name: UpdateSerat :one
UPDATE serats
SET
  title = @title,
  description = @description,
  body = @body,
  cover_image_url = @cover_image_url,
  thumbnail_image_url = @thumbnail_image_url,
  updated_at = NOW()
WHERE serats.id = @id
RETURNING
  id,
  title,
  body,
  description,
  cover_image_url,
  thumbnail_image_url,
  created_at,
  updated_at;