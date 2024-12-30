package models

import (
	"github.com/google/uuid"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type SchedulesAppointmentRequest struct {
	Shift     int `json:"shift,omitempty"`
	Service   int `json:"service,omitempty"`
	DayOfWeek int `json:"day,omitempty"`
}

type OfficeScheduleResponse struct {
	ID           int
	OfficeID     int
	ShiftID      int
	ServiceID    int
	DoctorID     uuid.UUID
	StatusID     int
	DayOfWeek    int
	TimeStart    string `json:"time_start"`
	TimeEnd      string `json:"time_end"`
	TimeDuration string `json:"time_duration"`
}
