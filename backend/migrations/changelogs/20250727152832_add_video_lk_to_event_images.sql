-- +goose Up

INSERT into event_image_lks (event_image_lk, event_image_description)
VALUES ('video', 'Video Preview');

INSERT into event_images (event_id, image_url, image_type)
VALUES (1000, '/film/screening-1001/video.mp4', 'video');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DELETE FROM event_images_lk WHERE event_image_lk = 'video';
