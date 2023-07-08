package main

import (
	"backend/internal/application/dto"
	"backend/internal/pkg/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.POST("/ride/calculate-price", calculateRidePrice)
	router.Run("localhost:8080")
}

func calculateRidePrice(context *gin.Context) {
	var rideDTO dto.RideDTO
	if err := context.BindJSON(&rideDTO); err != nil {
		apiErr := domain.NewBindJsonError(err.Error())
		context.IndentedJSON(http.StatusBadRequest, apiErr)
		return
	}

	ride := domain.Ride{}
	for i, segment := range rideDTO.Segments {
		layout := "2006-01-02T15:04:05.000Z"
		date, err := time.Parse(layout, segment.Date)
		if err != nil {
			apiErr := domain.NewApiError("error_parse_date", fmt.Sprintf("Segment #%v: invalid date", i), err.Error())
			context.IndentedJSON(http.StatusBadRequest, apiErr)
			return
		}

		ride.AddSegment(segment.Distance, date)
	}

	price := ride.CalculatePrice()
	context.IndentedJSON(http.StatusOK, price)
}