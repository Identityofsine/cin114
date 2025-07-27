-- +goose Up

CREATE TABLE IF NOT EXISTS tickets (
  ticket_id SERIAL PRIMARY KEY,
  event_id INTEGER NOT NULL,
  stripe_payment_id VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Foreign key constraints
  CONSTRAINT fk_tickets_event_id 
    FOREIGN KEY (event_id) 
    REFERENCES events(event_id) 
    ON DELETE CASCADE,
    
  CONSTRAINT fk_tickets_stripe_payment_id 
    FOREIGN KEY (stripe_payment_id) 
    REFERENCES stripe_payments(id) 
    ON DELETE SET NULL
);

-- Set the starting value for ticket_id to 100,000
ALTER SEQUENCE tickets_ticket_id_seq RESTART WITH 100000;

-- Add indexes for common queries
CREATE INDEX idx_tickets_event_id ON tickets(event_id);
CREATE INDEX idx_tickets_stripe_payment_id ON tickets(stripe_payment_id);
CREATE INDEX idx_tickets_created_at ON tickets(created_at);

-- +goose Down

DROP TABLE IF EXISTS tickets; 