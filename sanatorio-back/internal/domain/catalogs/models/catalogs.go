package models

import "github.com/google/uuid"

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