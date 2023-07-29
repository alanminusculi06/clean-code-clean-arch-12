package main

import (
	"backend/internal/application/repository"
	createDriver2 "backend/internal/application/useCases/drivers/createDriver"
	getDriver2 "backend/internal/application/useCases/drivers/getDriver"
	createPassenger2 "backend/internal/application/useCases/passengers/createPassenger"
	getPassenger2 "backend/internal/application/useCases/passengers/getPassenger"
	"backend/internal/application/useCases/rides"
	"backend/internal/pkg/infra/database"
	"backend/internal/pkg/infra/http"
	"database/sql"
)

var db *sql.DB
var passengerDatabase repository.PassengerRepository
var driverDatabase repository.DriverRepository

func main() {
	sqlDatabase := database.NewDataBase()
	db = sqlDatabase.Config()

	passengerDatabase = database.NewPassengerDatabase(db)
	driverDatabase = database.NewDriverDatabase(db)

	createDriver := createDriver2.NewCreateDriverUseCase(driverDatabase)
	createPassenger := createPassenger2.NewCreatePassengerUseCase(passengerDatabase)
	getDriver := getDriver2.NewGetDriverUseCase(driverDatabase)
	getPassenger := getPassenger2.NewGetPassengerUseCase(passengerDatabase)
	calculateRide := rides.NewCalculateRideUseCase()

	router := http.NewRouter(createDriver, createPassenger, getDriver, getPassenger, calculateRide)
	router.SetupRoutes()
}
