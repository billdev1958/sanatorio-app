package models

import (
	"time"

	"github.com/google/uuid"
)

type Office struct {
	ID   int    `json:"office_id"`
	Name string `json:"office_name"`
}

type DayOfWeek struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CatShift struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Services struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Doctor struct {
	AccountID uuid.UUID `json:"account_id"`
	FirstName string    `json:"first_name"`
	LastName1 string    `json:"last_name_1"`
	LastName2 string    `json:"last_name_2"`
}

type OfficeSchedule struct {
	ID           int
	OfficeID     int
	ShiftID      int
	ServiceID    int
	DoctorID     uuid.UUID
	StatusID     int
	DayOfWeek    int
	TimeStart    time.Time
	TimeEnd      time.Time
	TimeDuration time.Duration
}
