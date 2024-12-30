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

type AdminUser struct {
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
	AccountID        uuid.UUID
	MedicalLicense   string
	SpecialtyLicense string
	FirstName        string
	LastName1        string
	LastName2        string
	Sex              string
	Created_At       time.Time
	Updated_At       time.Time
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

type MedicalHistory struct {
	ID                       uuid.UUID  // Primary key
	MedicalHistoryID         string     // Unique identifier
	DateOfRecord             *time.Time // Fecha de registro
	TimeOfRecord             *time.Time // Hora de registro (formato HH:MM:SS)
	PatientName              string     // Nombre del paciente
	Curp                     string     // CURP
	BirthDate                *time.Time // Fecha de nacimiento
	Age                      string     // Edad
	Gender                   string     // Género
	PlaceOfOrigin            string     // Lugar de procedencia
	EthnicGroup              string     // Grupo étnico
	PhoneNumber              string     // Teléfono
	Address                  string     // Domicilio
	Occupation               string     // Ocupación
	GuardianName             string     // Nombre del tutor
	FamilyMedicalHistory     string     // Antecedentes médicos familiares
	NonPathologicalHistory   string     // Antecedentes no patológicos
	PathologicalHistory      string     // Antecedentes patológicos
	GynecObstetricHistory    string     // Antecedentes gineco-obstétricos
	CurrentCondition         string     // Condición actual
	Cardiovascular           string     // Sistema cardiovascular
	Respiratory              string     // Sistema respiratorio
	Gastrointestinal         string     // Sistema gastrointestinal
	Genitourinary            string     // Sistema genitourinario
	HematicLymphatic         string     // Sistema hemático y linfático
	Endocrine                string     // Sistema endocrino
	NervousSystem            string     // Sistema nervioso
	Musculoskeletal          string     // Sistema musculoesquelético
	Skin                     string     // Piel
	BodyTemperature          string     // Temperatura corporal
	Weight                   string     // Peso
	Height                   string     // Altura
	BMI                      string     // Índice de masa corporal (IMC)
	HeartRate                string     // Frecuencia cardíaca
	RespiratoryRate          string     // Frecuencia respiratoria
	BloodPressure            string     // Presión arterial
	Physical                 string     // Examen físico
	Head                     string     // Cabeza
	NeckAndChest             string     // Cuello y tórax
	Abdomen                  string     // Abdomen
	Genital                  string     // Genitales
	Extremities              string     // Extremidades
	PreviousResults          string     // Resultados anteriores
	Diagnoses                string     // Diagnósticos
	PharmacologicalTreatment string     // Tratamiento farmacológico
	Prognosis                string     // Pronóstico
	DoctorName               string     // Nombre del médico
	MedicalLicense           string     // Cédula médica
	SpecialtyLicense         string     // Cédula de especialidad
	Status                   bool       // Estado completo o incompleto
	Created_At               time.Time
	Updated_At               time.Time
}
