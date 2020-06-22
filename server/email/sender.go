package email

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type email interface {
	subject() string
	hermesEmail() *hermes.Email
}

type EmailAddress struct {
	Name  string
	Email string
}

type Sender struct {
	SendGridApiKey string
	From           EmailAddress
	hermes         hermes.Hermes
}

func NewSender(SendGridApiKey string, From EmailAddress) *Sender {
	sender := &Sender{
		SendGridApiKey: SendGridApiKey,
		From:           From,
	}

	sender.hermes = hermes.Hermes{
		Product: hermes.Product{
			Name:      "Cruciforme",
			Link:      "https://crucifor.me",
			Copyright: "Cruciforme is free software under the GPLv3 license.",
		},
	}

	return sender
}

func (s *Sender) Send(to EmailAddress, email email) error {
	return s.sendWithSendGrid(to, email)
}

func (s *Sender) sendWithSendGrid(to EmailAddress, email email) error {
	fromEmail := mail.NewEmail(s.From.Name, s.From.Email)
	toEmail := mail.NewEmail(to.Name, to.Email)

	hermesEmail := email.hermesEmail()

	plainTextContent, err := s.hermes.GeneratePlainText(*hermesEmail)
	if err != nil {
		return err
	}
	htmlContent, err := s.hermes.GenerateHTML(*hermesEmail)
	if err != nil {
		return err
	}

	message := mail.NewSingleEmail(fromEmail, email.subject(), toEmail, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(s.SendGridApiKey)

	_, err = client.Send(message)
	if err != nil {
		return fmt.Errorf("could not send the email using SendGrid: %w", err)
	}

	return nil
}
