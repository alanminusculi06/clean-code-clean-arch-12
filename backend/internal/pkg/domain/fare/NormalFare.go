package fare

import "backend/internal/pkg/domain/segment"

const (
	normalFare = 2.1
)

type NormalFare struct {
}

func NewNormalFare() Calculator {
	return &NormalFare{}
}

func (fare NormalFare) Execute(segment segment.Segment) float64 {
	return segment.Distance * normalFare
}
