package cites

import (
	"context"
	"sanatorioApp/internal/domain/cites/entities"

	"github.com/google/uuid"
)

type CitesRepository interface {
	Register
	//Get
	//Update
	//Delete
}

type Register interface {
	RegisterSpecialty(ctx context.Context, sp entities.Specialty) error
	RegisterOffice(ctx context.Context, of entities.Office) error
	RegisterSchedule(ctx context.Context, sc entities.Schedule) error
	RegisterAppointment(ctx context.Context, ap entities.Appointment) (entities.Appointment, error)
}

type Get interface {
	GetSpecialty(ctx context.Context) (entities.Specialty, error)
	GetOffice(ctx context.Context) (entities.Office, error)
	GetSchedule(ctx context.Context) (entities.Schedule, error)
	GetAppointment(ctx context.Context) (entities.Appointment, error)
}

type Update interface {
	UpdateSpecialty(ctx context.Context, sp entities.Specialty) (entities.Specialty, error)
	UpdateOffice(ctx context.Context, of entities.Office) (entities.Office, error)
	UpdateSchedule(ctx context.Context, sc entities.Schedule) (entities.Schedule, error)
	UpdateAppointment(ctx context.Context, ap entities.Appointment) (entities.Appointment, error)
	AssignDoctorToOffice(ctx context.Context, officeID int, doctorAccountID uuid.UUID) error
}

type Delete interface {
	DeleteSpecialty(ctx context.Context, id int) (entities.Specialty, error)
	DeleteOffice(ctx context.Context, id int) (entities.Office, error)
	DeleteSchedule(ctx context.Context, id int) (entities.Schedule, error)
	DeleteAppointment(ctx context.Context, id int) (entities.Appointment, error)
}
