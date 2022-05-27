CREATE TABLE IF NOT EXISTS wulangan_babs (
  id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title       VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  kaca_count  INT          NOT NULL DEFAULT 0,
  wulangan_id UUID NOT NULL REFERENCES wulangans(id),
  created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at  TIMESTAMP
);

COMMENT ON TABLE wulangan_babs IS 'Store every bab on a wulangan (course) on kasusastran.io'; 
COMMENT ON COLUMN wulangan_babs.title IS 'The title on a Bab of a Wulangan.'; 
COMMENT ON COLUMN wulangan_babs.description IS 'The summary of a Bab in a Wulangan.'; 
COMMENT ON COLUMN wulangan_babs.kaca_count IS 'Total kaca on a Bab of Wulangan'; 
COMMENT ON COLUMN wulangan_babs.wulangan_id IS 'The id of wulangan.'; 
