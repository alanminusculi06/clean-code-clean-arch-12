package email

import "testing"

func TestEmail_IsValid(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Given an invalid email, when IsValid is called, then it should return false",
			fields: fields{
				Address: "test.com",
			},
			want: false,
		},
		{
			name: "Given an empty email, when IsValid is called, then it should return false",
			fields: fields{
				Address: "",
			},
			want: false,
		},
		{
			name: "Given a valid email, when IsValid is called, then it should return true",
			fields: fields{
				Address: "test@email.com",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := NewEmail(tt.fields.Address)
			if got := email.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
