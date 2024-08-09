package email

import (
	"fmt"
	"net/smtp"
	"os"
)

type Email struct {
	msg       string
	reccpient []string
	subject   string
}

func New(m string, r []string, s string) *Email {
	return &Email{m, r, s}
}

func (email *Email) SendMail() (*Email, error) {
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")

	to := email.reccpient

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	message := []byte("Subject: " + email.subject + "\r\n" + email.msg + "\r\n")
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return email, nil
}
