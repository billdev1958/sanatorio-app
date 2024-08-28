package users

import (
	"time"

	"github.com/google/uuid"
)

type Role int32

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

type Doctor struct {
	User
	Account
	CedulaProfesional string
}

type Patient struct {
	User
	Account
	FechaNacimiento time.Time
	Curp            string
}
