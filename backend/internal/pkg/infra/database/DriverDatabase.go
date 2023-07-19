package database

import (
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/user"
	"database/sql"
	"fmt"
)

type DriverDatabase struct {
	db *sql.DB
}

func NewDriverDatabase(db *sql.DB) DriverDatabase {
	return DriverDatabase{db: db}
}

func (driverDatabase DriverDatabase) Save(driver user.Driver) (*user.Driver, *domain.ApiError) {
	_, err := driverDatabase.db.Exec("INSERT INTO driver(`id`,`name`,`email`,`car_plate`,`cpf`) VALUES (?,?,?,?,?);", driver.ID, driver.Name, driver.Email, driver.CarPlate, driver.Cpf.Number)
	if err != nil {
		return nil, domain.NewInternalServerError("error_save_driver", "Unexpected error while saving driver", err.Error())
	}
	return &driver, nil
}

func (driverDatabase DriverDatabase) Get(driverID string) (*user.Driver, *domain.ApiError) {
	var driver user.Driver
	row := driverDatabase.db.QueryRow("SELECT * FROM driver WHERE id = ?", driverID)
	if err := row.Scan(&driver.ID, &driver.Name, &driver.Email, &driver.CarPlate, &driver.Cpf.Number); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewNotFoundError("error_get_driver", fmt.Sprintf("Not found driver with ID %s", driverID), err.Error())
		}
		return nil, domain.NewInternalServerError("error_get_driver", "Unexpected error on get driver", err.Error())
	}
	return &driver, nil
}
