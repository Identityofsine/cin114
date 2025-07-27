-- +goose Up

-- Clear all data from stripe_prices table
DELETE FROM stripe_prices;

-- Update longitude and latitude for event_locations
UPDATE event_locations 
SET longitude = 40.737500, latitude = -73.082608 
WHERE event_location_id IS NOT NULL;

-- +goose Down

-- Note: This migration is destructive and cannot be fully reversed
-- The stripe_prices data that was deleted cannot be restored
-- The previous coordinates are also lost and cannot be restored automatically 