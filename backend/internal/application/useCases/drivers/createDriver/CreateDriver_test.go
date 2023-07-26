package createDriver

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

func TestCreateDriver_Execute(t *testing.T) {
	type fields struct {
		driverDatabase repository.DriverRepository
	}
	type args struct {
		input Input
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedDriver *Output
		expectedError  *domain.ApiError
	}{
		{
			name: "Given a valid input When execute Then return nil",
			fields: fields{
				driverDatabase: func() repository.DriverRepository {
					r := mocks.DriverRepository{}
					r.On("Save", mock.Anything).Return(&user.Driver{
						ID:       "123-abc",
						Name:     "Test",
						Email:    email.Email{Address: "test@email.com"},
						CarPlate: "ABC1234",
						Cpf:      cpf.Cpf{Number: "871.756.595-20"},
					}, nil)
					return &r
				}(),
			},
			args: args{
				input: Input{
					Name:     "Test",
					Email:    "test@test.com",
					Cpf:      "871.756.595-20",
					CarPlate: "ABC1234",
				},
			},
			expectedDriver: &Output{
				ID:       "123-abc",
				Name:     "Test",
				Email:    "test@email.com",
				CarPlate: "ABC1234",
				Cpf:      "871.756.595-20",
			},
			expectedError: nil,
		},
		{
			name: "Given a invalid email When execute Then return error",
			fields: fields{
				driverDatabase: func() repository.DriverRepository {
					r := mocks.DriverRepository{}
					r.On("Save", mock.Anything).Return(&user.Driver{
						ID:       "123-abc",
						Name:     "Test",
						Email:    email.Email{Address: "test@email.com"},
						CarPlate: "ABC1234",
						Cpf:      cpf.Cpf{Number: "871.756.595-20"},
					}, nil)
					return &r
				}(),
			},
			args: args{
				input: Input{
					Name:     "Test",
					Email:    "test.com",
					Cpf:      "871.756.595-20",
					CarPlate: "ABC1234",
				},
			},
			expectedDriver: nil,
			expectedError:  domain.NewUnprocessableEntityError("invalid_email", "Given e-mail is not valid.", ""),
		},
		{
			name: "Given a invalid cpf When execute Then return error",
			fields: fields{
				driverDatabase: func() repository.DriverRepository {
					r := mocks.DriverRepository{}
					r.On("Save", mock.Anything).Return(&user.Driver{
						ID:       "123-abc",
						Name:     "Test",
						Email:    email.Email{Address: "test@email.com"},
						CarPlate: "ABC1234",
						Cpf:      cpf.Cpf{Number: "871.756.595-20"},
					}, nil)
					return &r
				}(),
			},
			args: args{
				input: Input{
					Name:     "Test",
					Email:    "test@test.com",
					Cpf:      "871.756.595-21",
					CarPlate: "ABC1234",
				},
			},
			expectedDriver: nil,
			expectedError:  domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", ""),
		},
		{
			name: "Given a valid input When execute and repository return error Then return error",
			fields: fields{
				driverDatabase: func() repository.DriverRepository {
					r := mocks.DriverRepository{}
					r.On("Save", mock.Anything).Return(nil, domain.NewInternalServerError("error_save_driver", "Unexpected error while saving driver", ""))
					return &r
				}(),
			},
			args: args{
				input: Input{
					Name:     "Test",
					Email:    "test@test.com",
					Cpf:      "871.756.595-20",
					CarPlate: "ABC1234",
				},
			},
			expectedDriver: nil,
			expectedError:  domain.NewInternalServerError("error_save_driver", "Unexpected error while saving driver", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := CreateDriver{
				driverDatabase: tt.fields.driverDatabase,
			}
			driver, err := useCase.Execute(tt.args.input)
			if driver != nil && !reflect.DeepEqual(driver, tt.expectedDriver) {
				t.Errorf("Execute() driver = %v, want %v", err, tt.expectedDriver)
			}
			if !reflect.DeepEqual(err, tt.expectedError) {
				t.Errorf("Execute() err = %v, want %v", err, tt.expectedError)
			}
		})
	}
}
