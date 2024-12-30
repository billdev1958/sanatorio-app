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
	Shift     int `json:"shift"`
	Service   int `json:"service"`
	DayOfWeek int `json:"day"`
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
