package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/cites"
	"sanatorioApp/internal/domain/cites/entities"
	postgres "sanatorioApp/internal/infraestructure/db"
)

type citesRepository struct {
	storage *postgres.PgxStorage
}

func NewCitesRepository(storage *postgres.PgxStorage) cites.CitesRepository {
	return &citesRepository{storage: storage}
}

func (cr citesRepository) RegisterSpecialty(ctx context.Context, sp entities.Specialty) (string, error) {
	query := "INSERT INTO cat_specialty (name) VALUES ($1)"
	_, err := cr.storage.DbPool.Exec(ctx, query, sp.Name)
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
	tx, err := cr.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Primero actualiza el status de la oficina a "asignado" (por ejemplo, status_id = 1)
	queryUpdate := "UPDATE office SET status_id = 1 WHERE id = $1" // Cambia 1 al ID de "asignado" si es diferente
	_, err = tx.Exec(ctx, queryUpdate, sc.OfficeID)
	if err != nil {
		log.Printf("error al actualizar el estado de la oficina '%d': %v", sc.OfficeID, err)
		return "", fmt.Errorf("failed to update office status: %w", err)
	}

	// Inserta el nuevo horario
	queryInsert := "INSERT INTO schedule (office_id, day_of_week, time_start, time_end) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(ctx, queryInsert, sc.OfficeID, sc.DayOfWeek, sc.TimeStart, sc.TimeEnd)
	if err != nil {
		log.Printf("error al registrar el horario para la oficina '%d' en el día '%d' con hora inicio '%s' y hora fin '%s': %v", sc.OfficeID, sc.DayOfWeek, sc.TimeStart, sc.TimeEnd, err)
		return "", fmt.Errorf("failed to insert schedule: %w", err)
	}

	// Confirma la transacción
	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("Horario registrado con éxito para la oficina '%d' el día '%d'", sc.OfficeID, sc.DayOfWeek), nil
}

func (cr citesRepository) RegisterAppointment(ctx context.Context, ap entities.Appointment) (string, error) {
	// Inserta la nueva cita en la tabla appointment
	query := `
		INSERT INTO appointment (id, patient_account_id, office_id, date, schedule_id, status_id)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := cr.storage.DbPool.Exec(ctx, query, ap.ID, ap.PatientAccountID, ap.OfficeID, ap.Date, ap.ScheduleID, ap.StatusID)
	if err != nil {
		log.Printf("error al registrar la cita para el paciente '%s' en la oficina '%d': %v", ap.PatientAccountID.String(), ap.OfficeID, err)
		return "", fmt.Errorf("failed to insert appointment: %w", err)
	}

	return fmt.Sprintf("Cita registrada con éxito para el paciente '%s'", ap.PatientAccountID.String()), nil
}
