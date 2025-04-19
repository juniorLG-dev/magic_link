package smtp

import (
	gomail "gopkg.in/gomail.v2"

	"fmt"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = 587
)

type senderCredentials struct {
	email 	 string
	password string
}

func NewSenderCredentials(email, password string) *senderCredentials {
	return &senderCredentials{
		email: email,
		password: password,
	}
}

func (sc *senderCredentials) SendEmail(recipient, code string) error {
	message := gomail.NewMessage()
	baseURL := "http://localhost:8080/verify"
	messageBody := fmt.Sprintf("<h1>Click on the link to validate your email:<strong>%s/%s</strong></h1>", baseURL, code)

	message.SetHeader("From", sc.email)
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", "Email verification")
	message.SetBody("text/html", messageBody)

	smtpClient := gomail.NewDialer(smtpHost, smtpPort, sc.email, sc.password)

	if err := smtpClient.DialAndSend(message); err != nil {
		return err
	}

	return nil
}

