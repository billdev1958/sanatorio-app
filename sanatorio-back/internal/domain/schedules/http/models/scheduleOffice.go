package models

import "sanatorioApp/internal/domain/catalogs/models"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type Office struct {
	OfficeID   int    `json:"office_id"`
	OfficeName string `json:"office_name"`
}

type DayOfWeek = models.DayOfWeek
type CatShift = models.CatShift
type Services = models.Services
type Doctor = models.Doctor

type GetInfoOfficeSchedule struct {
	CatDays     []DayOfWeek `json:"day_of_week"`
	CatShifts   []CatShift  `json:"cat_shift"`
	CatServices []Services  `json:"cat_services"`
	Doctors     []Doctor    `json:"doctor"`
}
