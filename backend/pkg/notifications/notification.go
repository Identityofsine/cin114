package notifications

import (
	"grapplelogs/service/mail/template"
	"os"

	gomail "gopkg.in/mail.v2"
)

const (
	MAIL_VERIFICATION   = "mail-verification"
	MAIL_PASSWORD_RESET = "mail-password-reset"
	//mail layout -- maybe load from a file
	HTML_VERIFICATION = `email/verify.tmpl`
	HTML_WELCOME      = `email/welcome.tmpl`
)

var (
	mailClient *MailClient = nil
	templates  *template.TemplateStore
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
	email      Email
}

type MailClientProps struct {
	SmtpServer string `json:"smtp_server"`
	SmtpPort   int    `json:"smtp_port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
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
		log.Error("Mail Client not initialized due to missing fields")
		return nil
	}

	//template stuff
	templates = template.NewTemplateStore()
	//load templates
	templates.LoadTemplates("templates")

	mailClient = &MailClient{
		smtpServer: props.SmtpServer,
		smtpPort:   props.SmtpPort,
		username:   props.Username,
		password:   props.Password,
	}

	log.Debug("Mail Client Initialized: %s:%d", mailClient.smtpServer, mailClient.smtpPort)
	return mailClient
}

func (MailClient) SetEmailObject(email Email) {
	mailClient.email = email
}

func (MailClient) send() error {
	m := gomail.NewMessage()
	m.SetHeader("From", mailClient.username)
	m.SetHeader("To", mailClient.email.To)
	m.SetHeader("Subject", mailClient.email.Subject)
	m.SetBody("text/html", mailClient.email.Body)

	d := gomail.NewDialer(mailClient.smtpServer, mailClient.smtpPort, mailClient.username, mailClient.password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
