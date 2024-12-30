package repository

import (
	"sanatorioApp/internal/domain/appointment"
	postgres "sanatorioApp/internal/infraestructure/db"
)

type appointmentRepository struct {
	storage *postgres.PgxStorage
}

func NewAppointmentRepository(storage *postgres.PgxStorage) appointment.AppointmentRepository {
	return &appointmentRepository{storage: storage}
}
