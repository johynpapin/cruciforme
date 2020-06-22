package email

import (
	"github.com/matcornic/hermes/v2"
)

type VerificationEmail struct {
	NewFormConfirmation bool
	Link                string
}

func (e *VerificationEmail) subject() string {
	return "[Cruciforme] Verify your account"
}

func (e *VerificationEmail) hermesEmail() *hermes.Email {
	return &hermes.Email{
		Body: hermes.Body{
			Title: "Welcome to Cruciforme!",
			Intros: []string{
				"Before you can use Cruciforme, please check your email address.",
				"Why, you may ask? It allows to avoid that someone uses Cruciforme to spam you.",
				"If you have not created an account on Cruciforme, please just ignore this email.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To verify your email address, please click here:",
					Button: hermes.Button{
						Text: "Verify your account",
						Link: e.Link,
					},
				},
			},
		},
	}
}
