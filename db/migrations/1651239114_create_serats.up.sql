CREATE TABLE IF NOT EXISTS serats (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  cover_image_url VARCHAR(255) NOT NULL,
  thumbnail_image_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

COMMENT ON COLUMN serats.title IS 'The title of a Serat.'; 
COMMENT ON COLUMN serats.description IS 'The summary of a Serat.'; 
COMMENT ON COLUMN serats.cover_image_url IS 'The cover image URL of a Serat, shown on detail page.'; 
COMMENT ON COLUMN serats.thumbnail_image_url IS 'The thumbnail image URL of a Serat, shown when retrieved as list.'; 
