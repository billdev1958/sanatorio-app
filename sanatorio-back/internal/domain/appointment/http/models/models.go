package models

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
	ID           int
	TimeStart    string `json:"time_start"`
	TimeEnd      string `json:"time_end"`
	TimeDuration string `json:"time_duration"`
	OfficeName   string
	StatusID     int
}
