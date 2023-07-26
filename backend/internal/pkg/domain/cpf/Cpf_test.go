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

func TestCpf_TestValidate(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Given an invalid CPF, when IsValid is called, then it should return false #1",
			args: args{
				str: "834.326.160-76",
			},
			want: false,
		},
		{
			name: "Given an invalid CPF, when IsValid is called, then it should return false #2",
			args: args{
				str: "999.999.999-99",
			},
			want: false,
		},
		{
			name: "Given an invalid CPF, when IsValid is called, then it should return false #3",
			args: args{
				str: "834.326.160",
			},
			want: false,
		},
		{
			name: "Given an invalid CPF, when IsValid is called, then it should return false #4",
			args: args{
				str: "",
			},
			want: false,
		},
		{
			name: "Given an invalid CPF, when IsValid is called, then it should return false #5",
			args: args{
				str: "undefined",
			},
			want: false,
		},
		{
			name: "Given a valid CPF, when IsValid is called, then it should return true #1",
			args: args{
				str: "834.326.160-74",
			},
			want: true,
		},
		{
			name: "Given a valid CPF, when IsValid is called, then it should return true #2",
			args: args{
				str: "745.878.878-03",
			},
			want: true,
		},
		{
			name: "Given a valid CPF, when IsValid is called, then it should return true #3",
			args: args{
				str: "871.756.595-20",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.str); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
