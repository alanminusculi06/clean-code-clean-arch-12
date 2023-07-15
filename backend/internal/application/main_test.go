package main

import (
	"backend/internal/application/dto"
	"backend/internal/pkg/domain"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_CalculateRidePrice(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name               string
		fields             fields
		expectedResult     *dto.CalculateRidePriceResultDTO
		expectedApiError   *domain.ApiError
		expectedStatusCode int
	}{
		{
			name: "Given a valid segment when calculate ride price then return price",
			fields: fields{
				input: `{
					"segments": [
						{ "distance": 10, "date": "2021-03-01T10:00:00" }
					]
				}`,
			},
			expectedResult:     &dto.CalculateRidePriceResultDTO{Price: 21.0},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Given an invalid segment distance when calculate ride price then return price",
			fields: fields{
				input: `{
					"segments": [
						{ "distance": -10, "date": "2021-03-01T10:00:00" }
					]
				}`,
			},
			expectedResult:     nil,
			expectedApiError:   domain.NewUnprocessableEntityError("error_negative_distance", "Distance cannot be negative", ""),
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := http.Post("http://localhost:8080/rides/calculate-price", "application/json", strings.NewReader(tt.fields.input))
			if err != nil {
				t.Errorf(err.Error())
			}

			defer func() {
				_ = response.Body.Close()
			}()

			if tt.expectedResult != nil {
				apiResult := &dto.CalculateRidePriceResultDTO{}
				body, _ := io.ReadAll(response.Body)
				_ = json.Unmarshal(body, apiResult)

				if !reflect.DeepEqual(apiResult.Price, tt.expectedResult.Price) {
					t.Errorf("Expected result %.2f, got %.2f", tt.expectedResult.Price, apiResult.Price)
				}
			} else {
				apiResult := &domain.ApiError{}
				body, _ := io.ReadAll(response.Body)
				_ = json.Unmarshal(body, apiResult)

				if !reflect.DeepEqual(apiResult, tt.expectedApiError) {
					t.Errorf("Expected api error %v, got %v", tt.expectedApiError, apiResult)
				}
			}

			if !reflect.DeepEqual(response.StatusCode, tt.expectedStatusCode) {
				t.Errorf("Expected status code %v, got %v", tt.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func Test_CreateDriver(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name               string
		fields             fields
		expectedApiError   *domain.ApiError
		expectedStatusCode int
	}{
		{
			name: "Given a valid input when create driver then return created id",
			fields: fields{
				input: `{
					"name": "Alan",
					"email": "alan@email.com.br",
					"cpf": "871.756.595-20",
					"car-plate": "MKJ-6775"
				}`,
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Given an invalid input when create driver then return api error",
			fields: fields{
				input: `{
					"name": "Alan",
					"email": "alan@email.com.br",
					"cpf": "834.326.160-76",
					"car-plate": "MKJ-6775"
				}`,
			},
			expectedApiError:   domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", ""),
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := http.Post("http://localhost:8080/drivers", "application/json", strings.NewReader(tt.fields.input))
			if err != nil {
				t.Errorf(err.Error())
			}

			defer func() {
				_ = response.Body.Close()
			}()

			if tt.expectedApiError == nil {
				apiResult := &dto.DriverDTO{}
				body, _ := io.ReadAll(response.Body)
				_ = json.Unmarshal(body, apiResult)

				if apiResult.ID == "" {
					t.Errorf("Expected result ID, got empty")
				}
			} else {
				apiResult := &domain.ApiError{}
				body, _ := io.ReadAll(response.Body)
				_ = json.Unmarshal(body, apiResult)

				if !reflect.DeepEqual(apiResult, tt.expectedApiError) {
					t.Errorf("Expected api error %v, got %v", tt.expectedApiError, apiResult)
				}
			}

			if !reflect.DeepEqual(response.StatusCode, tt.expectedStatusCode) {
				t.Errorf("Expected status code %v, got %v", tt.expectedStatusCode, response.StatusCode)
			}
		})
	}
}

func Test_CreatePassenger(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name               string
		fields             fields
		expectedApiError   *domain.ApiError
		expectedStatusCode int
	}{
		{
			name: "Given a valid input when create passenger then return created id",
			fields: fields{
				input: `{
					"name": "Alan",
					"email": "alan@email.com.br",
					"cpf": "871.756.595-20"
				}`,
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Given an invalid input when create driver then return api error",
			fields: fields{
				input: `{
					"name": "Alan",
					"email": "alan@email.com.br",
					"cpf": "834.326.160-76"
				}`,
			},
			expectedApiError:   domain.NewUnprocessableEntityError("invalid_cpf", "Given CPF is not valid.", ""),
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := http.Post("http://localhost:8080/passengers", "application/json", strings.NewReader(tt.fields.input))
			if err != nil {
				t.Errorf(err.Error())
			}

			defer func() {
				_ = response.Body.Close()
			}()

			if tt.expectedApiError == nil {
				apiResult := &dto.DriverDTO{}
				body, _ := io.ReadAll(response.Body)
				_ = json.Unmarshal(body, apiResult)

				if apiResult.ID == "" {
					t.Errorf("Expected result ID, got empty")
				}
			} else {
				apiResult := &domain.ApiError{}
				body, _ := io.ReadAll(response.Body)
				_ = json.Unmarshal(body, apiResult)

				if !reflect.DeepEqual(apiResult, tt.expectedApiError) {
					t.Errorf("Expected api error %v, got %v", tt.expectedApiError, apiResult)
				}
			}

			if !reflect.DeepEqual(response.StatusCode, tt.expectedStatusCode) {
				t.Errorf("Expected status code %v, got %v", tt.expectedStatusCode, response.StatusCode)
			}
		})
	}
}
