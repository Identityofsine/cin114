package model

import (
	"encoding/json"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type StripePaymentDB struct {
	Id                   string  `json:"id"`
	Object               string  `json:"object"`
	Amount               int64   `json:"amount"`
	AmountCaptured       int64   `json:"amount_captured"`
	AmountRefunded       int64   `json:"amount_refunded"`
	Application          *string `json:"application"`
	ApplicationFee       *string `json:"application_fee"`
	ApplicationFeeAmount *int64  `json:"application_fee_amount"`
	BalanceTransaction   *string `json:"balance_transaction"`

	// Billing Details
	BillingEmail             *string `json:"billing_email"`
	BillingName              *string `json:"billing_name"`
	BillingPhone             *string `json:"billing_phone"`
	BillingAddressLine1      *string `json:"billing_address_line1"`
	BillingAddressLine2      *string `json:"billing_address_line2"`
	BillingAddressCity       *string `json:"billing_address_city"`
	BillingAddressState      *string `json:"billing_address_state"`
	BillingAddressPostalCode *string `json:"billing_address_postal_code"`
	BillingAddressCountry    *string `json:"billing_address_country"`

	CalculatedStatementDescriptor *string         `json:"calculated_statement_descriptor"`
	Captured                      bool            `json:"captured"`
	Created                       int64           `json:"created"` // Unix timestamp
	Currency                      string          `json:"currency"`
	Customer                      *string         `json:"customer"`
	Description                   *string         `json:"description"`
	Disputed                      bool            `json:"disputed"`
	FailureBalanceTransaction     *string         `json:"failure_balance_transaction"`
	FailureCode                   *string         `json:"failure_code"`
	FailureMessage                *string         `json:"failure_message"`
	FraudDetails                  json.RawMessage `json:"fraud_details"`
	Livemode                      bool            `json:"livemode"`
	Metadata                      json.RawMessage `json:"metadata"`
	OnBehalfOf                    *string         `json:"on_behalf_of"`

	// Outcome Details
	OutcomeNetworkStatus *string `json:"outcome_network_status"`
	OutcomeReason        *string `json:"outcome_reason"`
	OutcomeRiskLevel     *string `json:"outcome_risk_level"`
	OutcomeRiskScore     *int64  `json:"outcome_risk_score"`
	OutcomeSellerMessage *string `json:"outcome_seller_message"`
	OutcomeType          *string `json:"outcome_type"`

	Paid          bool    `json:"paid"`
	PaymentIntent *string `json:"payment_intent"`
	PaymentMethod *string `json:"payment_method"`

	// Payment Method Details
	PaymentMethodType *string `json:"payment_method_type"`
	CardBrand         *string `json:"card_brand"`
	CardCountry       *string `json:"card_country"`
	CardExpMonth      *int64  `json:"card_exp_month"`
	CardExpYear       *int64  `json:"card_exp_year"`
	CardFingerprint   *string `json:"card_fingerprint"`
	CardFunding       *string `json:"card_funding"`
	CardLast4         *string `json:"card_last4"`
	CardNetwork       *string `json:"card_network"`

	ReceiptEmail              *string         `json:"receipt_email"`
	ReceiptNumber             *string         `json:"receipt_number"`
	ReceiptUrl                *string         `json:"receipt_url"`
	Refunded                  bool            `json:"refunded"`
	Review                    *string         `json:"review"`
	Shipping                  json.RawMessage `json:"shipping"`
	SourceTransfer            *string         `json:"source_transfer"`
	StatementDescriptor       *string         `json:"statement_descriptor"`
	StatementDescriptorSuffix *string         `json:"statement_descriptor_suffix"`
	Status                    string          `json:"status"`
	TransferData              json.RawMessage `json:"transfer_data"`
	TransferGroup             *string         `json:"transfer_group"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllStripePayments() ([]StripePaymentDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_payments ORDER BY created_at DESC"
	rows, err := db.Query[StripePaymentDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetStripePaymentById(id string) (*StripePaymentDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_payments WHERE id = $1"
	rows, err := db.Query[StripePaymentDB](query, id)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetStripePaymentById", "Stripe payment not found", "stripe-payment-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetStripePaymentsByStatus(status string) ([]StripePaymentDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_payments WHERE status = $1 ORDER BY created_at DESC"
	rows, err := db.Query[StripePaymentDB](query, status)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetStripePaymentsByCustomer(customer string) ([]StripePaymentDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_payments WHERE customer = $1 ORDER BY created_at DESC"
	rows, err := db.Query[StripePaymentDB](query, customer)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetStripePaymentsByEmail(email string) ([]StripePaymentDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_payments WHERE billing_email = $1 ORDER BY created_at DESC"
	rows, err := db.Query[StripePaymentDB](query, email)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func CreateStripePayment(payment *StripePaymentDB) db.DatabaseError {
	query := `INSERT INTO stripe_payments (
		id, object, amount, amount_captured, amount_refunded, application, application_fee, 
		application_fee_amount, balance_transaction, billing_email, billing_name, billing_phone,
		billing_address_line1, billing_address_line2, billing_address_city, billing_address_state,
		billing_address_postal_code, billing_address_country, calculated_statement_descriptor,
		captured, created, currency, customer, description, disputed, failure_balance_transaction,
		failure_code, failure_message, fraud_details, livemode, metadata, on_behalf_of,
		outcome_network_status, outcome_reason, outcome_risk_level, outcome_risk_score,
		outcome_seller_message, outcome_type, paid, payment_intent, payment_method,
		payment_method_type, card_brand, card_country, card_exp_month, card_exp_year,
		card_fingerprint, card_funding, card_last4, card_network, receipt_email,
		receipt_number, receipt_url, refunded, review, shipping, source_transfer,
		statement_descriptor, statement_descriptor_suffix, status, transfer_data, transfer_group
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,
		$20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36,
		$37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53,
		$54, $55, $56, $57, $58, $59, $60, $61
	) RETURNING created_at, updated_at`

	rows, err := db.Query[StripePaymentDB](query,
		payment.Id, payment.Object, payment.Amount, payment.AmountCaptured, payment.AmountRefunded,
		payment.Application, payment.ApplicationFee, payment.ApplicationFeeAmount, payment.BalanceTransaction,
		payment.BillingEmail, payment.BillingName, payment.BillingPhone, payment.BillingAddressLine1,
		payment.BillingAddressLine2, payment.BillingAddressCity, payment.BillingAddressState,
		payment.BillingAddressPostalCode, payment.BillingAddressCountry, payment.CalculatedStatementDescriptor,
		payment.Captured, payment.Created, payment.Currency, payment.Customer, payment.Description,
		payment.Disputed, payment.FailureBalanceTransaction, payment.FailureCode, payment.FailureMessage,
		payment.FraudDetails, payment.Livemode, payment.Metadata, payment.OnBehalfOf,
		payment.OutcomeNetworkStatus, payment.OutcomeReason, payment.OutcomeRiskLevel, payment.OutcomeRiskScore,
		payment.OutcomeSellerMessage, payment.OutcomeType, payment.Paid, payment.PaymentIntent,
		payment.PaymentMethod, payment.PaymentMethodType, payment.CardBrand, payment.CardCountry,
		payment.CardExpMonth, payment.CardExpYear, payment.CardFingerprint, payment.CardFunding,
		payment.CardLast4, payment.CardNetwork, payment.ReceiptEmail, payment.ReceiptNumber,
		payment.ReceiptUrl, payment.Refunded, payment.Review, payment.Shipping, payment.SourceTransfer,
		payment.StatementDescriptor, payment.StatementDescriptorSuffix, payment.Status,
		payment.TransferData, payment.TransferGroup,
	)

	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		payment.CreatedAt = (*rows)[0].CreatedAt
		payment.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateStripePaymentStatus(id string, status string) db.DatabaseError {
	query := "UPDATE stripe_payments SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"
	_, err := db.Query[StripePaymentDB](query, status, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteStripePayment(id string) db.DatabaseError {
	query := "DELETE FROM stripe_payments WHERE id = $1"
	_, err := db.Query[StripePaymentDB](query, id)
	if err != nil {
		return err
	}
	return nil
}

func StripePaymentExists(id string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM stripe_payments WHERE id = $1)"
	rows, err := db.Query[bool](query, id)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}

func GetSuccessfulStripePayments() ([]StripePaymentDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_payments WHERE status = 'succeeded' AND paid = true ORDER BY created_at DESC"
	rows, err := db.Query[StripePaymentDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetFailedStripePayments() ([]StripePaymentDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_payments WHERE status = 'failed' ORDER BY created_at DESC"
	rows, err := db.Query[StripePaymentDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}
