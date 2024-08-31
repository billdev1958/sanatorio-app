package models

import "github.com/google/uuid"

type RegisterUserRequest struct {
	Name      string    `json:"name"`
	Lastname1 string    `json:"lastname1"`
	Lastname2 string    `json:"lastname2"`
	AccountID uuid.UUID `json:"account_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Rol       int       `json:"rol"`
	Curp      string    `json:"curp"`
}

type RegisterDoctorRequest struct {
	Name           string    `json:"name"`
	Lastname1      string    `json:"lastname1"`
	Lastname2      string    `json:"lastname2"`
	AccountID      uuid.UUID `json:"account_id"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Rol            int       `json:"rol"`
	MedicalLicense string    `json:"medical_license"`
	Specialty      int       `json:"specialty"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
