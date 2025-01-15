package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/appointment"
	"sanatorioApp/internal/domain/appointment/entities"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type appointmentRepository struct {
	storage *postgres.PgxStorage
}

func NewAppointmentRepository(storage *postgres.PgxStorage) appointment.AppointmentRepository {
	return &appointmentRepository{storage: storage}
}

func (ar *appointmentRepository) GetAvaliableSchedules(ctx context.Context, date string, dayOfWeek int, serviceID int, shiftID int) ([]entities.OfficeSchedule, error) {
	query := `
        SELECT
            os.id AS schedule_id,
            os.time_start,
            os.time_end,
            os.time_duration,
            o.name AS office_name,
            CASE
                WHEN a.id IS NOT NULL THEN 2  -- Ocupado (cita)
                WHEN sb.id IS NOT NULL THEN 3  -- Bloqueado
                ELSE 1  -- Disponible
            END AS status_id
        FROM
            office_schedule os
        JOIN office o ON os.office_id = o.id
        LEFT JOIN appointment a ON
            os.id = a.schedule_id AND DATE(a.time_start) = $1
        LEFT JOIN schedule_block sb ON
            os.id = sb.office_schedule_id AND sb.block_date = $1
            AND (
                (sb.time_start IS NULL AND sb.time_end IS NULL) OR -- Bloqueo de día completo
                (os.time_start >= sb.time_start AND os.time_end <= sb.time_end) OR
                (os.time_start < sb.time_end AND os.time_end > sb.time_start)
            )
        WHERE
            os.service_id = $2
            AND os.shift_id = $3
            AND os.day_of_week = $4
            AND os.status_id = 1
        ORDER BY os.time_start ASC;
    `

	// Log de los parámetros de entrada
	log.Printf("GetAvaliableSchedules - Ejecutando consulta con parámetros: date=%s, serviceID=%d, shiftID=%d, dayOfWeek=%d", date, serviceID, shiftID, dayOfWeek)

	rows, err := ar.storage.DbPool.Query(ctx, query, date, serviceID, shiftID, dayOfWeek)
	if err != nil {
		log.Printf("GetAvaliableSchedules - Error ejecutando la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	var schedules []entities.OfficeSchedule

	for rows.Next() {
		var schedule entities.OfficeSchedule
		err := rows.Scan(
			&schedule.ID,
			&schedule.TimeStart,
			&schedule.TimeEnd,
			&schedule.TimeDuration,
			&schedule.OfficeName,
			&schedule.StatusID,
		)
		if err != nil {
			log.Printf("GetAvaliableSchedules - Error escaneando fila: %v", err)
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	// Verificar errores del iterador
	if rows.Err() != nil {
		log.Printf("GetAvaliableSchedules - Error durante la iteración de filas: %v", rows.Err())
		return nil, rows.Err()
	}

	// Log de los resultados obtenidos
	log.Printf("GetAvaliableSchedules - Total de horarios obtenidos: %d", len(schedules))

	return schedules, nil
}

func (ar *appointmentRepository) RegisterAppointment(ctx context.Context, a entities.Appointment, c entities.Consultation) (bool, error) {
	tx, err := ar.storage.DbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return false, fmt.Errorf("error al iniciar la transacción: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			if commitErr := tx.Commit(ctx); commitErr != nil {
				err = fmt.Errorf("error al confirmar la transacción: %v", commitErr)
			}
		}
	}()

	queryRegisterAppointment := `
		INSERT INTO appointment (
			id, account_id, schedule_id, patient_id, beneficiary_id, 
			time_start, time_end, status_id
		) 
		VALUES (
			$1, $2, $3, $4, COALESCE($5, NULL), 
			$6, $7, $8
		) RETURNING id
	`

	var appointmentID uuid.UUID
	err = tx.QueryRow(
		ctx,
		queryRegisterAppointment,
		a.ID,
		a.AccountID,
		a.ScheduleID,
		a.PatientID,
		a.BeneficiaryID,
		a.TimeStart,
		a.TimeEnd,
		a.StatusID,
	).Scan(&appointmentID)
	if err != nil {
		return false, fmt.Errorf("error al registrar el appointment: %w", err)
	}

	queryRegisterConsultation := `
		INSERT INTO consultation (appointment_id, reason, symptoms)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(ctx, queryRegisterConsultation, appointmentID, c.Reason, c.Symptoms)
	if err != nil {
		return false, fmt.Errorf("error al registrar la consulta médica: %w", err)
	}

	return true, nil
}
