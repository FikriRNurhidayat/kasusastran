CREATE TABLE IF NOT EXISTS wulangans (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  cover_image_url VARCHAR(255) NOT NULL,
  thumbnail_image_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

COMMENT ON TABLE wulangans IS 'Store every wulangan (course) of kasusastran.io'; 
COMMENT ON COLUMN wulangans.title IS 'The title of a Wulangan.'; 
COMMENT ON COLUMN wulangans.description IS 'The summary of a Wulangan.'; 
COMMENT ON COLUMN wulangans.cover_image_url IS 'The cover image URL of a Wulangan, shown on detail page.'; 
COMMENT ON COLUMN wulangans.thumbnail_image_url IS 'The thumbnail image URL of a Wulangan, shown when retrieved as list.'; 
