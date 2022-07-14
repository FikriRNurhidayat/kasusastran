CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255),
  encrypted_password TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS users_email_deleted_at_key ON users (email) WHERE deleted_at IS NULL;

COMMENT ON TABLE users IS 'Store every users';
COMMENT ON COLUMN users.id IS 'The id of user';
COMMENT ON COLUMN users.name IS 'The name of user';
COMMENT ON COLUMN users.email IS 'The email of user';
COMMENT ON COLUMN users.encrypted_password IS 'The encrypted password of a user';
