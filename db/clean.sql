DO $$
  DECLARE statements CURSOR FOR
      SELECT tablename FROM pg_tables
      WHERE schemaname = 'public';
BEGIN
    FOR stmt IN statements LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
    END LOOP;
END$$;
