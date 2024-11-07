package entities

import (
	"time"

	"github.com/google/uuid"
)

type OfficeStatusType int

const (
	_ OfficeStatusType = iota
	OfficeStatusAvailable
	OfficeStatusUnavailable
	OfficeStatusUnassigned
)

type AppointmentStatusType int

const (
	_ AppointmentStatusType = iota
	AppointmentStatusPendiente
	AppointmentStatusConfirmada
	AppointmentStatusCancelada
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
	ID        int
	Name      string
	ServiceID int
	StatusID  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
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
	ID           int
	DayOfWeek    int
	TimeStart    time.Time
	TimeEnd      time.Time
	TimeDuration time.Duration
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type OfficeSchedule struct {
	ID         int
	ScheduleID int
	OfficeID   int
	ShiftID    int
	ServiceID  int
	DoctorID   uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type Specialty struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
