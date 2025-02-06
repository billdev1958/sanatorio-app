package entities

import (
	"time"

	"github.com/google/uuid"
)

type OfficeSchedule struct {
	ID           int
	OfficeID     int
	OfficeName   string
	ShiftID      int
	ServiceID    int
	DoctorID     string
	StatusID     int
	DayOfWeek    int
	TimeStart    time.Time
	TimeEnd      time.Time
	TimeDuration time.Duration
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

type Appointment struct {
	ID            uuid.UUID
	AccountID     uuid.UUID
	ScheduleID    int
	PatientID     uuid.UUID
	BeneficiaryID uuid.NullUUID
	TimeStart     time.Time
	TimeEnd       time.Time
	StatusID      int
	OfficeSchedule
	Consultation
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type AppointmentForPatient struct {
	AccountID     uuid.UUID
	PatientID     uuid.UUID
	BeneficiaryID uuid.NullUUID
	PatientName   string
	OfficeName    string
	ServiceName   string
	TimeStart     time.Time
	TimeEnd       time.Time
	StatusName    string
}

type AppointmentForReceptionist struct {
	AccountID     uuid.UUID
	PatientID     uuid.UUID
	BeneficiaryID uuid.NullUUID
	PatientName   string
	OfficeName    string
	ServiceName   string
	TimeStart     time.Time
	TimeEnd       time.Time
	StatusID      int
	StatusName    string
}

type AppointmentForDoctor struct {
	AccountID     uuid.UUID
	PatientID     uuid.UUID
	BeneficiaryID uuid.NullUUID
	TimeStart     time.Time
	TimeEnd       time.Time
	StatusID      string
	StatusName    string
	Consultation
}

type Consultation struct {
	ID            int
	AppointmentID uuid.UUID
	Reason        string
	Symptoms      string
	DoctorNotes   string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
}
