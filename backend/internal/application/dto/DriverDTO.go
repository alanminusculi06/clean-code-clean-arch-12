package dto

type DriverDTO struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
	CarPlate string `json:"car-plate"`
}
