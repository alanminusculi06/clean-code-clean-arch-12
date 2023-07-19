package createDriver

import (
	"backend/internal/application/repository"
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/user"
	"github.com/google/uuid"
)

type CreateDriver struct {
	driverDatabase repository.DriverRepository
}

func NewCreateDriverUseCase(driverDatabase repository.DriverRepository) CreateDriver {
	return CreateDriver{driverDatabase: driverDatabase}
}

func (useCase CreateDriver) Execute(input Input) (*Output, *domain.ApiError) {
	driver := user.NewDriver(uuid.New().String(), input.Name, input.Email, cpf.NewCpf(input.Cpf), input.CarPlate)
	if cpfIsValid := driver.GetCpf().IsValid(); !cpfIsValid {
		return nil, domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", "")
	}

	insertedDriver, apiErr := useCase.driverDatabase.Save(driver)
	if apiErr != nil {
		return nil, apiErr
	}

	return &Output{
		ID:       insertedDriver.ID,
		Name:     insertedDriver.Name,
		Email:    insertedDriver.Email,
		Cpf:      insertedDriver.Cpf.Number,
		CarPlate: insertedDriver.CarPlate,
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
