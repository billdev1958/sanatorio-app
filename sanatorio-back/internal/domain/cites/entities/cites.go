package cites

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID               uuid.UUID  `json:"id"`
	DoctorAccountID  uuid.UUID  `json:"doctor_account_id"`
	PatientAccountID uuid.UUID  `json:"patient_account_id"`
	OfficeID         int        `json:"office_id"`
	TimeStart        time.Time  `json:"time_start"`
	TimeEnd          time.Time  `json:"time_end"`
	ScheduleID       int        `json:"schedule_id"`
	StatusID         int        `json:"status_id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type Office struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	SpecialtyID     int        `json:"specialty_id"`
	StatusID        int        `json:"status_id"`
	DoctorAccountID uuid.UUID  `json:"doctor_account_id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

type OfficeStatus struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type AppointmentStatus struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Schedule struct {
	ID        int        `json:"id"`
	OfficeID  int        `json:"office_id"`
	DayOfWeek int        `json:"day_of_week"`
	TimeStart time.Time  `json:"time_start"`
	TimeEnd   time.Time  `json:"time_end"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Specialty struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
