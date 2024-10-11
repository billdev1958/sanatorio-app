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

type DoctorUser struct {
	ID uuid.UUID
	Account
	MedicalLicense string
	SpecialtyID    Specialties
	FirstName      string
	LastName1      string
	LastName2      string
	Sex            byte
	Created_At     time.Time
	Updated_At     time.Time
}

type PatientUser struct {
	ID               uuid.UUID
	MedicalHistoryID string
	LegacyID         string
	Account
	FirstName  string
	LastName1  string
	LastName2  string
	Curp       string
	Sex        byte
	Created_At time.Time
	Updated_At time.Time
}
