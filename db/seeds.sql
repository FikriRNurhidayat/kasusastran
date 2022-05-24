-- Seeding serats
COPY serats(id, title, description, body, cover_image_url, thumbnail_image_url) FROM '${PWD}/db/data/serats.csv' DELIMITER ',' CSV NULL AS 'NULL';
