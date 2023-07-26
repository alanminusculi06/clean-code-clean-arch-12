package user

import (
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/email"
)

type Driver struct {
	ID       string
	Name     string
	Email    email.Email
	CarPlate string
	Cpf      cpf.Cpf
}

func NewDriver(id string, name string, email email.Email, cpf cpf.Cpf, carPlate string) Driver {
	return Driver{
		ID:       id,
		Name:     name,
		Email:    email,
		Cpf:      cpf,
		CarPlate: carPlate,
	}
}

func (driver Driver) GetCpf() cpf.Cpf {
	return driver.Cpf
}

func (driver Driver) GetEmail() email.Email {
	return driver.Email
}
