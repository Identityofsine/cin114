-- +goose Up

DROP VIEW IF EXISTS owned_tickets;

CREATE OR REPLACE VIEW owned_tickets AS
SELECT 
  t.ticket_id,
  e.event_id as event_id,
  e.description AS event_description,
  sp.billing_email AS stripe_receipt_email,
  t.created_at,
  t.updated_at
FROM tickets t
JOIN events e ON t.event_id = e.event_id
LEFT JOIN stripe_payments sp ON t.stripe_payment_id = sp.id;

-- +goose Down

DROP VIEW IF EXISTS owned_tickets;

CREATE OR REPLACE VIEW owned_tickets AS
SELECT 
  t.ticket_id,
  e.event_id as event_id,
  e.description AS event_description,
  sp.billing_email AS stripe_receipt_email
FROM tickets t
JOIN events e ON t.event_id = e.event_id
LEFT JOIN stripe_payments sp ON t.stripe_payment_id = sp.id;
