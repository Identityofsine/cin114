-- +goose Up

UPDATE event_locations SET location_name = 'Sayville Theatre' WHERE event_id = 1000;

-- +goose Down

UPDATE event_locations SET location_name = 'Screening at the Sayville Theatre' WHERE event_id = 1000;
