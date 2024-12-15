package cites

import (
	"context"
	"sanatorioApp/internal/domain/cites/entities"
)

type CitesRepository interface {
	Register
	Get
	Update
	//Delete
}

type Register interface {
	RegisterOffice(ctx context.Context, of entities.Office) (string, error)
	RegisterOfficeSchedule(ctx context.Context, schedules []entities.Schedule, officeSchedule entities.OfficeSchedule) (string, error)
	RegisterAppointment(ctx context.Context, ap entities.Appointment) (string, error)
}

type Get interface {
	//GetSpecialty(ctx context.Context) (entities.Services, error)
	GetOffices(ctx context.Context) ([]entities.Office, error)
	GetSchedules(ctx context.Context, filters map[string]interface{}) ([]entities.OfficeSchedule, error)
	//GetAppointment(ctx context.Context) (entities.Appointment, error)
}

type Update interface {
	//UpdateSpecialty(ctx context.Context, sp entities.Services) (entities.Services, error)
	UpdateOffice(ctx context.Context, of entities.Office) (string, error)
	//UpdateSchedule(ctx context.Context, sc entities.Schedule) (entities.Schedule, error)
	//UpdateAppointment(ctx context.Context, ap entities.Appointment) (entities.Appointment, error)
}

type Delete interface {
	DeleteSpecialty(ctx context.Context, id int) (entities.Services, error)
	DeleteOffice(ctx context.Context, id int) (entities.Office, error)
	DeleteSchedule(ctx context.Context, id int) (entities.Schedule, error)
	DeleteAppointment(ctx context.Context, id int) (entities.Appointment, error)
}
