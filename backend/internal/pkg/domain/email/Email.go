package email

import "net/mail"

type Email struct {
	Address string
}

func NewEmail(address string) Email {
	return Email{Address: address}
}

func (email Email) IsValid() bool {
	return validate(email.Address)
}

func validate(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
