-- +goose Up

UPDATE events set expiration_date = '2025-08-17 15:00:00' 
WHERE event_id = 1000;
