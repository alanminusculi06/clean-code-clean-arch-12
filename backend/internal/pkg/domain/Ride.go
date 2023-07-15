package domain

import (
	"backend/internal/pkg/shared"
	"time"
)

type Ride struct {
	Segments []Segment
}

func (ride *Ride) AddSegment(distance float64, date time.Time) {
	ride.Segments = append(ride.Segments, NewSegment(distance, date))
}

func (ride *Ride) CalculatePrice() (float64, *ApiError) {
	price := 0.0

	for _, segment := range ride.Segments {
		if segment.Distance < 0 {
			return 0.0, NewUnprocessableEntityError("error_negative_distance", "Distance cannot be negative", "")
		}
		if segment.IsOvernight() && !segment.IsSunday() {
			price += segment.Distance * shared.OvernightFare
		}
		if segment.IsOvernight() && segment.IsSunday() {
			price += segment.Distance * shared.OvernightSundayFare
		}
		if !segment.IsOvernight() && segment.IsSunday() {
			price += segment.Distance * shared.SundayFare
		}
		if !segment.IsOvernight() && !segment.IsSunday() {
			price += segment.Distance * shared.NormalFare
		}
	}

	if price < shared.MinPrice {
		return shared.MinPrice, nil
	}

	return price, nil
}
