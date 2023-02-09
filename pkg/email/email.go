package email

import (
	"github.com/go-gomail/gomail"
)

type Email struct {
	From       string
	To         []string
	Subject    string
	Body       string
	Attachment []string
}

func (e *Email) SendEmail(host string, port int, sender, password string) error {
	email := gomail.NewMessage()
	email.SetHeader("From", e.From)
	email.SetHeader("To", e.To...)
	email.SetHeader("Subject", e.Subject)
	email.SetBody("text/plain;charset=UTF-8", e.Body)
	for _, fileName := range e.Attachment {
		email.Attach(fileName)
	}

	dialer := gomail.NewDialer(host, port, sender, password)
	return dialer.DialAndSend(email)
}
