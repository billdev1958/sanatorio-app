package models

import "sanatorioApp/internal/domain/catalogs/models"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type DayOfWeek = models.DayOfWeek
type CatShift = models.CatShift
type Services = models.Services
type Doctor = models.Doctor
type Office = models.Office

type GetInfoOfficeSchedule struct {
	CatDays     []DayOfWeek `json:"day_of_week"`
	CatShifts   []CatShift  `json:"cat_shift"`
	CatServices []Services  `json:"cat_services"`
	Doctors     []Doctor    `json:"doctor"`
	Offices     []Office    `json:"office"`
}
