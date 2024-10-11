package entities

import (
	"time"

	"github.com/google/uuid"
)

type Specialties int

const (
	_ = iota
	MedicinaGeneral
	Cardiologia
	Psiquiatria
)

type Account struct {
	AccountID           uuid.UUID
	AfiliationID        int
	UserID              int
	PhoneNumber         string
	Email               string
	Password            string
	Rol                 Roles
	Created_At          time.Time
	Updated_At          time.Time
	Password_Changed_At time.Time
	Deleted_At          time.Time
}

// UsersTypes
type SuperAdminUser struct {
	Account
	FirstName  string
	LastName1  string
	Lastname2  string
	Curp       string
	Sex        byte
	Created_At time.Time
	Updated_At time.Time
}

type AdminUser struct {
	Account
	FirstName  string
	LastName1  string
	Lastname2  string
	Curp       string
	Sex        byte
	Created_At time.Time
	Updated_At time.Time
}

type DoctorUser struct {
	Account
	MedicalLicense string
	SpecialtyID    Specialties
	FirstName      string
	LastName1      string
	Lastname2      string
	Sex            byte
	Created_At     time.Time
	Updated_At     time.Time
}

type PatientUser struct {
	ID               uuid.UUID
	MedicalHistoryID string
	Account
	FirstName  string
	LastName1  string
	Lastname2  string
	Curp       string
	Sex        byte
	Created_At time.Time
	Updated_At time.Time
}
