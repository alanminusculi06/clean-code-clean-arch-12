package drivers

import (
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/user"
	"github.com/google/uuid"
)

type CreateDriver struct {
}

func NewCreateDriverUseCase() CreateDriver {
	return CreateDriver{}
}

func (useCase CreateDriver) Execute(input Input) (*Output, *domain.ApiError) {
	driver := user.NewDriver(uuid.New().String(), input.Name, input.Email, cpf.NewCpf(input.Cpf), input.CarPlate)
	if cpfIsValid := driver.GetCpf().IsValid(); !cpfIsValid {
		return nil, domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", "")
	}

	//todo insert

	return &Output{
		ID:       driver.ID,
		Name:     driver.Name,
		Email:    driver.Email,
		Cpf:      driver.Cpf.Number,
		CarPlate: driver.CarPlate,
	}, nil
}

type Input struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
	CarPlate string `json:"car-plate"`
}

type Output struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
	CarPlate string `json:"car-plate"`
}
