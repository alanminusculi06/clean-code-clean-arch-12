package fare

import (
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/segment"
)

func GetFareCalculator(segment segment.Segment) (Calculator, *domain.ApiError) {
	if segment.IsOvernight() && !segment.IsSunday() {
		return NewOvernightFare(), nil
	}
	if segment.IsOvernight() && segment.IsSunday() {
		return NewOvernightSundayFare(), nil
	}
	if !segment.IsOvernight() && segment.IsSunday() {
		return NewSundayFare(), nil
	}
	if !segment.IsOvernight() && !segment.IsSunday() {
		return NewNormalFare(), nil
	}
	return nil, domain.NewUnprocessableEntityError("invalid_fare", "Invalid fare", "")
}
