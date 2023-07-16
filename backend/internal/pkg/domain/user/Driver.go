package user

import (
	"backend/internal/pkg/domain/cpf"
)

type Driver struct {
	ID       string
	Name     string
	Email    string
	CarPlate string
	Cpf      cpf.Cpf
}

func NewDriver(id string, name string, email string, cpf cpf.Cpf, carPlate string) Driver {
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
