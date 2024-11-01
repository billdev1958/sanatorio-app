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

type SuperAdminUser struct {
	AccountID  uuid.UUID
	FirstName  string
	LastName1  string
	LastName2  string
	Curp       string
	Sex        string
	Created_At time.Time
	Updated_At time.Time
}

type ReceptionistUser struct {
	AccountID  uuid.UUID
	FirstName  string
	LastName1  string
	LastName2  string
	Curp       string
	Sex        string
	Created_At time.Time
	Updated_At time.Time
}

type DoctorUser struct {
	AccountID      uuid.UUID
	MedicalLicense string
	SpecialtyID    int
	FirstName      string
	LastName1      string
	LastName2      string
	Sex            string
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
	Sex              string
	Created_At       time.Time
	Updated_At       time.Time
}

type BeneficiaryUser struct {
	ID               uuid.UUID
	AccountHolder    uuid.UUID
	MedicalHistoryID string
	Firstname        string
	Lastname1        string
	Lastname2        string
	Curp             string
	Sex              string
	Created_At       time.Time
	Updated_At       time.Time
}

type User struct {
	FirstName string
	LastName1 string
	LastName2 string
	Curp      string
	Sex       string
}
