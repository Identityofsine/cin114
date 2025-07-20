-- +goose Up

ALTER TABLE event_locations ADD column location_address VARCHAR(255) NULL; 

UPDATE event_locations SET location_address = '103 Railroad Ave, Sayville, NY 11782' WHERE event_id = 1000;

-- +goose Down

ALTER TABLE event_locations DROP COLUMN IF EXISTS location_address;
