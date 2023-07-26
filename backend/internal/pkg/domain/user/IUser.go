package user

import (
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/email"
)

type IUser interface {
	GetCpf() cpf.Cpf
	GetEmail() email.Email
}
