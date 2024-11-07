package models

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type RegisterOfficeRequest struct {
	Name      string `json:"name"`
	ServiceID int    `json:"service_id"`
}

type RegisterSpecialtyRequest struct {
	Name string `json:"name"`
}

type RegisterOfficeScheduleRequest struct {
	DayOfWeek    int       `json:"dayOfWeek"`
	TimeStart    string    `json:"timeStart"`
	TimeEnd      string    `json:"timeEnd"`
	TimeDuration string    `json:"timeDuration"`
	OfficeID     int       `json:"officeID"`
	ShiftID      int       `json:"shiftID"`
	ServiceID    int       `json:"serviceID"`
	DoctorID     uuid.UUID `json:"doctorID"`
}

type RegisterAppointmentRequest struct {
	PatientAccountID uuid.UUID `json:"patientAccountID"`
	OfficeID         int       `json:"officeID"`
	Date             time.Time `json:"date"`
	ScheduleID       int       `json:"scheduleID"`
}
