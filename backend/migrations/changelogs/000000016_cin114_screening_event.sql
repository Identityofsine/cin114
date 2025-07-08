-- +goose Up

INSERT INTO events (
  event_id,
  description,
  short_description,
  expiration_date,
  created_at,
  updated_at
) VALUES (
  1000,
  'Expensive to Die in America\nRedpine\nI Wish My Bones Were Unidentifiable',
  'Triple feature screening: 3:00pm August 17th',
  '2024-08-18 23:59:59',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
);

-- +goose Down

DELETE FROM events 
WHERE event_id = 1000; 
