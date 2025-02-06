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
        WITH AvailableSchedules AS (
            SELECT
                os.id AS schedule_id,
                os.time_start,
                os.time_end,
                os.time_duration,
                o.name AS office_name,
                os.status_id AS office_schedule_status_id,
                CASE
                    WHEN (
                        SELECT 1
                        FROM appointment a
                        WHERE os.id = a.schedule_id
                        AND DATE(a.time_start) = $1
                        AND a.status_id IN (SELECT id FROM appointment_status WHERE name IN ('Pendiente', 'Confirmada'))
                        AND a.time_start::time <= os.time_start::time AND a.time_end::time >= os.time_end::time
                    ) IS NOT NULL THEN 2  -- Ocupado (cita)
                    WHEN (
                        SELECT 1
                        FROM schedule_block sb
                        WHERE os.id = sb.office_schedule_id AND sb.block_date = $1
                        AND (
                            (sb.time_start IS NULL AND sb.time_end IS NULL) OR -- Bloqueo de día completo
                            (os.time_start::time >= sb.time_start::time AND os.time_end::time <= sb.time_end::time) OR
                            (os.time_start::time < sb.time_end::time AND os.time_end::time > sb.time_start::time)
                        )
                    ) IS NOT NULL THEN 3  -- Bloqueado
                    ELSE 1  -- Disponible
                END AS calculated_status_id,
                ROW_NUMBER() OVER (PARTITION BY os.id ORDER BY a.id DESC, sb.id DESC) as rn
            FROM
                office_schedule os
            JOIN office o ON os.office_id = o.id
            LEFT JOIN appointment a ON
                os.id = a.schedule_id AND DATE(a.time_start) = $1
                AND a.status_id IN (SELECT id FROM appointment_status WHERE name IN ('Pendiente', 'Confirmada'))
                AND a.time_start::time <= os.time_start::time AND a.time_end::time >= os.time_end::time
            LEFT JOIN schedule_block sb ON
                os.id = sb.office_schedule_id AND sb.block_date = $1
                AND (
                    (sb.time_start IS NULL AND sb.time_end IS NULL) OR
                    (os.time_start::time >= sb.time_start::time AND os.time_end::time <= sb.time_end::time) OR
                    (os.time_start::time < sb.time_end::time AND os.time_end::time > sb.time_start::time)
                )
            WHERE
                os.service_id = $2
                AND os.shift_id = $3
                AND os.day_of_week = $4
                AND os.status_id = 1
        )
        SELECT
            schedule_id,
            time_start,
            time_end,
            time_duration,
            office_name,
            CASE
                WHEN office_schedule_status_id = 1 AND calculated_status_id = 1 THEN 1  -- Disponible
                ELSE 2  -- No disponible
            END AS status_id
        FROM
            AvailableSchedules
        WHERE
            rn = 1
        ORDER BY
            time_start ASC;
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

	var beneficiaryID interface{}
	if a.BeneficiaryID.Valid {
		beneficiaryID = a.BeneficiaryID.UUID
	} else {
		beneficiaryID = nil
	}

	queryRegisterAppointment := `
		INSERT INTO appointment (
			id, account_id, schedule_id, patient_id, beneficiary_id, 
			time_start, time_end, status_id
		) 
		VALUES (
			$1, $2, $3, $4, $5, 
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
		beneficiaryID,
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

func (ar *appointmentRepository) GetAppointmentForPatient(ctx context.Context, PatientID uuid.UUID) ([]entities.AppointmentForPatient, error) {

	query := `
		SELECT
			appt.id
		    appt.account_id,
		    appt.patient_id,
		    appt.beneficiary_id,
		    CASE 
		        WHEN appt.beneficiary_id IS NOT NULL THEN 
		            CONCAT(bene.first_name, ' ', bene.last_name1, ' ', bene.last_name2)
		        ELSE 
		            CONCAT(pat.first_name, ' ', pat.last_name1, ' ', pat.last_name2)
		    END AS full_name,
		    offc.name AS office_name,
		    serv.name AS service_name,
		    appt.time_start,
		    appt.time_end,
		    stat.name AS status_name
		FROM appointment AS appt
		JOIN appointment_status AS stat ON appt.status_id = stat.id
		JOIN patient AS pat ON appt.patient_id = pat.account_id
		LEFT JOIN beneficiary AS bene ON appt.beneficiary_id = bene.id 
		JOIN office_schedule AS osched ON appt.schedule_id = osched.id
		JOIN office AS offc ON osched.office_id = offc.id
		JOIN services AS serv ON osched.service_id = serv.id
		WHERE appt.patient_id = $1
		ORDER BY appt.created_at DESC;
	`

	rows, err := ar.storage.DbPool.Query(ctx, query, PatientID)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando la consulta: %w", err)
	}
	defer rows.Close()

	var appointments []entities.AppointmentForPatient

	for rows.Next() {
		var appt entities.AppointmentForPatient
		err := rows.Scan(
			&appt.AppointentID,
			&appt.AccountID,
			&appt.PatientID,
			&appt.BeneficiaryID,
			&appt.PatientName,
			&appt.OfficeName,
			&appt.ServiceName,
			&appt.TimeStart,
			&appt.TimeEnd,
			&appt.StatusName,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando fila: %w", err)
		}
		appointments = append(appointments, appt)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando las filas: %w", err)
	}
	return appointments, nil

}

func (ar *appointmentRepository) GetAppointmentByID(ctx context.Context, appointmentID uuid.UUID) (entities.Appointment, error) {
	query := `
	SELECT 
		appt.id AS appointment_id,
		appt.patient_id,
		appt.beneficiary_id,
		os.service_id,
		os.shift_id,
		s.reason,
		s.symptoms,
		appt.time_start,
		appt.time_end
	FROM appointment AS appt
	JOIN office_schedule AS os ON os.id = appt.schedule_id
	JOIN consultation AS s ON s.appointment_id = appt.id
	WHERE appt.id = $1;
	`

	var appt entities.Appointment

	err := ar.storage.DbPool.QueryRow(ctx, query, appointmentID).Scan(
		&appt.ID,
		&appt.PatientID,
		&appt.BeneficiaryID,
		&appt.OfficeSchedule.ServiceID,
		&appt.OfficeSchedule.ShiftID,
		&appt.Consultation.Reason,
		&appt.Consultation.Symptoms,
		&appt.TimeStart,
		&appt.TimeEnd,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return appt, fmt.Errorf("no se encontró la cita con el ID proporcionado: %v", appointmentID)
		}
		return appt, fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	return appt, nil
}
