package fare

import "backend/internal/pkg/domain/segment"

const (
	sundayFare = 2.9
)

type SundayFare struct {
}

func NewSundayFare() Calculator {
	return &SundayFare{}
}

func (fare SundayFare) Execute(segment segment.Segment) float64 {
	return segment.Distance * sundayFare
}
