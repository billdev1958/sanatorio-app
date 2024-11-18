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

func (cr *citesRepository) RegisterSpecialty(ctx context.Context, sp entities.Services) (string, error) {
	query := "INSERT INTO cat_specialty (name) VALUES ($1)"
	_, err := cr.storage.DbPool.Exec(ctx, query, sp.Name)
	if err != nil {
		log.Printf("error al registrar la especialidad '%s' en la db: %v", sp.Name, err)
		return "", err
	}
	return fmt.Sprintf("Especialidad '%s' registrada con éxito", sp.Name), nil
}

func (cr *citesRepository) RegisterOffice(ctx context.Context, of entities.Office) (string, error) {
	query := `
		INSERT INTO office (name, status_id)
		VALUES ($1, $2, $3)`

	// Ejecutar la consulta
	_, err := cr.storage.DbPool.Exec(ctx, query, of.Name, entities.OfficeStatusUnassigned)
	if err != nil {
		log.Printf("error al registrar el consultorio '%s' en la db: %v", of.Name, err)
		return "", err
	}

	return fmt.Sprintf("Consultorio '%s' registrado con éxito", of.Name), nil
}

func (cr *citesRepository) RegisterAppointment(ctx context.Context, ap entities.Appointment) (string, error) {
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
