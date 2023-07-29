package fare

import "backend/internal/pkg/domain/segment"

const (
	overnightSundayFare = 5
)

type OvernightSundayFare struct {
}

func NewOvernightSundayFare() Calculator {
	return &OvernightSundayFare{}
}

func (fare OvernightSundayFare) Execute(segment segment.Segment) float64 {
	return segment.Distance * overnightSundayFare
}
