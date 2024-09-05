package entities

import (
	"time"

	"github.com/google/uuid"
)

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

type Users struct {
	User
	Email      string
	Rol        int
	Curp       string
	Created_At time.Time
	AccountID  uuid.UUID
}

type Doctors struct {
	User
	Email          string
	Rol            int
	MedicalLicense string
	Specialty      int
	AccountID      uuid.UUID
}

type SuperUser struct {
	User
	Email      string
	Rol        int
	Curp       string
	Created_At time.Time
	AccountID  uuid.UUID
}

type DoctorUser struct {
	ID int
	User
	Account
	MedicalLicense string
}

type Account struct {
	AccountID uuid.UUID
	UserID    int
	Email     string
	Password  string
	Rol       int
}

type PatientUser struct {
	Name      string
	Lastname1 string
	Lastname2 string
	AccountID uuid.UUID
	Email     string
	Password  string
	Rol       int
	Curp      string
}

type UserResponse struct {
	Name  string
	Email string
}

type AdminData struct {
	AccountAdminID uuid.UUID
	RolAdmmin      int
	AdminPassword  string
}

type RegisterUserByAdmin struct {
	AdminData
	User
	Account
	DocumentID string
}

type RegisterDoctorByAdmin struct {
	AdminData
	User
	Account
	SpecialtyID int
	DocumentID  string
}

type LoginUser struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccountID uuid.UUID
	Role      int
}

type UpdateUser struct {
	AdminData
	AccountID uuid.UUID
	Name      string
	Lastname1 string
	Lastname2 string
	Email     string
	Password  string
	Curp      string
}

type UpdateDoctor struct {
	AdminData
	AccountID      uuid.UUID
	Name           string
	Lastname1      string
	Lastname2      string
	Email          string
	Password       string
	SpecialtyID    int
	MedicalLicense string
}
