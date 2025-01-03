package entities

import "time"

type OfficeSchedule struct {
	ID           int
	OfficeID     int
	OfficeName   string
	ShiftID      int
	ServiceID    int
	DoctorID     string
	StatusID     int
	DayOfWeek    int
	TimeStart    time.Time
	TimeEnd      time.Time
	TimeDuration time.Duration
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}
