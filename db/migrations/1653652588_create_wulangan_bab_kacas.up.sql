CREATE TABLE IF NOT EXISTS wulangan_bab_kacas (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  wulangan_bab_id UUID NOT NULL REFERENCES wulangan_babs(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

COMMENT ON TABLE wulangan_bab_kacas IS 'Store every kaca of a bab of wulangan (course) on kasusastran.io'; 
COMMENT ON COLUMN wulangan_bab_kacas.title IS 'Kaca title.'; 
COMMENT ON COLUMN wulangan_bab_kacas.content IS 'Kaca content as HTML.'; 
COMMENT ON COLUMN wulangan_bab_kacas.wulangan_bab_id IS 'Bab id of kaca.'; 
