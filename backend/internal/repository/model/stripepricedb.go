package model

import (
	"encoding/json"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type StripePriceDB struct {
	Id               string          `json:"id"`
	Object           string          `json:"object"`
	Active           bool            `json:"active"`
	BillingScheme    string          `json:"billing_scheme"`
	Created          int64           `json:"created"` // Unix timestamp
	Currency         string          `json:"currency"`
	CustomUnitAmount json.RawMessage `json:"custom_unit_amount"`
	Livemode         bool            `json:"livemode"`
	LookupKey        *string         `json:"lookup_key"`
	Metadata         json.RawMessage `json:"metadata"`
	Nickname         *string         `json:"nickname"`
	Product          string          `json:"product"`

	// Recurring Details
	RecurringInterval        *string `json:"recurring_interval"`
	RecurringIntervalCount   *int64  `json:"recurring_interval_count"`
	RecurringTrialPeriodDays *int64  `json:"recurring_trial_period_days"`
	RecurringUsageType       *string `json:"recurring_usage_type"`

	TaxBehavior       *string         `json:"tax_behavior"`
	TiersMode         *string         `json:"tiers_mode"`
	TransformQuantity json.RawMessage `json:"transform_quantity"`
	Type              string          `json:"type"`
	UnitAmount        *int64          `json:"unit_amount"`
	UnitAmountDecimal *string         `json:"unit_amount_decimal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllStripePrices() ([]StripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_prices ORDER BY created_at DESC"
	rows, err := db.Query[StripePriceDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetStripePriceById(id string) (*StripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_prices WHERE id = $1"
	rows, err := db.Query[StripePriceDB](query, id)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetStripePriceById", "Stripe price not found", "stripe-price-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetActivestripePrices() ([]StripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_prices WHERE active = true ORDER BY created_at DESC"
	rows, err := db.Query[StripePriceDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetStripePricesByProduct(product string) ([]StripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_prices WHERE product = $1 ORDER BY created_at DESC"
	rows, err := db.Query[StripePriceDB](query, product)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetStripePricesByType(priceType string) ([]StripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM stripe_prices WHERE type = $1 ORDER BY created_at DESC"
	rows, err := db.Query[StripePriceDB](query, priceType)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func CreateStripePrice(price *StripePriceDB) db.DatabaseError {
	query := `INSERT INTO stripe_prices (
		id, object, active, billing_scheme, created, currency, custom_unit_amount,
		livemode, lookup_key, metadata, nickname, product, recurring_interval,
		recurring_interval_count, recurring_trial_period_days, recurring_usage_type,
		tax_behavior, tiers_mode, transform_quantity, type, unit_amount, unit_amount_decimal
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22
	) RETURNING created_at, updated_at`

	type TimestampResult struct {
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	rows, err := db.Query[TimestampResult](query,
		price.Id, price.Object, price.Active, price.BillingScheme, price.Created,
		price.Currency, price.CustomUnitAmount, price.Livemode, price.LookupKey,
		price.Metadata, price.Nickname, price.Product, price.RecurringInterval,
		price.RecurringIntervalCount, price.RecurringTrialPeriodDays, price.RecurringUsageType,
		price.TaxBehavior, price.TiersMode, price.TransformQuantity, price.Type,
		price.UnitAmount, price.UnitAmountDecimal,
	)

	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		price.CreatedAt = (*rows)[0].CreatedAt
		price.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateStripePriceActive(id string, active bool) db.DatabaseError {
	query := "UPDATE stripe_prices SET active = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"
	_, err := db.Query[StripePriceDB](query, active, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteStripePrice(id string) db.DatabaseError {
	query := "DELETE FROM stripe_prices WHERE id = $1"
	_, err := db.Query[StripePriceDB](query, id)
	if err != nil {
		return err
	}
	return nil
}

func StripePriceExists(id string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM stripe_prices WHERE id = $1)"
	rows, err := db.Query[bool](query, id)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}
