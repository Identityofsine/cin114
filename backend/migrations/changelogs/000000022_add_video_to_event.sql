-- +goose Up

ALTER TABLE events ADD COLUMN video_id INTEGER NULL;
-- Add foreign key constraint for video_id in videos TABLE
ALTER TABLE videos ADD CONSTRAINT fk_videos_event_id 
  FOREIGN KEY (video_id) 
  REFERENCES events(event_id) 
  ON DELETE SET NULL;
ALTER TABLE videos ADD COLUMN previewable BOOLEAN DEFAULT TRUE;

-- Add foreign key constraint for video_id in events table
ALTER TABLE events ADD CONSTRAINT fk_events_video_id 
  FOREIGN KEY (video_id) 
  REFERENCES videos(video_id) 
  ON DELETE SET NULL;

-- Insert the video
INSERT INTO videos (
  title,
  description,
  weight,
  previewable,
  use_boxart_as_preview,
  boxart_title,
  boxart_caption,
  boxart_img,
  boxart_video,
  url,
  date,
  img,
  video_type,
  created_at,
  updated_at
) VALUES
  (
    'I Wish My Bones Were Unidentifiable',
    'A young keylime pie maker gets ran over by a car and wakes up in a random hospital in Laos.',
    0,
    false,
    false,
    '',
    '',
    '',
    '',
    '',
    '',
    '',
    'screening',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  );

-- Update the event with the video_id
UPDATE events SET video_id = (
    SELECT video_id FROM videos WHERE title = 'I Wish My Bones Were Unidentifiable'
) WHERE event_id = 1000;

-- +goose Down

-- Remove the foreign key constraint
ALTER TABLE events DROP CONSTRAINT IF EXISTS fk_events_video_id;

-- Remove the video_id column from events table
ALTER TABLE events DROP COLUMN IF EXISTS video_id;

-- Remove the previewable column from videos table
ALTER TABLE videos DROP COLUMN IF EXISTS previewable;

-- Delete the inserted video
DELETE FROM videos WHERE title = 'I Wish My Bones Were Unidentifiable';
