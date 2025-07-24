package notifications

import (
	"errors"
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	gomail "gopkg.in/mail.v2"
)

const (
	MAIL_VERIFICATION        = "mail-verification"
	MAIL_PASSWORD_RESET      = "mail-password-reset"
	MAIL_TICKET_CONFIRMATION = "mail-ticket-confirmation"
	//mail layout -- maybe load from a file
	HTML_VERIFICATION        = `email/verify.tmpl`
	HTML_TICKET_CONFIRMATION = `email/ticket-purchase-success.tmpl`
)

var (
	mailClient *MailClient = nil
	templates  *TemplateStore
)

type Email struct {
	To      string
	Subject string
	Body    string
}

type MailClient struct {
	smtpServer string
	smtpPort   int
	username   string
	password   string
	fromEmail  string
	fromName   string
	email      Email
}

type MailClientProps struct {
	SmtpServer string `json:"smtp_server"`
	SmtpPort   int    `json:"smtp_port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FromEmail  string `json:"from_email"`
	FromName   string `json:"from_name"`
}

// Ticket represents a single ticket for template rendering
type Ticket struct {
	Number string `json:"number"`
	Type   string `json:"type"`
	Price  string `json:"price"`
}

// TicketConfirmationData contains all data needed for ticket confirmation email
type TicketConfirmationData struct {
	EventTitle    string   `json:"event_title"`
	EventDate     string   `json:"event_date"`
	EventTime     string   `json:"event_time"`
	EventVenue    string   `json:"event_venue"`
	EventAddress  string   `json:"event_address"`
	Tickets       []Ticket `json:"tickets"`
	Subtotal      string   `json:"subtotal"`
	ProcessingFee string   `json:"processing_fee"`
	TaxAmount     string   `json:"tax_amount"`
	TotalAmount   string   `json:"total_amount"`
	PaymentMethod string   `json:"payment_method"`
	TransactionID string   `json:"transaction_id"`
	PurchaseDate  string   `json:"purchase_date"`
	CalendarURL   string   `json:"calendar_url"`
}

// InitMailClient - this will initialize the mail client - also will act as a constructor
func InitMailClient(props *MailClientProps) *MailClient {
	if mailClient != nil {
		return mailClient
	}

	if props == nil {
		return nil
	}

	if props.SmtpServer == "" || props.SmtpPort == 0 || props.Username == "" || props.Password == "" {
		storedlogs.LogError("Mail Client not initialized due to missing fields", errors.New("missing fields in MailClientProps"))
		return nil
	}

	//template stuff
	templates = NewTemplateStore()
	//load templates
	templates.LoadTemplates("templates")

	mailClient = &MailClient{
		smtpServer: props.SmtpServer,
		smtpPort:   props.SmtpPort,
		username:   props.Username,
		password:   props.Password,
		fromEmail:  props.FromEmail,
		fromName:   props.FromName,
	}

	storedlogs.LogDebug("Mail Client Initialized")
	return mailClient
}

// InitMailClientFromConfig - initializes the mail client using the config package
func InitMailClientFromConfig() *MailClient {
	if mailClient != nil {
		return mailClient
	}

	mailConfig := config.GetMailSettings()
	if mailConfig == nil {
		storedlogs.LogError("Mail Client not initialized due to missing mail configuration", errors.New("mail configuration not found"))
		return nil
	}

	props := &MailClientProps{
		SmtpServer: mailConfig.SMTPServer,
		SmtpPort:   mailConfig.SMTPPort,
		Username:   mailConfig.SMTPUsername,
		Password:   mailConfig.SMTPPassword,
		FromEmail:  mailConfig.FromEmail,
		FromName:   mailConfig.FromName,
	}

	fmt.Println("Initializing Mail Client with config:", props)

	return InitMailClient(props)
}

func (MailClient) SetEmailObject(email Email) {
	mailClient.email = email
}

func (MailClient) send() error {
	m := gomail.NewMessage()

	// Use configured from email and name
	if mailClient.fromName != "" {
		m.SetHeader("From", m.FormatAddress(mailClient.fromEmail, mailClient.fromName))
	} else {
		m.SetHeader("From", mailClient.fromEmail)
	}

	m.SetHeader("To", mailClient.email.To)
	m.SetHeader("Subject", mailClient.email.Subject)
	m.SetBody("text/html", mailClient.email.Body)

	d := gomail.NewDialer(mailClient.smtpServer, mailClient.smtpPort, mailClient.username, mailClient.password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// SendEmail - convenience function to send an email using the configured mail client
func SendEmail(to, subject, body string) error {
	if mailClient == nil {
		// Try to initialize from config if not already done
		if InitMailClientFromConfig() == nil {
			return errors.New("mail client not initialized")
		}
	}

	email := Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	mailClient.SetEmailObject(email)
	return mailClient.send()
}

// SendTicketConfirmationEmail - sends a ticket confirmation email using the template
func SendTicketConfirmationEmail(to string, data TicketConfirmationData) error {
	if mailClient == nil {
		// Try to initialize from config if not already done
		if InitMailClientFromConfig() == nil {
			return errors.New("mail client not initialized")
		}
	}

	if templates == nil {
		return errors.New("templates not initialized")
	}

	// Convert struct to H map for template
	templateData := H{
		"EventTitle":    data.EventTitle,
		"EventDate":     data.EventDate,
		"EventTime":     data.EventTime,
		"EventVenue":    data.EventVenue,
		"EventAddress":  data.EventAddress,
		"Tickets":       data.Tickets,
		"Subtotal":      data.Subtotal,
		"ProcessingFee": data.ProcessingFee,
		"TaxAmount":     data.TaxAmount,
		"TotalAmount":   data.TotalAmount,
		"PaymentMethod": data.PaymentMethod,
		"TransactionID": data.TransactionID,
		"PurchaseDate":  data.PurchaseDate,
		"CalendarURL":   data.CalendarURL,
	}

	// Parse the template
	body, err := templates.ParseTemplate(HTML_TICKET_CONFIRMATION, templateData)
	if err != nil {
		storedlogs.LogError("Failed to parse ticket confirmation template", err)
		return err
	}

	// Send the email
	subject := "Ticket Purchase Confirmation - CIN114"
	return SendEmail(to, subject, body)
}
