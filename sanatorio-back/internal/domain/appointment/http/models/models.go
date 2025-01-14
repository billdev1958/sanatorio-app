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

type SchedulesAppointmentRequest struct {
	Shift           int    `json:"shift,omitempty"`
	Service         int    `json:"service,omitempty"`
	AppointmentDate string `json:"appointmentDate,omitempty"`
}

type OfficeScheduleResponse struct {
	ID           int    `json:"id"`
	TimeStart    string `json:"timeStart"`
	TimeEnd      string `json:"timeEnd"`
	TimeDuration string `json:"timeDuration"`
	OfficeName   string `json:"officeName"`
	StatusID     int    `json:"statusID"`
}

type RegisterAppointmentRequest struct {
	ScheduleID    int       `json:"scheduleID"`
	PatientID     uuid.UUID `json:"patientID"`
	BeneficiaryID uuid.UUID `json:"beneficiaryID,omitempty"`
	TimeStart     time.Time `json:"timeStart"`
	TimeEnd       time.Time `json:"timeEnd"`
	Reason        string    `json:"reason,omitempty"`
	Symptoms      string    `json:"symptoms,omitempty"`
}
