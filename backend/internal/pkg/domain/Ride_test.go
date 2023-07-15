package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestRide_CalculatePrice(t *testing.T) {
	type fields struct {
		Segments []Segment
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
		want1  *ApiError
	}{
		{
			name: "Given a segment with negative distance when calculate ride price then return error",
			fields: fields{
				Segments: []Segment{
					{Distance: -10, Date: time.Date(2023, 7, 15, 14, 30, 00, 000, time.Local)},
				},
			},
			want:  0.0,
			want1: NewUnprocessableEntityError("error_negative_distance", "Distance cannot be negative", ""),
		},
		{
			name: "Given an over night not sunday segment when calculate ride price then return price",
			fields: fields{
				Segments: []Segment{
					{Distance: 10, Date: time.Date(2023, 7, 15, 22, 30, 00, 000, time.Local)},
				},
			},
			want:  39.0,
			want1: nil,
		},
		{
			name: "Given an over night sunday segment when calculate ride price then return price",
			fields: fields{
				Segments: []Segment{
					{Distance: 10, Date: time.Date(2023, 7, 16, 22, 30, 00, 000, time.Local)},
				},
			},
			want:  50.0,
			want1: nil,
		},
		{
			name: "Given a not over night sunday segment when calculate ride price then return price",
			fields: fields{
				Segments: []Segment{
					{Distance: 10, Date: time.Date(2023, 7, 16, 14, 30, 00, 000, time.Local)},
				},
			},
			want:  29.0,
			want1: nil,
		},
		{
			name: "Given a not over night not sunday segment when calculate ride price then return price",
			fields: fields{
				Segments: []Segment{
					{Distance: 10, Date: time.Date(2023, 7, 15, 14, 30, 00, 000, time.Local)},
				},
			},
			want:  21.0,
			want1: nil,
		},
		{
			name: "Given a less than min price segment when calculate ride price then return min price",
			fields: fields{
				Segments: []Segment{
					{Distance: 1, Date: time.Date(2023, 7, 15, 14, 30, 00, 000, time.Local)},
				},
			},
			want:  10.0,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ride := &Ride{}
			for _, segment := range tt.fields.Segments {
				ride.AddSegment(segment.Distance, segment.Date)
			}

			got, got1 := ride.CalculatePrice()
			if got != tt.want {
				t.Errorf("CalculatePrice() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CalculatePrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
