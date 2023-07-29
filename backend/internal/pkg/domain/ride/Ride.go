package ride

import (
	"backend/internal/pkg/domain"
	fareCalculator "backend/internal/pkg/domain/fare"
	"backend/internal/pkg/domain/segment"
	"backend/internal/pkg/shared"
	"time"
)

type Ride struct {
	Segments []segment.Segment
}

func (ride *Ride) AddSegment(distance float64, date time.Time) {
	ride.Segments = append(ride.Segments, segment.NewSegment(distance, date))
}

func (ride *Ride) CalculatePrice() (float64, *domain.ApiError) {
	price := 0.0

	for _, segmentItem := range ride.Segments {
		if isValid := isValidSegment(segmentItem); !isValid {
			return 0.0, domain.NewUnprocessableEntityError("error_negative_distance", "Distance cannot be negative", "")
		}

		if fare, err := fareCalculator.GetFareCalculator(segmentItem); err != nil {
			return 0.0, err
		} else {
			price += fare.Execute(segmentItem)
		}
	}

	return priceOrMinValue(price), nil
}

func isValidSegment(segment segment.Segment) bool {
	return segment.Distance > 0
}

func priceOrMinValue(price float64) float64 {
	if price < shared.MinPrice {
		return shared.MinPrice
	}
	return price
}
