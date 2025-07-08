-- +goose Up

INSERT INTO events (
  description,
  short_description,
  expiration_date,
  created_at,
  updated_at
) VALUES (
  'CIN 114 Screening - Triple Feature: Join us for an exclusive screening of three compelling films: "Expensive to Die in America", "Redpine", and "I Wish My Bones Were Unidentifiable". This special triple feature event showcases powerful storytelling and cinematic artistry. Tickets available for purchase.',
  'Triple feature screening: 3:00pm August 17th',
  '2024-08-18 23:59:59',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
);

-- +goose Down

DELETE FROM events 
WHERE description LIKE 'CIN 114 Screening - Triple Feature:%' 
AND short_description = 'Triple feature screening: 3:00pm August 17th'; 