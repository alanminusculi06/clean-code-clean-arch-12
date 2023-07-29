package fare

import (
	"backend/internal/pkg/domain/segment"
)

const (
	overnightFare = 3.90
)

type OvernightFare struct {
}

func NewOvernightFare() Calculator {
	return &OvernightFare{}
}

func (fare OvernightFare) Execute(segment segment.Segment) float64 {
	return segment.Distance * overnightFare
}
