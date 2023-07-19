package repository

import (
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/user"
)

type DriverRepository interface {
	Save(driver user.Driver) (*user.Driver, *domain.ApiError)
	Get(driverID string) (*user.Driver, *domain.ApiError)
}
