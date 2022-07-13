ALTER TABLE wulangan_bab_kacas ADD COLUMN idx INT DEFAULT 0;
ALTER TABLE wulangan_bab_kacas ADD COLUMN gidx INT DEFAULT 0;

COMMENT ON COLUMN wulangan_bab_kacas.idx IS 'The kaca order inside bab.';
COMMENT ON COLUMN wulangan_bab_kacas.gidx IS 'The kaca order on wulangan perspective';
