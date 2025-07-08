-- +goose Up

CREATE TABLE IF NOT EXISTS stripe_prices (
  id VARCHAR(255) PRIMARY KEY,  -- Stripe price ID like "price_1MoBy5LkdIwHu7ixZhnattbh"
  object VARCHAR(50) NOT NULL DEFAULT 'price',
  active BOOLEAN NOT NULL DEFAULT true,
  billing_scheme VARCHAR(50) NOT NULL DEFAULT 'per_unit',
  created INTEGER NOT NULL,  -- Unix timestamp from Stripe
  currency VARCHAR(3) NOT NULL DEFAULT 'usd',
  custom_unit_amount JSONB,
  livemode BOOLEAN NOT NULL DEFAULT false,
  lookup_key VARCHAR(255),
  metadata JSONB,
  nickname VARCHAR(255),
  product VARCHAR(255) NOT NULL,  -- Stripe product ID
  
  -- Recurring details (flattened)
  recurring_interval VARCHAR(50),  -- month, year, week, day
  recurring_interval_count INTEGER,
  recurring_trial_period_days INTEGER,
  recurring_usage_type VARCHAR(50),
  
  tax_behavior VARCHAR(50),
  tiers_mode VARCHAR(50),
  transform_quantity JSONB,
  type VARCHAR(50) NOT NULL,  -- one_time, recurring
  unit_amount INTEGER,
  unit_amount_decimal VARCHAR(50),
  
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS event_stripe_price (
  id SERIAL PRIMARY KEY,
  event_id INTEGER NOT NULL REFERENCES events(event_id) ON DELETE CASCADE,
  stripe_price_id VARCHAR(255) NOT NULL REFERENCES stripe_prices(id) ON DELETE CASCADE,
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  -- Ensure an event can only have one active price at a time
  UNIQUE(event_id, is_active) DEFERRABLE INITIALLY DEFERRED
);

-- Add indexes for common queries
CREATE INDEX idx_stripe_prices_active ON stripe_prices(active);
CREATE INDEX idx_stripe_prices_created ON stripe_prices(created);
CREATE INDEX idx_stripe_prices_product ON stripe_prices(product);
CREATE INDEX idx_stripe_prices_type ON stripe_prices(type);

CREATE INDEX idx_event_stripe_price_event_id ON event_stripe_price(event_id);
CREATE INDEX idx_event_stripe_price_stripe_price_id ON event_stripe_price(stripe_price_id);
CREATE INDEX idx_event_stripe_price_active ON event_stripe_price(is_active);

-- +goose Down

DROP TABLE IF EXISTS event_stripe_price;
DROP TABLE IF EXISTS stripe_prices; 