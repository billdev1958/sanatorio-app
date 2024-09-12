package entities

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID               uuid.UUID
	PatientAccountID uuid.UUID
	OfficeID         int
	Date             time.Time
	ScheduleID       int
	StatusID         int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

type Office struct {
	ID              int
	Name            string
	SpecialtyID     int
	StatusID        int
	DoctorAccountID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

type OfficeStatus struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type AppointmentStatus struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Schedule struct {
	ID        int
	OfficeID  int
	DayOfWeek int
	TimeStart time.Duration
	TimeEnd   time.Duration
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Specialty struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
