package user

import "backend/internal/pkg/domain/cpf"

type IUser interface {
	GetCpf() cpf.Cpf
}
