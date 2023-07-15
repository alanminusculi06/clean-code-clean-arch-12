package segment

import "time"

type Segment struct {
	Distance float64
	Date     time.Time
}

func NewSegment(distance float64, date time.Time) Segment {
	return Segment{Distance: distance, Date: date}
}

func (segment *Segment) IsOvernight() bool {
	return segment.Date.Hour() >= 22 || segment.Date.Hour() <= 6
}

func (segment *Segment) IsSunday() bool {
	if weekday := segment.Date.Weekday(); weekday == time.Sunday {
		return true
	}
	return false
}
