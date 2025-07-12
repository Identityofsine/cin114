-- +goose Up

-- Update image URLs to begin with forward slash
UPDATE event_images 
SET image_url = '/' || image_url 
WHERE image_url NOT LIKE '/%' 
AND event_id = 1000;

-- +goose Down

-- Revert image URLs to remove forward slash
UPDATE event_images 
SET image_url = SUBSTRING(image_url FROM 2) 
WHERE image_url LIKE '/%' 
AND event_id = 1000;
