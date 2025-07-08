-- +goose Up

CREATE TABLE IF NOT EXISTS event_image_lks (
  event_image_lk VARCHAR(50) PRIMARY KEY,
  event_image_description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS event_images (
  event_image_id SERIAL PRIMARY KEY,
  event_id INTEGER NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  image_type VARCHAR(50) NOT NULL, -- e.g., 'poster', 'thumbnail', etc.
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  -- Foreign key constraints
  CONSTRAINT fk_event_images_event_id 
    FOREIGN KEY (event_id) 
    REFERENCES events(event_id) 
    ON DELETE CASCADE,

  CONSTRAINT fk_event_images_event_image_lk 
    FOREIGN KEY (image_type) 
    REFERENCES event_image_lks(event_image_lk) 
    ON DELETE CASCADE
);

INSERT INTO event_image_lks (event_image_lk, event_image_description)
VALUES
  ('poster', 'Poster image for the event'),
  ('poster-mobile', 'Mobile version of the poster image'),
  ('thumbnail', 'Thumbnail image for the event'),
  ('banner', 'Banner image for the event');

INSERT INTO event_images (
  event_id,
  image_url,
  image_type,
  created_at,
  updated_at
) VALUES (
  1000, -- Assuming this is the event_id from the previous migration
  'film/screening-1001/poster.jpg',
  'poster',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
);

INSERT INTO event_images (
  event_id,
  image_url,
  image_type,
  created_at,
  updated_at
) VALUES (
  1000, -- Assuming this is the event_id from the previous migration
  'film/screening-1001/poster-mobile.jpg',
  'poster-mobile',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE IF EXISTS event_images;
DROP TABLE IF EXISTS event_image_lks;
