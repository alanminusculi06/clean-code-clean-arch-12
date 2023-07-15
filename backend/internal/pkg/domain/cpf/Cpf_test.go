package cpf

import (
	"testing"
)

func TestCpf_IsValid(t *testing.T) {
	type fields struct {
		Number string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Given an invalid CPF, when IsValid is called, then it should return false",
			fields: fields{
				Number: "834.326.160-76",
			},
			want: false,
		},
		{
			name: "Given a valid CPF, when IsValid is called, then it should return true",
			fields: fields{
				Number: "871.756.595-20",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpf := NewCpf(tt.fields.Number)
			if got := cpf.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
