package user

import (
	"backend/internal/pkg/domain/cpf"
)

type Passenger struct {
	ID    string
	Name  string
	Email string
	Cpf   cpf.Cpf
}

func NewPassenger(id string, name string, email string, cpf cpf.Cpf) Passenger {
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
