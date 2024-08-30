package users

import (
	"time"

	"github.com/google/uuid"
)

type Role int

const (
	Admin = iota
	Doctor
	Patient
)

type User struct {
	ID        int
	AccountID uuid.UUID
	Name      string
	Lastname1 string
	Lastname2 string
}

type Account struct {
	ID       uuid.UUID
	UserID   int
	Email    string
	Password string
	Rol      int
}

type DoctorUser struct {
	User
	Account
	Especialidad      int
	CedulaProfesional string
}

type PatientUser struct {
	User
	Account
	FechaNacimiento time.Time
	Curp            string
}
