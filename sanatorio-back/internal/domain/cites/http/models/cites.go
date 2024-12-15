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
	Name string `json:"name"`
}

type RegisterSpecialtyRequest struct {
	Name string `json:"name"`
}

type RegisterOfficeSchedule struct {
	SelectedDays []int     `json:"selectedDays"` // Array de días seleccionados (IDs de días)
	TimeStart    string    `json:"timeStart"`    // Hora de inicio en formato ISO 8601
	TimeEnd      string    `json:"timeEnd"`      // Hora de fin en formato ISO 8601
	TimeDuration string    `json:"timeDuration"` // Duración en formato hh:mm
	ShiftID      int       `json:"shiftID"`      // ID del turno (puede ser cadena porque viene como string en el JSON)
	ServiceID    int       `json:"serviceID"`    // ID del servicio (cadena)
	DoctorID     uuid.UUID `json:"doctorID"`     // UUID del doctor
	OfficeID     int       `json:"officeID"`     // ID de la oficina
	TimeSlots    []string  `json:"timeSlots"`    // Array de intervalos de tiempo generados
}
type RegisterAppointmentRequest struct {
	PatientAccountID uuid.UUID `json:"patientAccountID"`
	OfficeID         int       `json:"officeID"`
	Date             time.Time `json:"date"`
	ScheduleID       int       `json:"scheduleID"`
}
