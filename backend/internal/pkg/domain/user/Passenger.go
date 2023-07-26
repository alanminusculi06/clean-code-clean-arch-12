package user

import (
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/email"
)

type Passenger struct {
	ID    string
	Name  string
	Email email.Email
	Cpf   cpf.Cpf
}

func NewPassenger(id string, name string, email email.Email, cpf cpf.Cpf) Passenger {
	return Passenger{
		ID:    id,
		Name:  name,
		Email: email,
		Cpf:   cpf,
	}
}

func (passenger Passenger) GetCpf() cpf.Cpf {
	return passenger.Cpf
}

func (passenger Passenger) GetEmail() email.Email {
	return passenger.Email
}
