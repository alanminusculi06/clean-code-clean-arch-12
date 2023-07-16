package passengers

import (
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/user"
	"github.com/google/uuid"
)

type CreatePassenger struct {
}

func NewCreatePassengerUseCase() CreatePassenger {
	return CreatePassenger{}
}

func (useCase CreatePassenger) Execute(input Input) (*Output, *domain.ApiError) {
	passenger := user.NewPassenger(uuid.New().String(), input.Name, input.Email, cpf.NewCpf(input.Cpf))
	if cpfIsValid := passenger.GetCpf().IsValid(); !cpfIsValid {
		return nil, domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", "")
	}

	//todo insert

	return &Output{
		ID:    passenger.ID,
		Name:  passenger.Name,
		Email: passenger.Email,
		Cpf:   passenger.Cpf.Number,
	}, nil
}

type Input struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}

type Output struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}
