package fare

import "backend/internal/pkg/domain/segment"

type Calculator interface {
	Execute(segment segment.Segment) float64
}
