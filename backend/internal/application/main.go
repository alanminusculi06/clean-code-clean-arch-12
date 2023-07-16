package main

import (
	"backend/internal/application/useCases/drivers"
	"backend/internal/application/useCases/passengers"
	"backend/internal/application/useCases/rides"
	"backend/internal/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/rides/calculate-price", calculateRidePrice)
	router.POST("/drivers", createDriver)
	router.POST("/passengers", createPassenger)
	_ = router.Run("localhost:8080")
}

func createPassenger(context *gin.Context) {
	var input passengers.Input
	if apiErr := bindJson(context, &input); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	output, apiErr := passengers.NewCreatePassengerUseCase().Execute(input)
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusCreated, output)
}

func createDriver(context *gin.Context) {
	var input drivers.Input
	if apiErr := bindJson(context, &input); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	output, apiErr := drivers.NewCreateDriverUseCase().Execute(input)
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
