package rides

import (
	"backend/internal/pkg/domain"
	rideDomain "backend/internal/pkg/domain/ride"
	"fmt"
	"time"
)

const DateTimeLayout = "2006-01-02T15:04:05"

type CalculateRide struct {
}

func NewCalculateRideUseCase() CalculateRide {
	return CalculateRide{}
}

func (useCase CalculateRide) Execute(input Input) (*Output, *domain.ApiError) {
	ride := rideDomain.Ride{}
	for i, segment := range input.Segments {
		date, err := time.Parse(DateTimeLayout, segment.Date)
		if err != nil {
			return nil, domain.NewUnprocessableEntityError("error_parse_date", fmt.Sprintf("Segment #%v: invalid date", i), err.Error())
		}
		ride.AddSegment(segment.Distance, date)
	}

	price, apiErr := ride.CalculatePrice()
	if apiErr != nil {
		return nil, apiErr
	}

	return &Output{Price: price}, nil
}

type Input struct {
	Segments []InputSegment
}

type InputSegment struct {
	Distance float64 `json:"distance"`
	Date     string  `json:"date"`
}

type Output struct {
	Price float64 `json:"price"`
}
