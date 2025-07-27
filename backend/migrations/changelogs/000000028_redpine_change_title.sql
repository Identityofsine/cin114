-- +goose Up

UPDATE events SET short_description = 'CIN114 Triple Feature Screening' WHERE event_id = 1000;
