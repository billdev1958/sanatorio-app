package repository

import (
	"context"
	"log"
	"sanatorioApp/internal/domain/cites"
	"sanatorioApp/internal/domain/cites/entities"
	postgres "sanatorioApp/internal/infraestructure/db"
	"time"

	"github.com/google/uuid"
)

type citesRepository struct {
	storage *postgres.PgxStorage
}

func NewCitesRepository(storage *postgres.PgxStorage) cites.CitesRepository {
	return &citesRepository{storage: storage}
}

func (cr citesRepository) RegisterSpecialty(ctx context.Context, sp entities.Specialty) error {
	query := "INSERT INTO cat_specialty (name, created_at) VALUES ($1, $2)"
	_, err := cr.storage.DbPool.Exec(ctx, query, sp.Name, time.Now())
	if err != nil {
		log.Printf("error al registrar la especialidad '%s' en la db: %v", sp.Name, err)
		return err
	}
	return nil
}

func (cr citesRepository) RegisterOffice(ctx context.Context, of entities.Office) error {
	query := `
		INSERT INTO office (name, specialty_id, status_id, created_at)
		VALUES ($1, $2, $3, $4)`

	// Ejecutar la consulta
	_, err := cr.storage.DbPool.Exec(ctx, query, of.Name, of.SpecialtyID, of.StatusID, time.Now())
	if err != nil {
		log.Printf("error al registrar el consultorio '%s' en la db: %v", of.Name, err)
		return err
	}

	return nil
}

func (cr citesRepository) AssignDoctorToOffice(ctx context.Context, officeID int, doctorAccountID uuid.UUID) error {
	query := "UPDATE office SET doctor_account_id = $1 WHERE id = $2"
	_, err := cr.storage.DbPool.Exec(ctx, query, doctorAccountID, officeID)
	if err != nil {
		log.Printf("error al asignar el doctor '%s' al consultorio '%d': %v", doctorAccountID, officeID, err)
		return err
	}

	return nil
}

func (cr citesRepository) RegisterSchedule(ctx context.Context, sc entities.Schedule) error {
	// Proceder con el registro del horario, el trigger se encargará de validar el estado de la oficina
	query := "INSERT INTO schedule (office_id, day_of_week, time_start, time_end, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := cr.storage.DbPool.Exec(ctx, query, sc.OfficeID, sc.DayOfWeek, sc.TimeStart, sc.TimeEnd, time.Now())
	if err != nil {
		log.Printf("error al registrar el horario para la oficina '%d' en el día '%d' con hora inicio '%s' y hora fin '%s': %v", sc.OfficeID, sc.DayOfWeek, sc.TimeStart, sc.TimeEnd, err)
		return err
	}

	return nil
}

func (cr citesRepository) RegisterAppointment(ctx context.Context, ap entities.Appointment) (entities.Appointment, error)
