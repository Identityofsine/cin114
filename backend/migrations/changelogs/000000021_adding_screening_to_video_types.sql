-- +goose Up
INSERT INTO video_type_lks (video_type_lk, video_type_description) VALUES
  ('screening', 'Screening of a film');

-- +goose Down

DELETE FROM video_type_lks WHERE video_type_lk = 'screening';
