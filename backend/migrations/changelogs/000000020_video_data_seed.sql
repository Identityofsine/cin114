-- +goose Up

-- Insert cast members first
INSERT INTO cast_members (name, created_at, updated_at) VALUES
  ('Kai Helenius', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Antonio Venticinque', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Shane Keeley', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Erin Hennig', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Noah Fields', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Kristopher King', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Louis Clarke', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Meredith Reed', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Gabriel Patrascu', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Sebastian Caldwell', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Kenrhon Anthony', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Cayson Rhodes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Makayla Russo', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert Films data
INSERT INTO videos (
  title,
  description,
  weight,
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
    'Interviewing John Ford',
    'A young interviewer gets the opportunity to interview famed film director John Ford.',
    0,
    false,
    '',
    'Ask the right questions...',
    '/film/john-ford/boxart.png',
    '/film/john-ford/video.mp4',
    '/catalog/john-ford',
    '02/15/2024',
    '/film/john-ford/image1.png',
    'film',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    'UNO',
    'Three friends play a game of UNO',
    0,
    false,
    'DRAW...',
    'Three friends get togehter to play a game of UNO until things turn sour.',
    '/film/uno/boxart.png',
    '/film/uno/video.mp4',
    '/catalog/uno',
    '2024',
    '/film/uno/image1.png',
    'film',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    'Just Give Me The Night',
    '',
    0,
    false,
    'It''s a simple task...',
    'A man ventures into the night to buy food for for his cat, but complications soon arise.',
    '/film/just-give-me-tonight/boxart.png',
    '/film/just-give-me-tonight/video.mp4',
    '/catalog/just-give-me-the-night',
    '2024',
    '/film/just-give-me-tonight/image1.png',
    'film',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    'The Imaginary Rules of Engagement',
    '',
    0,
    false,
    '',
    'Two kids are outside playing with toys while using their imaginations. What''s the worst that can happen?',
    '/film/the-rules-of-engagement/boxart.png',
    '/film/the-rules-of-engagement/video.mp4',
    '/catalog/imaginary-rules-of-engagement',
    '2024',
    '/film/the-rules-of-engagement/image1.png',
    'film',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    '16 ROUND Wood Fired Pizza',
    'A promotional video for 16 ROUND Wood Fired Pizza',
    0,
    true,
    '',
    'A promotional video for 16 ROUND Wood Fired Pizza',
    '/video/16round/image1.png',
    '/video/16round/video.mp4',
    '#',
    '2024',
    '/video/16round/image1.png',
    'promotional',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  );

-- Insert video links (YouTube links)
INSERT INTO video_links (video_id, link_type, link_url, created_at, updated_at)
SELECT 
  v.video_id,
  'youtube',
  'https://www.youtube.com/watch?v=gZaosS7-l5w&ab_channel=CIN114',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v 
WHERE v.title = 'Interviewing John Ford';

INSERT INTO video_links (video_id, link_type, link_url, created_at, updated_at)
SELECT 
  v.video_id,
  'youtube',
  'https://www.youtube.com/watch?v=usKIKXzoXyM&ab_channel=CIN114',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v 
WHERE v.title = 'UNO';

INSERT INTO video_links (video_id, link_type, link_url, created_at, updated_at)
SELECT 
  v.video_id,
  'youtube',
  'https://www.youtube.com/watch?v=Gk0a63sfaF0&ab_channel=CIN114',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v 
WHERE v.title = 'The Imaginary Rules of Engagement';

-- Insert video credits for "Interviewing John Ford"
INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  1,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Interviewing John Ford' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  2,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Interviewing John Ford' AND cm.name = 'Antonio Venticinque';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'director',
  3,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Interviewing John Ford' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'writer',
  4,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Interviewing John Ford' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'cinematographer',
  5,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Interviewing John Ford' AND cm.name = 'Shane Keeley';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  6,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Interviewing John Ford' AND cm.name = 'Erin Hennig';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  7,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Interviewing John Ford' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  8,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Interviewing John Ford' AND cm.name = 'Noah Fields';

-- Insert video credits for "UNO"
INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  1,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'UNO' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  2,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'UNO' AND cm.name = 'Shane Keeley';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  3,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'UNO' AND cm.name = 'Noah Fields';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'director',
  4,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'UNO' AND cm.name = 'Kristopher King';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'writer',
  5,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'UNO' AND cm.name = 'Kristopher King';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'cinematographer',
  6,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'UNO' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  7,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'UNO' AND cm.name = 'Kristopher King';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  8,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'UNO' AND cm.name = 'Kai Luckey';

-- Insert video credits for "Just Give Me The Night"
INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  1,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Just Give Me The Night' AND cm.name = 'Kristopher King';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  2,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Just Give Me The Night' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'director',
  3,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Just Give Me The Night' AND cm.name = 'Shane Keeley';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'writer',
  4,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Just Give Me The Night' AND cm.name = 'Shane Keeley';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'cinematographer',
  5,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Just Give Me The Night' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  6,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Just Give Me The Night' AND cm.name = 'Shane Keeley';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  7,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'Just Give Me The Night' AND cm.name = 'Kai Luckey';

-- Insert video credits for "The Imaginary Rules of Engagement"
INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  1,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Louis Clarke';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  2,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Meredith Reed';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  3,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Gabriel Patrascu';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  4,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Sebastian Caldwell';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  5,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Kenrhon Anthony';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  6,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Cayson Rhodes';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'starring',
  7,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Makayla Russo';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'director',
  8,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Noah Fields';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'a_camera',
  9,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Kristopher King';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'cinematographer',
  10,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Kai Luckey';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  11,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Sebastian Caldwell';

INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order, created_at, updated_at)
SELECT 
  v.video_id,
  cm.cast_member_id,
  'producer',
  12,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
FROM videos v, cast_members cm
WHERE v.title = 'The Imaginary Rules of Engagement' AND cm.name = 'Kai Luckey';

-- Update style_json for mv3 (music video with background position)
UPDATE videos 
SET style_json = '{"backgroundPosition": "25% center"}'::jsonb
WHERE title = '' AND img = '/mv/mv3/image1.png';

-- +goose Down

-- Note: This is a data migration, so the Down migration would need to delete all the inserted data
-- For safety, we'll leave this empty to prevent accidental data deletion 