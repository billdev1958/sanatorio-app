package entities

import (
	"sanatorioApp/internal/domain/users/entities"
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

type CatStatus struct {
	ID   int
	Name string
}

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
	ID   int
	Name string
	OfficeStatus
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
	ID int
	Services
	Schedule
	Office
	StatusName string
	ShiftID    int
	ShiftName  string
	entities.DoctorUser
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Services struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
