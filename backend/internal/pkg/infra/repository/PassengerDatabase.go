package repository

import (
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/user"
	"database/sql"
	"fmt"
)

type PassengerDatabase struct {
	db *sql.DB
}

func NewPassengerDatabase(db *sql.DB) PassengerDatabase {
	return PassengerDatabase{db: db}
}

func (passengerDatabase PassengerDatabase) Save(passenger user.Passenger) (*user.Passenger, *domain.ApiError) {
	_, err := passengerDatabase.db.Exec("INSERT INTO passenger (`id`,`name`,`email`,`cpf`) VALUES (?,?,?,?);", passenger.ID, passenger.Name, passenger.Email.Address, passenger.Cpf.Number)
	if err != nil {
		return nil, domain.NewInternalServerError("error_save_passenger", "Unexpected error while saving passenger", err.Error())
	}
	return &passenger, nil
}

func (passengerDatabase PassengerDatabase) Get(passengerID string) (*user.Passenger, *domain.ApiError) {
	var passenger user.Passenger
	row := passengerDatabase.db.QueryRow("SELECT * FROM passenger WHERE id = ?", passengerID)
	if err := row.Scan(&passenger.ID, &passenger.Name, &passenger.Email.Address, &passenger.Cpf.Number); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewNotFoundError("error_get_passenger", fmt.Sprintf("Not found passenger with ID %s", passengerID), err.Error())
		}
		return nil, domain.NewInternalServerError("error_get_passenger", "Unexpected error on get passenger", err.Error())
	}
	return &passenger, nil
}
