package main

import (
	"backend/internal/application/dto"
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/ride"
	"backend/internal/pkg/domain/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.POST("/rides/calculate-price", calculateRidePrice)
	router.POST("/drivers", createDriver)
	router.POST("/passengers", createPassenger)
	_ = router.Run("localhost:8080")
}

func createPassenger(context *gin.Context) {
	var passengerDTO dto.PassengerDTO
	if apiErr := bindJson(context, &passengerDTO); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	passenger := user.NewPassenger(uuid.New().String(), passengerDTO.Name, passengerDTO.Email, cpf.NewCpf(passengerDTO.Cpf))
	if apiErr := validateCpf(passenger); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	passengerDTO.ID = passenger.ID

	context.IndentedJSON(http.StatusCreated, passengerDTO)
}

func createDriver(context *gin.Context) {
	var driverDTO dto.DriverDTO
	if apiErr := bindJson(context, &driverDTO); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	driver := user.NewDriver(uuid.New().String(), driverDTO.Name, cpf.NewCpf(driverDTO.Cpf), driverDTO.CarPlate)
	if apiErr := validateCpf(driver); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	driverDTO.ID = driver.ID

	context.IndentedJSON(http.StatusCreated, driverDTO)
}

func calculateRidePrice(context *gin.Context) {
	var rideDTO dto.RideDTO
	if apiErr := bindJson(context, &rideDTO); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	ride := ride.Ride{}
	for i, segment := range rideDTO.Segments {
		layout := "2006-01-02T15:04:05"
		date, err := time.Parse(layout, segment.Date)
		if err != nil {
			apiErr := domain.NewUnprocessableEntityError("error_parse_date", fmt.Sprintf("Segment #%v: invalid date", i), err.Error())
			context.IndentedJSON(http.StatusBadRequest, apiErr)
			return
		}

		ride.AddSegment(segment.Distance, date)
	}

	price, apiErr := ride.CalculatePrice()
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusOK, dto.CalculateRidePriceResultDTO{Price: price})
}

func bindJson(context *gin.Context, obj interface{}) *domain.ApiError {
	if err := context.BindJSON(obj); err != nil {
		return domain.NewBindJsonError(err.Error())
	}
	return nil
}

func validateCpf(user user.IUser) *domain.ApiError {
	if cpfIsValid := user.GetCpf().IsValid(); !cpfIsValid {
		return domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", "")
	}
	return nil
}
