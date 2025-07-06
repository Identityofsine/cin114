-- +goose Up

CREATE TABLE IF NOT EXISTS user_email_whitelist (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users ADD CONSTRAINT fk_user_email_whitelist
  FOREIGN KEY (username)
  REFERENCES user_email_whitelist(email)
  ON DELETE CASCADE;
