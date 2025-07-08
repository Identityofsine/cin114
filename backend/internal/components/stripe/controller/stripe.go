package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/charge"
	"github.com/stripe/stripe-go/v82/webhook"
)

var (
	stripeConfig = config.GetStripeSettings()
)

func HandleStripeWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		storedlogs.LogError("Error reading request body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// This is your Stripe CLI webhook secret for testing your endpoint locally.
	// In production, you should set this to your webhook endpoint secret
	endpointSecret := stripeConfig.WebhookSecret
	if endpointSecret == "" {
		storedlogs.LogWarn("Webhook secret not configured, skipping signature verification")
	}

	// Verify the payload came from Stripe
	if endpointSecret != "" {
		signatureHeader := c.GetHeader("Stripe-Signature")
		event, err := webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
		if err != nil {
			storedlogs.LogError("Error verifying webhook signature", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
			return
		}

		handleStripeEvent(c, event)
	} else {
		// For development without signature verification
		var event stripe.Event
		if err := json.Unmarshal(payload, &event); err != nil {
			storedlogs.LogError("Error parsing webhook JSON", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		handleStripeEvent(c, event)
	}
}

func handleStripeEvent(c *gin.Context, event stripe.Event) {
	switch event.Type {
	case "checkout.session.completed":
		handleCheckoutSessionCompleted(c, event)
	default:
		storedlogs.LogInfo("Unhandled event type: " + string(event.Type))
		c.JSON(http.StatusOK, gin.H{"message": "Event type not handled"})
	}
}

func handleCheckoutSessionCompleted(c *gin.Context, event stripe.Event) {
	var checkoutSession stripe.CheckoutSession
	if err := json.Unmarshal(event.Data.Raw, &checkoutSession); err != nil {
		storedlogs.LogError("Error parsing checkout session", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checkout session data"})
		return
	}

	storedlogs.LogInfo("Processing checkout session completed: " + checkoutSession.ID)

	// Extract event_id from metadata
	eventIdStr, exists := checkoutSession.Metadata["event_id"]
	if !exists {
		storedlogs.LogError("No event_id found in checkout session metadata", nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing event_id in metadata"})
		return
	}

	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		storedlogs.LogError("Invalid event_id in metadata: "+eventIdStr, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event_id"})
		return
	}

	// Get the payment intent ID from the checkout session
	paymentIntentId := checkoutSession.PaymentIntent.ID
	if paymentIntentId == "" {
		storedlogs.LogError("No payment intent found in checkout session", nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No payment intent found"})
		return
	}

	// Retrieve the charge from the payment intent
	stripe.Key = stripeConfig.StripeSecretKey

	// Get the charge using the payment intent
	params := &stripe.ChargeListParams{
		PaymentIntent: stripe.String(paymentIntentId),
	}
	iter := charge.List(params)

	var stripeCharge *stripe.Charge
	for iter.Next() {
		stripeCharge = iter.Charge()
		break // We only need the first (and should be only) charge
	}

	if err := iter.Err(); err != nil {
		storedlogs.LogError("Error retrieving charge", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving charge"})
		return
	}

	if stripeCharge == nil {
		storedlogs.LogError("No charge found for payment intent: "+paymentIntentId, nil)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No charge found"})
		return
	}

	// Create the payment record in the database
	if dbErr := createPaymentRecord(stripeCharge); dbErr != nil {
		storedlogs.LogError("Error creating payment record", dbErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating payment record"})
		return
	}

	// Calculate the number of tickets to create
	quantity := int64(1) // Default to 1 if not found
	if checkoutSession.Metadata["quantity"] != "" {
		if q, err := strconv.ParseInt(checkoutSession.Metadata["quantity"], 10, 64); err == nil {
			quantity = q
		}
	}

	// Extract quantity from line items if available
	if checkoutSession.LineItems != nil && len(checkoutSession.LineItems.Data) > 0 {
		quantity = checkoutSession.LineItems.Data[0].Quantity
	}

	// Create ticket records
	for i := int64(0); i < quantity; i++ {
		ticket := &model.TicketDB{
			EventId:         eventId,
			StripePaymentId: &stripeCharge.ID,
		}

		if err := model.CreateTicket(ticket); err != nil {
			storedlogs.LogError("Error creating ticket", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating ticket"})
			return
		}
	}

	storedlogs.LogInfo(fmt.Sprintf("Successfully created payment record and %d tickets for event %d", quantity, eventId))
	c.JSON(http.StatusOK, gin.H{"message": "Webhook processed successfully"})
}

func createPaymentRecord(charge *stripe.Charge) db.DatabaseError {
	payment := &model.StripePaymentDB{
		Id:             charge.ID,
		Object:         string(charge.Object),
		Amount:         charge.Amount,
		AmountCaptured: charge.AmountCaptured,
		AmountRefunded: charge.AmountRefunded,
		Captured:       charge.Captured,
		Created:        charge.Created,
		Currency:       string(charge.Currency),
		Disputed:       charge.Disputed,
		Livemode:       charge.Livemode,
		Paid:           charge.Paid,
		Refunded:       charge.Refunded,
		Status:         string(charge.Status),
	}

	// Handle optional fields
	if charge.Application != nil {
		payment.Application = &charge.Application.ID
	}
	if charge.ApplicationFee != nil {
		payment.ApplicationFee = &charge.ApplicationFee.ID
	}
	if charge.ApplicationFeeAmount != 0 {
		payment.ApplicationFeeAmount = &charge.ApplicationFeeAmount
	}
	if charge.BalanceTransaction != nil {
		payment.BalanceTransaction = &charge.BalanceTransaction.ID
	}
	if charge.CalculatedStatementDescriptor != "" {
		payment.CalculatedStatementDescriptor = &charge.CalculatedStatementDescriptor
	}
	if charge.Customer != nil {
		payment.Customer = &charge.Customer.ID
	}
	if charge.Description != "" {
		payment.Description = &charge.Description
	}
	if charge.FailureBalanceTransaction != nil {
		payment.FailureBalanceTransaction = &charge.FailureBalanceTransaction.ID
	}
	if charge.FailureCode != "" {
		payment.FailureCode = &charge.FailureCode
	}
	if charge.FailureMessage != "" {
		payment.FailureMessage = &charge.FailureMessage
	}
	if charge.OnBehalfOf != nil {
		payment.OnBehalfOf = &charge.OnBehalfOf.ID
	}
	if charge.PaymentIntent != nil {
		payment.PaymentIntent = &charge.PaymentIntent.ID
	}
	if charge.PaymentMethod != "" {
		payment.PaymentMethod = &charge.PaymentMethod
	}
	if charge.ReceiptEmail != "" {
		payment.ReceiptEmail = &charge.ReceiptEmail
	}
	if charge.ReceiptNumber != "" {
		payment.ReceiptNumber = &charge.ReceiptNumber
	}
	if charge.ReceiptURL != "" {
		payment.ReceiptUrl = &charge.ReceiptURL
	}
	if charge.Review != nil {
		payment.Review = &charge.Review.ID
	}
	if charge.SourceTransfer != nil {
		payment.SourceTransfer = &charge.SourceTransfer.ID
	}
	if charge.StatementDescriptor != "" {
		payment.StatementDescriptor = &charge.StatementDescriptor
	}
	if charge.StatementDescriptorSuffix != "" {
		payment.StatementDescriptorSuffix = &charge.StatementDescriptorSuffix
	}
	if charge.TransferGroup != "" {
		payment.TransferGroup = &charge.TransferGroup
	}

	// Handle billing details
	if charge.BillingDetails != nil {
		if charge.BillingDetails.Email != "" {
			payment.BillingEmail = &charge.BillingDetails.Email
		}
		if charge.BillingDetails.Name != "" {
			payment.BillingName = &charge.BillingDetails.Name
		}
		if charge.BillingDetails.Phone != "" {
			payment.BillingPhone = &charge.BillingDetails.Phone
		}
		if charge.BillingDetails.Address != nil {
			if charge.BillingDetails.Address.Line1 != "" {
				payment.BillingAddressLine1 = &charge.BillingDetails.Address.Line1
			}
			if charge.BillingDetails.Address.Line2 != "" {
				payment.BillingAddressLine2 = &charge.BillingDetails.Address.Line2
			}
			if charge.BillingDetails.Address.City != "" {
				payment.BillingAddressCity = &charge.BillingDetails.Address.City
			}
			if charge.BillingDetails.Address.State != "" {
				payment.BillingAddressState = &charge.BillingDetails.Address.State
			}
			if charge.BillingDetails.Address.PostalCode != "" {
				payment.BillingAddressPostalCode = &charge.BillingDetails.Address.PostalCode
			}
			if charge.BillingDetails.Address.Country != "" {
				payment.BillingAddressCountry = &charge.BillingDetails.Address.Country
			}
		}
	}

	// Handle outcome details
	if charge.Outcome != nil {
		if charge.Outcome.NetworkStatus != "" {
			payment.OutcomeNetworkStatus = &charge.Outcome.NetworkStatus
		}
		if charge.Outcome.Reason != "" {
			payment.OutcomeReason = &charge.Outcome.Reason
		}
		if charge.Outcome.RiskLevel != "" {
			payment.OutcomeRiskLevel = &charge.Outcome.RiskLevel
		}
		if charge.Outcome.RiskScore != 0 {
			payment.OutcomeRiskScore = &charge.Outcome.RiskScore
		}
		if charge.Outcome.SellerMessage != "" {
			payment.OutcomeSellerMessage = &charge.Outcome.SellerMessage
		}
		if charge.Outcome.Type != "" {
			payment.OutcomeType = &charge.Outcome.Type
		}
	}

	// Handle payment method details
	if charge.PaymentMethodDetails != nil && charge.PaymentMethodDetails.Card != nil {
		card := charge.PaymentMethodDetails.Card
		paymentMethodType := "card"
		payment.PaymentMethodType = &paymentMethodType

		if card.Brand != "" {
			brand := string(card.Brand)
			payment.CardBrand = &brand
		}
		if card.Country != "" {
			payment.CardCountry = &card.Country
		}
		if card.ExpMonth != 0 {
			payment.CardExpMonth = &card.ExpMonth
		}
		if card.ExpYear != 0 {
			payment.CardExpYear = &card.ExpYear
		}
		if card.Fingerprint != "" {
			payment.CardFingerprint = &card.Fingerprint
		}
		if card.Funding != "" {
			funding := string(card.Funding)
			payment.CardFunding = &funding
		}
		if card.Last4 != "" {
			payment.CardLast4 = &card.Last4
		}
		if card.Network != "" {
			network := string(card.Network)
			payment.CardNetwork = &network
		}
	}

	// Handle JSON fields
	if charge.FraudDetails != nil {
		if fraudDetailsJSON, err := json.Marshal(charge.FraudDetails); err == nil {
			payment.FraudDetails = fraudDetailsJSON
		} else {
			payment.FraudDetails = []byte("{}")
		}
	} else {
		payment.FraudDetails = []byte("{}")
	}

	if charge.Metadata != nil {
		if metadataJSON, err := json.Marshal(charge.Metadata); err == nil {
			payment.Metadata = metadataJSON
		} else {
			payment.Metadata = []byte("{}")
		}
	} else {
		payment.Metadata = []byte("{}")
	}

	if charge.Shipping != nil {
		if shippingJSON, err := json.Marshal(charge.Shipping); err == nil {
			payment.Shipping = shippingJSON
		} else {
			payment.Shipping = []byte("{}")
		}
	} else {
		payment.Shipping = []byte("{}")
	}

	if charge.TransferData != nil {
		if transferDataJSON, err := json.Marshal(charge.TransferData); err == nil {
			payment.TransferData = transferDataJSON
		} else {
			payment.TransferData = []byte("{}")
		}
	} else {
		payment.TransferData = []byte("{}")
	}

	return model.CreateStripePayment(payment)
}
