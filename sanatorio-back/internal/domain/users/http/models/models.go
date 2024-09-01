package models

import "github.com/google/uuid"

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

// Admin for obtions admin
type AdminData struct {
	AccountAdminID uuid.UUID `json:"account_admin_id"`
	RolAdmmin      int       `json:"rol_admin"`
	AdminPassword  string    `json:"admin_password"`
}

// Registers
type RegisterUserByAdminRequest struct {
	AdminData
	Name      string    `json:"name"`
	Lastname1 string    `json:"lastname1"`
	Lastname2 string    `json:"lastname2"`
	AccountID uuid.UUID `json:"account_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Rol       int       `json:"rol"`
	Curp      string    `json:"curp"`
}

type RegisterDoctorByAdminRequest struct {
	AdminData
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

type RegisterPatient struct {
	Name      string    `json:"name"`
	Lastname1 string    `json:"lastname1"`
	Lastname2 string    `json:"lastname2"`
	AccountID uuid.UUID `json:"account_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Curp      string    `json:"curp"`
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

// Objeto user
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Lastname1 string `json:"lastname1"`
	Lastname2 string `json:"lastname2"`
	Email     string
	Curp      string
}

// Get users
type Users struct {
	User
	Email string `json:"email"`
	Curp  string `json:"curp"`
}

type Doctors struct {
	User
	Email          string `json:"email"`
	MedicalLicense string `json:"medical_license"`
	Specialty      int    `json:"specialty"`
}

// Updates
type UpdateUser struct {
	AccountID uuid.UUID `json:"account_id"`
	Name      string    `json:"name,omitempty"`
	Lastname1 string    `json:"lastname1,omitempty"`
	Lastname2 string    `json:"lastname2,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Curp      string    `json:"curp,omitempty"`
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
}
