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

type RegisterOfficeScheduleRequest struct {
	DayOfWeek int    `json:"dayOfWeek"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
	// TimeDuration string    `json:"timeDuration"`
	OfficeID  int       `json:"officeID"`
	ShiftID   int       `json:"shiftID"`
	ServiceID int       `json:"serviceID"`
	DoctorID  uuid.UUID `json:"doctorID"`
}

type RegisterAppointmentRequest struct {
	PatientAccountID uuid.UUID `json:"patientAccountID"`
	OfficeID         int       `json:"officeID"`
	Date             time.Time `json:"date"`
	ScheduleID       int       `json:"scheduleID"`
}

type OfficeScheduleResponse struct {
	OfficeScheduleID int       `json:"office_schedule_id"`
	ServiceID        int       `json:"service_id"`
	ScheduleID       int       `json:"schedule_id"`
	OfficeID         int       `json:"office_id"`
	OfficeStatus     int       `json:"office_status"`
	ShiftID          int       `json:"shift_id"`
	DoctorID         uuid.UUID `json:"doctor_id"`
	ServiceName      string    `json:"service_name"`
	DaySchedule      int       `json:"day_schedule"`
	TimeStart        string    `json:"time_start"`
	TimeEnd          string    `json:"time_end"`
	TimeDuration     string    `json:"time_duration"`
	OfficeName       string    `json:"office_name"`
	OfficeStatusName string    `json:"officeStatusName"`
	ShiftName        string    `json:"shift_name"`
	DoctorName       string    `json:"doctor_name"`
	DoctorLastName1  string    `json:"doctor_lastname1"`
	DoctorLastName2  string    `json:"doctor_lastname2"`
	MedicalLicense   string    `json:"medical_license"`
}
