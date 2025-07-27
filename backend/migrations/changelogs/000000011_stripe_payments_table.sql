-- +goose Up

CREATE TABLE IF NOT EXISTS stripe_payments (
  id VARCHAR(255) PRIMARY KEY,  -- Stripe charge ID like "ch_3MmlLrLkdIwHu7ix0snN0B15"
  object VARCHAR(50) NOT NULL DEFAULT 'charge',
  amount INTEGER NOT NULL,
  amount_captured INTEGER NOT NULL,
  amount_refunded INTEGER NOT NULL DEFAULT 0,
  application VARCHAR(255),
  application_fee VARCHAR(255),
  application_fee_amount INTEGER,
  balance_transaction VARCHAR(255),
  
  -- Billing Details (flattened)
  billing_email VARCHAR(255),
  billing_name VARCHAR(255),
  billing_phone VARCHAR(50),
  billing_address_line1 VARCHAR(255),
  billing_address_line2 VARCHAR(255),
  billing_address_city VARCHAR(100),
  billing_address_state VARCHAR(100),
  billing_address_postal_code VARCHAR(20),
  billing_address_country VARCHAR(2),
  
  calculated_statement_descriptor VARCHAR(255),
  captured BOOLEAN NOT NULL DEFAULT true,
  created INTEGER NOT NULL,  -- Unix timestamp from Stripe
  currency VARCHAR(3) NOT NULL DEFAULT 'usd',
  customer VARCHAR(255),
  description TEXT,
  disputed BOOLEAN NOT NULL DEFAULT false,
  failure_balance_transaction VARCHAR(255),
  failure_code VARCHAR(100),
  failure_message TEXT,
  fraud_details JSONB,
  livemode BOOLEAN NOT NULL DEFAULT false,
  metadata JSONB,
  on_behalf_of VARCHAR(255),
  
  -- Outcome details (flattened)
  outcome_network_status VARCHAR(50),
  outcome_reason VARCHAR(255),
  outcome_risk_level VARCHAR(50),
  outcome_risk_score INTEGER,
  outcome_seller_message TEXT,
  outcome_type VARCHAR(50),
  
  paid BOOLEAN NOT NULL DEFAULT false,
  payment_intent VARCHAR(255),
  payment_method VARCHAR(255),
  
  -- Payment Method Details (flattened for card payments)
  payment_method_type VARCHAR(50),
  card_brand VARCHAR(50),
  card_country VARCHAR(2),
  card_exp_month INTEGER,
  card_exp_year INTEGER,
  card_fingerprint VARCHAR(255),
  card_funding VARCHAR(20),
  card_last4 VARCHAR(4),
  card_network VARCHAR(50),
  
  receipt_email VARCHAR(255),
  receipt_number VARCHAR(255),
  receipt_url TEXT,
  refunded BOOLEAN NOT NULL DEFAULT false,
  review VARCHAR(255),
  shipping JSONB,
  source_transfer VARCHAR(255),
  statement_descriptor VARCHAR(255),
  statement_descriptor_suffix VARCHAR(255),
  status VARCHAR(50) NOT NULL,
  transfer_data JSONB,
  transfer_group VARCHAR(255),
  
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes for common queries
CREATE INDEX idx_stripe_payments_status ON stripe_payments(status);
CREATE INDEX idx_stripe_payments_created ON stripe_payments(created);
CREATE INDEX idx_stripe_payments_billing_email ON stripe_payments(billing_email);
CREATE INDEX idx_stripe_payments_customer ON stripe_payments(customer); 