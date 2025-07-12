-- +goose Up

-- Video type lookup table
CREATE TABLE IF NOT EXISTS video_type_lks (
  video_type_lk VARCHAR(50) PRIMARY KEY,
  video_type_description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Credit role lookup table
CREATE TABLE IF NOT EXISTS credit_role_lks (
  credit_role_lk VARCHAR(50) PRIMARY KEY,
  credit_role_description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Link type lookup table
CREATE TABLE IF NOT EXISTS link_type_lks (
  link_type_lk VARCHAR(50) PRIMARY KEY,
  link_type_description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Cast members table
CREATE TABLE IF NOT EXISTS cast_members (
  cast_member_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Main videos table
CREATE TABLE IF NOT EXISTS videos (
  video_id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  weight INTEGER DEFAULT 0,
  use_boxart_as_preview BOOLEAN DEFAULT false,
  boxart_title VARCHAR(255),
  boxart_caption TEXT,
  boxart_img VARCHAR(255),
  boxart_video VARCHAR(255),
  url VARCHAR(255) NOT NULL,
  date VARCHAR(50),
  img VARCHAR(255),
  style_json JSONB,
  video_type VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  -- Foreign key constraints
  CONSTRAINT fk_videos_video_type 
    FOREIGN KEY (video_type) 
    REFERENCES video_type_lks(video_type_lk) 
    ON DELETE CASCADE
);

-- Video credits table (many-to-many relationship)
CREATE TABLE IF NOT EXISTS video_credits (
  video_credit_id SERIAL PRIMARY KEY,
  video_id INTEGER NOT NULL,
  cast_member_id INTEGER NOT NULL,
  credit_role VARCHAR(50) NOT NULL,
  credit_order INTEGER DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  -- Foreign key constraints
  CONSTRAINT fk_video_credits_video_id 
    FOREIGN KEY (video_id) 
    REFERENCES videos(video_id) 
    ON DELETE CASCADE,

  CONSTRAINT fk_video_credits_cast_member_id
    FOREIGN KEY (cast_member_id)
    REFERENCES cast_members(cast_member_id)
    ON DELETE CASCADE,

  CONSTRAINT fk_video_credits_credit_role 
    FOREIGN KEY (credit_role) 
    REFERENCES credit_role_lks(credit_role_lk) 
    ON DELETE CASCADE
);

-- Video links table (many-to-many relationship)
CREATE TABLE IF NOT EXISTS video_links (
  video_link_id SERIAL PRIMARY KEY,
  video_id INTEGER NOT NULL,
  link_type VARCHAR(50) NOT NULL,
  link_url VARCHAR(500) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  -- Foreign key constraints
  CONSTRAINT fk_video_links_video_id 
    FOREIGN KEY (video_id) 
    REFERENCES videos(video_id) 
    ON DELETE CASCADE,

  CONSTRAINT fk_video_links_link_type 
    FOREIGN KEY (link_type) 
    REFERENCES link_type_lks(link_type_lk) 
    ON DELETE CASCADE
);

-- Insert lookup data
INSERT INTO video_type_lks (video_type_lk, video_type_description) VALUES
  ('film', 'Feature film or short film'),
  ('music_video', 'Music video'),
  ('promotional', 'Promotional or commercial video');

INSERT INTO credit_role_lks (credit_role_lk, credit_role_description) VALUES
  ('starring', 'Actor or performer in the video'),
  ('director', 'Director of the video'),
  ('writer', 'Writer or screenwriter'),
  ('cinematographer', 'Cinematographer or director of photography'),
  ('producer', 'Producer of the video'),
  ('a_camera', 'A Camera operator'),
  ('editor', 'Editor of the video'),
  ('composer', 'Composer of the music');

INSERT INTO link_type_lks (link_type_lk, link_type_description) VALUES
  ('youtube', 'YouTube video link'),
  ('vimeo', 'Vimeo video link'),
  ('instagram', 'Instagram post link'),
  ('website', 'External website link');

-- Set the starting value for video_id to 1000
ALTER SEQUENCE videos_video_id_seq RESTART WITH 1000;
ALTER SEQUENCE video_credits_video_credit_id_seq RESTART WITH 1000;
ALTER SEQUENCE video_links_video_link_id_seq RESTART WITH 1000;

-- Add indexes for common queries
CREATE INDEX idx_videos_video_type ON videos(video_type);
CREATE INDEX idx_videos_weight ON videos(weight);
CREATE INDEX idx_videos_created_at ON videos(created_at);
CREATE INDEX idx_video_credits_video_id ON video_credits(video_id);
CREATE INDEX idx_video_credits_cast_member_id ON video_credits(cast_member_id);
CREATE INDEX idx_video_credits_credit_role ON video_credits(credit_role);
CREATE INDEX idx_video_links_video_id ON video_links(video_id);
CREATE INDEX idx_video_links_link_type ON video_links(link_type);
CREATE UNIQUE INDEX idx_cast_members_name ON cast_members(name);

-- +goose Down

DROP TABLE IF EXISTS video_links;
DROP TABLE IF EXISTS video_credits;
DROP TABLE IF EXISTS cast_members;
DROP TABLE IF EXISTS videos;
DROP TABLE IF EXISTS link_type_lks;
DROP TABLE IF EXISTS credit_role_lks;
DROP TABLE IF EXISTS video_type_lks; 