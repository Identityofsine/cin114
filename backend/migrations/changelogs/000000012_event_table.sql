-- +goose Up

CREATE TABLE IF NOT EXISTS events (
  event_id SERIAL PRIMARY KEY,
  description TEXT NOT NULL,
  short_description VARCHAR(255),
  expiration_date TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Set the starting value for event_id to 1000
ALTER SEQUENCE events_event_id_seq RESTART WITH 1000;

-- Add indexes for common queries
CREATE INDEX idx_events_expiration_date ON events(expiration_date);
CREATE INDEX idx_events_created_at ON events(created_at);

-- +goose Down

DROP TABLE IF EXISTS events; 