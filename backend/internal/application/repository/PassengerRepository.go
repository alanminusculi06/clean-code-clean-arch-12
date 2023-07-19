package repository

import (
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/user"
)

type PassengerRepository interface {
	Save(passenger user.Passenger) (*user.Passenger, *domain.ApiError)
	Get(passengerID string) (*user.Passenger, *domain.ApiError)
}
