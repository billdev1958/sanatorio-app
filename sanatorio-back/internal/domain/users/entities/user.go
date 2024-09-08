package entities

import (
	"time"

	"github.com/google/uuid"
)

type Roles int

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

type AdminData struct {
	AccountID     uuid.UUID
	RoleAdmin     int
	PasswordAdmin string
}

type Account struct {
	AccountID           uuid.UUID
	UserID              int
	Email               string
	Password            string
	Rol                 Roles
	Created_At          time.Time
	Updated_At          time.Time
	Password_Changed_At time.Time
}

type User struct {
	ID         int
	Name       string
	Lastname1  string
	Lastname2  string
	Created_At time.Time
	Updated_At time.Time
}

type SuperUser struct {
	User
	Account
	Curp       string
	Created_At time.Time
	Updated_At time.Time
}

type DoctorUser struct {
	User
	Account
	MedicalLicense string
	SpecialtyID    Specialties
	Created_At     time.Time
	Updated_At     time.Time
}

type PatientUser struct {
	User
	Account
	Curp       string
	Created_At time.Time
	Updated_At time.Time
}
