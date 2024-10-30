package models

import (
	"github.com/google/uuid"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type UserData struct {
	Name string `json:"name"`
}

// Registers

type RegisterPatientRequest struct {
	AfiliationID int    `json:"dependency_id"`
	Name         string `json:"name"`
	Lastname1    string `json:"lastname1"`
	Lastname2    string `json:"lastname2"`
	Curp         string `json:"curp"`
	Sex          string `json:"sex"`
	PhoneNumber  string `json:"phone"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type RegisterBeneficiaryRequest struct {
	Firstname string `json:"name"`
	Lastname1 string `json:"lastname1"`
	Lastname2 string `json:"lastname2"`
	Curp      string `json:"curp"`
	Sex       string `json:"sex"`
}

// Login
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccountID uuid.UUID `json:"account_id"`
	Role      int       `json:"role"`
	Token     string    `json:"token"`
}

// Get users
type UserRequest struct {
	AccountID  uuid.UUID `json:"account_id"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Lastname1  string    `json:"lastname1"`
	Lastname2  string    `json:"lastname2"`
	Email      string    `json:"email"`
	Curp       string    `json:"curp"`
	Created_At string    `json:"created_at"`
}

type DoctorRequest struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Lastname1      string    `json:"lastname1"`
	Lastname2      string    `json:"lastname2"`
	Email          string    `json:"email"`
	MedicalLicense string    `json:"medical_license"`
	SpecialtyID    int       `json:"specialty"`
	AccountID      uuid.UUID `json:"account_id"`
}

type PatientRequest struct {
	AccountID  string `json:"account_id"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Lastname1  string `json:"lastname1"`
	Lastname2  string `json:"lastname2"`
	Email      string `json:"email"`
	Curp       string `json:"curp"`
	Created_At string `json:"created_at"`
}

// Updates
type UpdateUser struct {
	AccountID     uuid.UUID `json:"account_id"`
	Name          string    `json:"name,omitempty"`
	Lastname1     string    `json:"lastname1,omitempty"`
	Lastname2     string    `json:"lastname2,omitempty"`
	Email         string    `json:"email,omitempty"`
	Password      string    `json:"password,omitempty"`
	Curp          string    `json:"curp,omitempty"`
	AdminPassword string    `json:"admin_password"`
}

type UpdateDoctor struct {
	AccountID      uuid.UUID `json:"account_id"`
	Name           string    `json:"name,omitempty"`
	Lastname1      string    `json:"lastname1,omitempty"`
	Lastname2      string    `json:"lastname2,omitempty"`
	Email          string    `json:"email,omitempty"`
	Password       string    `json:"password,omitempty"`
	SpecialtyID    int       `json:"specialty_id,omitempty"`
	MedicalLicense string    `json:"medical_license,omitempty"`
	AdminPassword  string    `json:"admin_password"`
}
