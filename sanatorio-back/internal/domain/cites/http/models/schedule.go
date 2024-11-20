package models

import "github.com/google/uuid"

type OfficeScheduleResponse struct {
	OfficeScheduleID int       `json:"office_schedule_id"`
	ServiceID        int       `json:"service_id"`
	ScheduleID       int       `json:"schedule_id"`
	OfficeID         int       `json:"office_id"`
	OfficeStatusID   int       `json:"office_status"`
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

// Filters for servicesID, OfficeID, StatusID, ShiftID, DoctorID, DayOfWeek,
type OfficeSCheduleFiltersRequest struct {
	ServiceID      *int       `json:"serviceID,omitempty"`
	DoctorID       *uuid.UUID `json:"doctorID,omitempty"`
	OfficeID       *int       `json:"officeID,omitempty"`
	OfficeStatusID *int       `json:"officeStatusID,omitempty"`
	ShiftID        *int       `json:"shiftID,omitempty"`
	DayOfWeek      *int       `json:"dayOfWeek,omitempty"`
}
