package infrastructure

import (
	"bytes"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

const (
	CONFIG_SMTP_HOST     = "smtp.zoho.com"
	CONFIG_SMTP_PORT     = 587
	CONFIG_SENDER_NAME   = "no_reply@pemudapeduli.org"
	CONFIG_AUTH_EMAIL    = "no_reply@pemudapeduli.org"
	CONFIG_AUTH_PASSWORD = "bismill@hPP2020"
)

// Mailer ...
type Mailer struct {
	from    string
	to      []string
	subject string
	body    string
}

// NewMailer ...
func NewMailer(to []string, subject, message string) *Mailer {
	return &Mailer{
		to:      to,
		subject: subject,
		body:    message,
	}
}

func (r *Mailer) SendEmail() error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", r.to...)
	// mailer.SetAddressHeader("Cc", "")
	mailer.SetHeader("Subject", r.subject)
	mailer.SetBody("text/html", r.body)
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,     // CONFIG_SMTP_HOST
		587,                  // CONFIG_SMTP_PORT
		CONFIG_AUTH_EMAIL,    // CONFIG_AUTH_EMAIL
		CONFIG_AUTH_PASSWORD, // CONFIG_AUTH_PASSWORD
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

func (r *Mailer) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
