ALTER TABLE wulangan_babs ADD COLUMN idx INT DEFAULT 0;
COMMENT ON COLUMN wulangan_babs.idx IS 'The Bab order';
