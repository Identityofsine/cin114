-- +goose Up

CREATE TABLE IF NOT EXISTS event_locations (
  event_location_id SERIAL PRIMARY KEY,
  event_id INTEGER NOT NULL,
  location_name VARCHAR(255) NOT NULL,
  location_description TEXT,
  longitude DECIMAL(9, 6) NOT NULL,
  latitude DECIMAL(9, 6) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  -- Foreign key constraints
  CONSTRAINT fk_event_locations_event_id 
    FOREIGN KEY (event_id) 
    REFERENCES events(event_id) 
    ON DELETE CASCADE
);

INSERT INTO event_locations (
  event_id,
  location_name,
  location_description,
  longitude,
  latitude,
  created_at,
  updated_at
) VALUES (
  1000, -- Assuming this is the event_id from the previous migration
  'Screening at the Sayville Theatre',
  'Located at 103 Railroad Ave in downtown Sayville, the Theatre greets guests with a classic neon-lit marquee and old-school ticket window.',
  40.73593, -- Example longitude
  -73.08206, -- Example latitude
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE IF EXISTS event_locations;
