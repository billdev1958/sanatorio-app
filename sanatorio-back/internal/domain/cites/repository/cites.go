package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/cites"
	"sanatorioApp/internal/domain/cites/entities"
	postgres "sanatorioApp/internal/infraestructure/db"
	"time"
)

type citesRepository struct {
	storage *postgres.PgxStorage
}

func NewCitesRepository(storage *postgres.PgxStorage) cites.CitesRepository {
	return &citesRepository{storage: storage}
}

func (cr citesRepository) RegisterSpecialty(ctx context.Context, sp entities.Specialty) (string, error) {
	query := "INSERT INTO cat_specialty (name) VALUES ($1, $2)"
	_, err := cr.storage.DbPool.Exec(ctx, query, sp.Name, time.Now())
	if err != nil {
		log.Printf("error al registrar la especialidad '%s' en la db: %v", sp.Name, err)
		return "", err
	}
	return fmt.Sprintf("Especialidad '%s' registrada con éxito", sp.Name), nil
}

func (cr citesRepository) RegisterOffice(ctx context.Context, of entities.Office) (string, error) {
	query := `
		INSERT INTO office (name, specialty_id, status_id)
		VALUES ($1, $2, $3)`

	// Ejecutar la consulta
	_, err := cr.storage.DbPool.Exec(ctx, query, of.Name, of.SpecialtyID, entities.OfficeStatusUnassigned)
	if err != nil {
		log.Printf("error al registrar el consultorio '%s' en la db: %v", of.Name, err)
		return "", err
	}

	return fmt.Sprintf("Consultorio '%s' registrado con éxito", of.Name), nil
}

func (cr citesRepository) RegisterSchedule(ctx context.Context, sc entities.Schedule) (string, error) {
	query := "INSERT INTO schedule (office_id, day_of_week, time_start, time_end) VALUES ($1, $2, $3, $4)"
	_, err := cr.storage.DbPool.Exec(ctx, query, sc.OfficeID, sc.DayOfWeek, sc.TimeStart, sc.TimeEnd)
	if err != nil {
		log.Printf("error al registrar el horario para la oficina '%d' en el día '%d' con hora inicio '%s' y hora fin '%s': %v", sc.OfficeID, sc.DayOfWeek, sc.TimeStart, sc.TimeEnd, err)
		return "", err
	}

	return fmt.Sprintf("Horario registrado con éxito para la oficina '%d' el día '%d'", sc.OfficeID, sc.DayOfWeek), nil
}

func (cr citesRepository) RegisterAppointment(ctx context.Context, ap entities.Appointment) (string, error)
