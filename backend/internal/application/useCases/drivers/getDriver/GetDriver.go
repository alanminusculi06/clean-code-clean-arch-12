package getDriver

import (
	"backend/internal/application/repository"
	"backend/internal/pkg/domain"
)

type GetDriver struct {
	driverDatabase repository.DriverRepository
}

func NewGetDriverUseCase(driverDatabase repository.DriverRepository) GetDriver {
	return GetDriver{driverDatabase: driverDatabase}
}

func (useCase GetDriver) Execute(input Input) (*Output, *domain.ApiError) {
	driver, apiErr := useCase.driverDatabase.Get(input.ID)
	if apiErr != nil {
		return nil, apiErr
	}

	return &Output{
		ID:       driver.ID,
		Name:     driver.Name,
		Email:    driver.Email.Address,
		Cpf:      driver.Cpf.Number,
		CarPlate: driver.CarPlate,
	}, nil
}

type Input struct {
	ID string
}

type Output struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
	CarPlate string `json:"car-plate"`
}
