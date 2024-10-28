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
	Name        string `json:"name"`
	SpecialtyID int    `json:"specialtyID"`
}

type RegisterSpecialtyRequest struct {
	Name string `json:"name"`
}

type RegisterScheduleRequest struct {
	OfficeID  int    `json:"officeID"`
	DayOfWeek int    `json:"dayOfWeek"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
}

type RegisterAppointmentRequest struct {
	PatientAccountID uuid.UUID `json:"patientAccountID"`
	OfficeID         int       `json:"officeID"`
	Date             time.Time `json:"date"`
	ScheduleID       int       `json:"scheduleID"`
}
