package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/cites/entities"

	"github.com/jackc/pgx/v5"
)

func (cr *citesRepository) RegisterOfficeSchedule(ctx context.Context, sc entities.Schedule, os entities.OfficeSchedule) (string, error) {
	// Inicia la transacción
	tx, err := cr.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Primero actualiza el estado de la oficina a "asignado"
	if err := cr.updateOfficeStatus(ctx, tx, os.OfficeID, 1); err != nil {
		return "", fmt.Errorf("failed to update office status: %w", err)
	}

	// Inserta el nuevo horario en la tabla `schedule` y obtén el `scheduleID`
	scheduleID, err := cr.insertSchedule(ctx, tx, sc)
	if err != nil {
		return "", fmt.Errorf("failed to insert schedule: %w", err)
	}

	// Inserta el nuevo registro en la tabla `office_schedule`, incluyendo el `scheduleID`
	os.ScheduleID = scheduleID
	if err := cr.insertOfficeSchedule(ctx, tx, os); err != nil {
		return "", fmt.Errorf("failed to insert office schedule: %w", err)
	}

	// Confirma la transacción
	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("Horario registrado con éxito para la oficina '%d' el día '%d'", os.OfficeID, sc.DayOfWeek), nil
}

func (cr *citesRepository) updateOfficeStatus(ctx context.Context, tx pgx.Tx, officeID int, statusID int) error {
	queryUpdate := "UPDATE office SET status_id = $1 WHERE id = $2"
	_, err := tx.Exec(ctx, queryUpdate, statusID, officeID)
	if err != nil {
		log.Printf("error al actualizar el estado de la oficina '%d' a '%d': %v", officeID, statusID, err)
		return err
	}
	return nil
}

func (cr *citesRepository) insertSchedule(ctx context.Context, tx pgx.Tx, sc entities.Schedule) (int, error) {
	queryInsert := `
		INSERT INTO schedule (day_of_week, time_start, time_end, time_duration)
		VALUES ($1, $2, $3, $4) RETURNING id
	`
	var scheduleID int
	err := tx.QueryRow(ctx, queryInsert, sc.DayOfWeek, sc.TimeStart, sc.TimeEnd, sc.TimeDuration).Scan(&scheduleID)
	if err != nil {
		log.Printf("error al registrar el horario para la oficina '%d' en el día '%d' con hora inicio '%s' y hora fin '%s': %v", sc.DayOfWeek, sc.TimeStart, sc.TimeEnd, sc.TimeDuration, err)
		return 0, err
	}
	return scheduleID, nil
}

func (cr *citesRepository) insertOfficeSchedule(ctx context.Context, tx pgx.Tx, sc entities.OfficeSchedule) error {
	queryInsert := `
		INSERT INTO office_schedule (schedule_id, office_id, shift_id, service_id, doctor_id)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (office_id, shift_id, service_id, doctor_id)
		DO NOTHING
	`
	_, err := tx.Exec(ctx, queryInsert, sc.ScheduleID, sc.OfficeID, sc.ShiftID, sc.ServiceID, sc.DoctorID)
	if err != nil {
		log.Printf("error al registrar la programación de la oficina '%d' con turno '%d' y servicio '%d': %v", sc.OfficeID, sc.ShiftID, sc.ServiceID, err)
		return err
	}
	return nil
}
