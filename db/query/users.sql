-- name: CreateUser :one
INSERT INTO users (
       name,
       email,
       encrypted_password,
       created_at,
       updated_at
)
VALUES (
  @name,
  @email,
  @encrypted_password,
  NOW(),
  NOW()
)
RETURNING
    id,
    name,
    email,
    encrypted_password,
    created_at,
    updated_at;


-- name: DoesUserEmailExist :one
SELECT EXISTS (
       SELECT email FROM users
       WHERE users.deleted_at IS NULL
       AND users.email = $1
);

-- name: GetUserByEmail :one
SELECT id, name, email, encrypted_password, created_at, updated_at
FROM users
WHERE email = $1
AND deleted_at IS NULL;
