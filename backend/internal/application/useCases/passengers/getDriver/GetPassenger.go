package getDriver

import (
	"backend/internal/application/repository"
	"backend/internal/pkg/domain"
)

type GetPassenger struct {
	passengerDatabase repository.PassengerRepository
}

func NewGetPassengerUseCase(passengerDatabase repository.PassengerRepository) GetPassenger {
	return GetPassenger{passengerDatabase: passengerDatabase}
}

func (useCase GetPassenger) Execute(input Input) (*Output, *domain.ApiError) {
	passenger, apiErr := useCase.passengerDatabase.Get(input.ID)
	if apiErr != nil {
		return nil, apiErr
	}

	return &Output{
		ID:    passenger.ID,
		Name:  passenger.Name,
		Email: passenger.Email.Address,
		Cpf:   passenger.Cpf.Number,
	}, nil
}

type Input struct {
	ID string
}

type Output struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}
