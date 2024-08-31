package entities

import "github.com/google/uuid"

type Rol int

const (
	_ = iota
	SuperUsuario
	Doctor
	Patient
)

type Specialties int

const (
	_ = iota
	MedicinaGeneral
	Cardiologia
	Psiquiatria
)

type User struct {
	ID        int
	Name      string
	Lastname1 string
	Lastname2 string
}

type Account struct {
	AccountID uuid.UUID
	UserID    int
	Email     string
	Password  string
	Rol       int
}

type SuperUser struct {
	ID int
	User
	Account
	Curp string
}

type DoctorUser struct {
	ID int
	User
	Account
	MedicalLicense string
}

type PatientUser struct {
	ID int
	User
	Account
	Curp string
}

type UserResponse struct {
	Name  string
	Email string
}

type RegisterUser struct {
	User
	Account
	DocumentID string
}

type RegisterDoctor struct {
	User
	Account
	SpecialtyID int
	DocumentID  string
}
