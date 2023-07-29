package http

import (
	"backend/internal/application/useCases/drivers/createDriver"
	"backend/internal/application/useCases/drivers/getDriver"
	"backend/internal/application/useCases/passengers/createPassenger"
	"backend/internal/application/useCases/passengers/getPassenger"
	"backend/internal/application/useCases/rides"
	"backend/internal/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	createDriverUseCase    createDriver.CreateDriver
	createPassengerUseCase createPassenger.CreatePassenger
	getDriverUseCase       getDriver.GetDriver
	getPassengerUseCase    getPassenger.GetPassenger
	calculateRide          rides.CalculateRide
}

func NewRouter(
	createDriver createDriver.CreateDriver,
	createPassenger createPassenger.CreatePassenger,
	getDriver getDriver.GetDriver,
	getPassenger getPassenger.GetPassenger,
	calculateRide rides.CalculateRide) Router {
	return Router{
		createDriverUseCase:    createDriver,
		createPassengerUseCase: createPassenger,
		getDriverUseCase:       getDriver,
		getPassengerUseCase:    getPassenger,
		calculateRide:          calculateRide}
}

func (r Router) SetupRoutes() {
	router := gin.Default()
	router.POST("/rides/calculate-price", r.calculateRidePrice)
	router.POST("/drivers", r.createDriver)
	router.GET("/drivers/:id", r.getDriver)
	router.POST("/passengers", r.createPassenger)
	router.GET("/passengers/:id", r.getPassenger)
	_ = router.Run("localhost:8080")
}

func (r Router) calculateRidePrice(context *gin.Context) {
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

func (r Router) createDriver(context *gin.Context) {
	var input createDriver.Input
	if apiErr := bindJson(context, &input); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	output, apiErr := r.createDriverUseCase.Execute(input)
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusCreated, output)
}

func (r Router) getDriver(context *gin.Context) {
	output, apiErr := r.getDriverUseCase.Execute(getDriver.Input{ID: context.Param("id")})
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusOK, output)
}

func (r Router) createPassenger(context *gin.Context) {
	var input createPassenger.Input
	if apiErr := bindJson(context, &input); apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	output, apiErr := r.createPassengerUseCase.Execute(input)
	if apiErr != nil {
		context.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	context.IndentedJSON(http.StatusCreated, output)
}

func (r Router) getPassenger(context *gin.Context) {
	output, apiErr := r.getPassengerUseCase.Execute(getPassenger.Input{ID: context.Param("id")})
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
