package payment

import (
	"fmt"
	"strconv"

	. "github.com/identityofsine/fofx-go-gin-api-template/api/dto/event"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
	"github.com/stripe/stripe-go/v82/price"
	"github.com/stripe/stripe-go/v82/webhookendpoint"
)

type EventStripePriceForm struct {
	Event Event `json:"event"`

	Price int64 `json:"price"` // in cents
}

type CheckoutSessionForm struct {
	EventId    int64  `json:"event_id" binding:"required"`
	Quantity   int64  `json:"quantity" binding:"required"`
	SuccessURL string `json:"success_url" binding:"required"`
	CancelURL  string `json:"cancel_url" binding:"required"`
}

var (
	stripeConfig = config.GetStripeSettings()
)

func CreateStripePriceForScreening(form EventStripePriceForm) (*model.EventStripePriceDB, db.DatabaseError) {

	stripe.Key = stripeConfig.StripeSecretKey

	if exists, err := doesStripePriceExist(form.Event.EventId); err != nil || exists {
		return nil, db.NewDatabaseError(
			"CreateStripePriceForScreening",
			"Stripe price already exists for this event",
			"stripe-price-exists",
			400,
		)
	}

	params := &stripe.PriceParams{
		Currency:   stripe.String(stripe.CurrencyUSD),
		UnitAmount: stripe.Int64(form.Price),
		ProductData: &stripe.PriceProductDataParams{
			Name:                stripe.String(string(*form.Event.ShortDescription)),
			StatementDescriptor: stripe.String("CIN114 - Scrn Ticket"),
			UnitLabel:           stripe.String("ticket"),
		},
	}

	price, err := price.New(params)
	if err != nil {
		return nil, db.NewDatabaseError(
			"CreateStripePriceForScreening",
			err.Error(),
			"stripe-price-creation-failed",
			500,
		)
	}

	if stripePayment, err := createStripePaymentInDatabase(price, form.Event); err != nil || stripePayment == nil {
		return nil, err
	} else {
		return stripePayment, nil
	}

}

func CreateCheckoutSessionForEvent(form CheckoutSessionForm) (*stripe.CheckoutSession, db.DatabaseError) {
	stripe.Key = stripeConfig.StripeSecretKey

	// Get the event from database
	event, err := model.GetEventById(form.EventId)
	if err != nil {
		return nil, db.NewDatabaseError(
			"CreateCheckoutSessionForEvent",
			"Event not found",
			"event-not-found",
			404,
		)
	}

	// Get the active stripe price for this event
	eventStripePrice, err := model.GetActiveEventStripePriceByEventId(form.EventId)
	if err != nil {
		return nil, db.NewDatabaseError(
			"CreateCheckoutSessionForEvent",
			"No active price found for this event",
			"no-active-price",
			404,
		)
	}

	// Create checkout session parameters
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(form.SuccessURL),
		CancelURL:  stripe.String(form.CancelURL),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(eventStripePrice.StripePriceId),
				Quantity: stripe.Int64(form.Quantity),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		Metadata: map[string]string{
			"event_id":              strconv.FormatInt(form.EventId, 10),
			"event_description":     event.Description,
			"event_stripe_price_id": strconv.FormatInt(eventStripePrice.Id, 10),
			"quantity":              strconv.FormatInt(form.Quantity, 10),
		},
	}

	// Create the checkout session
	checkoutSession, sessionErr := session.New(params)
	if sessionErr != nil {
		return nil, db.NewDatabaseError(
			"CreateCheckoutSessionForEvent",
			sessionErr.Error(),
			"checkout-session-creation-failed",
			500,
		)
	}

	return checkoutSession, nil
}

func doesStripePriceExist(eventId int64) (bool, error) {
	model, err := model.GetEventStripePricesByEventId(eventId)
	if err != nil {
		return false, err
	}

	if model == nil {
		return false, nil
	} else if len(model) == 0 {
		return false, nil
	}

	return true, nil
}

func createStripePaymentInDatabase(price *stripe.Price, event Event) (*model.EventStripePriceDB, db.DatabaseError) {
	// First, create the stripe price record
	stripePriceDB := model.StripePriceDB{
		Id:                price.ID,
		Object:            string(price.Object),
		Active:            price.Active,
		BillingScheme:     string(price.BillingScheme),
		Created:           price.Created,
		Currency:          string(price.Currency),
		CustomUnitAmount:  []byte("{}"), // Empty JSON object
		Livemode:          price.Livemode,
		Metadata:          []byte("{}"), // Empty JSON object
		Product:           price.Product.ID,
		TransformQuantity: []byte("{}"), // Empty JSON object
		Type:              string(price.Type),
		UnitAmount:        &price.UnitAmount,
	}

	// Handle recurring details if present
	if price.Recurring != nil {
		interval := string(price.Recurring.Interval)
		stripePriceDB.RecurringInterval = &interval
		stripePriceDB.RecurringIntervalCount = &price.Recurring.IntervalCount
		if price.Recurring.TrialPeriodDays != 0 {
			stripePriceDB.RecurringTrialPeriodDays = &price.Recurring.TrialPeriodDays
		}
		usageType := string(price.Recurring.UsageType)
		stripePriceDB.RecurringUsageType = &usageType
	}

	// Handle optional fields
	if price.LookupKey != "" {
		stripePriceDB.LookupKey = &price.LookupKey
	}
	if price.Nickname != "" {
		stripePriceDB.Nickname = &price.Nickname
	}
	if price.TaxBehavior != "" {
		taxBehavior := string(price.TaxBehavior)
		stripePriceDB.TaxBehavior = &taxBehavior
	}

	// Create the stripe price in the database
	if err := model.CreateStripePrice(&stripePriceDB); err != nil {
		return nil, err
	}

	// Now create the event-stripe price association
	eventStripePriceDB := model.EventStripePriceDB{
		EventId:       event.EventId,
		StripePriceId: price.ID,
		IsActive:      true,
	}

	if err := model.CreateEventStripePrice(&eventStripePriceDB); err != nil {
		return nil, err
	}

	return &eventStripePriceDB, nil
}

func StripeSetupWebhookOnCheckoutSession() db.DatabaseError {
	stripe.Key = stripeConfig.StripeSecretKey
	serverCfg := config.GetServerDetails()

	// Skip webhook creation in development
	if serverCfg.Environment == "development" || serverCfg.Environment == "dev" {
		storedlogs.LogInfo("Skipping webhook creation in development environment")
		return nil
	}

	// Construct webhook URL from server config
	webhookURL := fmt.Sprintf("%s/api/v1/stripe/webhook", serverCfg.WebServerConfig.URI)

	params := &stripe.WebhookEndpointParams{
		URL: stripe.String(webhookURL),
		EnabledEvents: stripe.StringSlice([]string{
			"checkout.session.completed",
		}),
	}

	_, err := webhookendpoint.New(params)
	if err != nil {
		return db.NewDatabaseError(
			"StripeSetupWebhookOnCheckoutSession",
			err.Error(),
			"stripe-webhook-setup-failed",
			500,
		)
	}

	storedlogs.LogInfo("Stripe webhook endpoint created successfully")

	return nil
}
