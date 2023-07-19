package main

import (
	"backend/internal/application/repository"
	createDriver2 "backend/internal/application/useCases/drivers/createDriver"
	getDriver2 "backend/internal/application/useCases/drivers/getDriver"
	createPassenger2 "backend/internal/application/useCases/passengers/createPassenger"
	getDriver3 "backend/internal/application/useCases/passengers/getDriver"
	"backend/internal/application/useCases/rides"
	"backend/internal/pkg/domain"
	"backend/internal/pkg/infra/database"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB
var passengerDatabase repository.PassengerRepository
var driverDatabase repository.DriverRepository

func main() {
	db = configDatabase()
	passengerDatabase = database.NewPassengerDatabase(db)
	driverDatabase = database.NewDriverDatabase(db)

	router := gin.Default()
	router.POST("/rides/calculate-price", calculateRidePrice)
	router.POST("/drivers", createDriver)
	router.GET("/drivers/:id", getDriver)
	router.POST("/passengers", createPassenger)
	router.GET("/passengers/:id", getPassenger)
	_ = router.Run("localhost:8080")
}

func getPassenger(context *gin.Context) {
	output, apiErr := getDriver3.NewGetPassengerUseCase(passengerDatabase).Execute(getDriver3.Input{ID: context.Param("id")})
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusOK, output)
}

func getDriver(context *gin.Context) {
	output, apiErr := getDriver2.NewGetDriverUseCase(driverDatabase).Execute(getDriver2.Input{ID: context.Param("id")})
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusOK, output)
}

func createPassenger(context *gin.Context) {
	var input createPassenger2.Input
	if apiErr := bindJson(context, &input); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	output, apiErr := createPassenger2.NewCreatePassengerUseCase(passengerDatabase).Execute(input)
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusCreated, output)
}

func createDriver(context *gin.Context) {
	var input createDriver2.Input
	if apiErr := bindJson(context, &input); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	output, apiErr := createDriver2.NewCreateDriverUseCase(driverDatabase).Execute(input)
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusCreated, output)
}

func calculateRidePrice(context *gin.Context) {
	var input []rides.InputSegment
	if apiErr := bindJson(context, &input); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	output, apiErr := rides.NewCalculateRideUseCase().Execute(rides.Input{Segments: input})
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusOK, output)
}

func bindJson(context *gin.Context, obj interface{}) *domain.ApiError {
	if err := context.BindJSON(obj); err != nil {
		return domain.NewBindJsonError(err.Error())
	}
	return nil
}

func configDatabase() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "mysqlroot",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "ccca_t12",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("DB Connected!")

	return db
}
