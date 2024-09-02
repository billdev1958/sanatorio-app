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

type Users struct {
	User
	Email string
	Rol   int
	Curp  string
}

type Doctors struct {
	User
	Email          string
	Rol            int
	MedicalLicense string
	Specialty      int
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
	AccountID uuid.UUID
	Name      string
	Lastname1 string
	Lastname2 string
	Email     string
	Password  string
	Curp      string
}

type UpdateDoctor struct {
	AccountID      uuid.UUID
	Name           string
	Lastname1      string
	Lastname2      string
	Email          string
	Password       string
	SpecialtyID    int
	MedicalLicense string
}
