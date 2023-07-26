package createPassenger

import (
	"backend/internal/application/repository"
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/email"
	"backend/internal/pkg/domain/user"
	"github.com/google/uuid"
)

type CreatePassenger struct {
	passengerDatabase repository.PassengerRepository
}

func NewCreatePassengerUseCase(passengerDatabase repository.PassengerRepository) CreatePassenger {
	return CreatePassenger{passengerDatabase: passengerDatabase}
}

func (useCase CreatePassenger) Execute(input Input) (*Output, *domain.ApiError) {
	passenger := user.NewPassenger(uuid.New().String(), input.Name, email.NewEmail(input.Email), cpf.NewCpf(input.Cpf))

	if cpfIsValid := passenger.GetCpf().IsValid(); !cpfIsValid {
		return nil, domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", "")
	}

	if emailIsValid := passenger.GetEmail().IsValid(); !emailIsValid {
		return nil, domain.NewUnprocessableEntityError("invalid_email", "Given e-mail is not valid.", "")
	}

	insertedPassenger, apiErr := useCase.passengerDatabase.Save(passenger)
	if apiErr != nil {
		return nil, apiErr
	}

	return &Output{
		ID:    insertedPassenger.ID,
		Name:  insertedPassenger.Name,
		Email: insertedPassenger.Email.Address,
		Cpf:   insertedPassenger.Cpf.Number,
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
