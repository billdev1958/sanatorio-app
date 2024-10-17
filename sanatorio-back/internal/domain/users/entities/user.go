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
	ID                  uuid.UUID
	AfiliationID        int
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
	AccountID      uuid.UUID
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
	AccountID        uuid.UUID
	MedicalHistoryID string
	LegacyID         string
	FirstName        string
	LastName1        string
	LastName2        string
	Curp             string
	Sex              byte
	Created_At       time.Time
	Updated_At       time.Time
}
