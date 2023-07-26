package createPassenger

import (
	"backend/internal/application/repository"
	"backend/internal/pkg/domain"
	"backend/internal/pkg/domain/cpf"
	"backend/internal/pkg/domain/email"
	"backend/internal/pkg/domain/user"
	"backend/internal/pkg/infra/database/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestCreatePassenger_Execute(t *testing.T) {
	type fields struct {
		passengerDatabase repository.PassengerRepository
	}
	type args struct {
		input Input
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		expectedPassenger *Output
		expectedError     *domain.ApiError
	}{
		{
			name: "Given a valid input When execute Then return nil",
			fields: fields{
				passengerDatabase: func() repository.PassengerRepository {
					r := mocks.PassengerRepository{}
					r.On("Save", mock.Anything).Return(&user.Passenger{
						ID:    "123-abc",
						Name:  "Test",
						Email: email.Email{Address: "test@email.com"},
						Cpf:   cpf.Cpf{Number: "871.756.595-20"},
					}, nil)
					return &r
				}(),
			},
			args: args{
				input: Input{
					Name:  "Test",
					Email: "test@test.com",
					Cpf:   "871.756.595-20",
				},
			},
			expectedPassenger: &Output{
				ID:    "123-abc",
				Name:  "Test",
				Email: "test@email.com",
				Cpf:   "871.756.595-20",
			},
			expectedError: nil,
		},
		{
			name: "Given a invalid email When execute Then return error",
			fields: fields{
				passengerDatabase: func() repository.PassengerRepository {
					r := mocks.PassengerRepository{}
					r.On("Save", mock.Anything).Return(&user.Passenger{
						ID:    "123-abc",
						Name:  "Test",
						Email: email.Email{Address: "test@email.com"},
						Cpf:   cpf.Cpf{Number: "871.756.595-20"},
					}, nil)
					return &r
				}(),
			},
			args: args{
				input: Input{
					Name:  "Test",
					Email: "test.com",
					Cpf:   "871.756.595-20",
				},
			},
			expectedPassenger: nil,
			expectedError:     domain.NewUnprocessableEntityError("invalid_email", "Given e-mail is not valid.", ""),
		},
		{
			name: "Given a invalid cpf When execute Then return error",
			fields: fields{
				passengerDatabase: func() repository.PassengerRepository {
					r := mocks.PassengerRepository{}
					r.On("Save", mock.Anything).Return(&user.Passenger{
						ID:    "123-abc",
						Name:  "Test",
						Email: email.Email{Address: "test@email.com"},
						Cpf:   cpf.Cpf{Number: "871.756.595-20"},
					}, nil)
					return &r
				}(),
			},
			args: args{
				input: Input{
					Name:  "Test",
					Email: "test@test.com",
					Cpf:   "871.756.595-21",
				},
			},
			expectedPassenger: nil,
			expectedError:     domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", ""),
		},
		{
			name: "Given a valid input When execute and repository return error Then return error",
			fields: fields{
				passengerDatabase: func() repository.PassengerRepository {
					r := mocks.PassengerRepository{}
					r.On("Save", mock.Anything).Return(nil, domain.NewInternalServerError("error_save_driver", "Unexpected error while saving driver", ""))
					return &r
				}(),
			},
			args: args{
				input: Input{
					Name:  "Test",
					Email: "test@test.com",
					Cpf:   "871.756.595-20",
				},
			},
			expectedPassenger: nil,
			expectedError:     domain.NewInternalServerError("error_save_driver", "Unexpected error while saving driver", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewCreatePassengerUseCase(tt.fields.passengerDatabase)
			expectedPassenger, expectedError := useCase.Execute(tt.args.input)
			if !reflect.DeepEqual(expectedPassenger, tt.expectedPassenger) {
				t.Errorf("Execute() expectedPassenger = %v, want %v", expectedPassenger, tt.expectedPassenger)
			}
			if !reflect.DeepEqual(expectedError, tt.expectedError) {
				t.Errorf("Execute() got1 = %v, expectedError %v", expectedError, tt.expectedError)
			}
		})
	}
}
