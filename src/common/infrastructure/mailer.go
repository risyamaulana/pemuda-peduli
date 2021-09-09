package infrastructure

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const (
	CONFIG_SMTP_HOST     = "smtp.zoho.com"
	CONFIG_SMTP_PORT     = 587
	CONFIG_SENDER_NAME   = "Pemuda Peduli <no_reply@pemudapeduli.org>"
	CONFIG_AUTH_EMAIL    = "no_reply@pemudapeduli.org"
	CONFIG_AUTH_PASSWORD = "bismill@hPP2020"
)

// Mailer ...
type Mailer struct{}

// NewMailer ...
func NewMailer() *Mailer {
	return &Mailer{}
}

func SendMail(to []string, subject, message string) (err error) {
	cc := []string{""}

	err = sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("Mail sent!")
	return
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	log.Println("Body Message Email : ", body)

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
