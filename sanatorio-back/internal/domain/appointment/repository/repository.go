package repository

import (
	"context"
	"sanatorioApp/internal/domain/appointment"
	"sanatorioApp/internal/domain/appointment/entities"
	postgres "sanatorioApp/internal/infraestructure/db"
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
		        WHEN a.id IS NOT NULL THEN 2 -- Ocupado
		        WHEN sb.id IS NOT NULL THEN 3 -- Bloqueado
		        ELSE 1 -- Disponible
		    END AS status_id
		FROM
		    office_schedule os
		JOIN office o ON os.office_id = o.id
		LEFT JOIN appointment a ON
		    os.id = a.schedule_id AND DATE(a.time_start) = $1 -- Comparar fecha completa en citas
		LEFT JOIN schedule_block sb ON
		    os.id = sb.schedule_id AND sb.block_date = $1 -- Comparar fecha completa en bloqueos
		WHERE
		    os.service_id = $2 -- Filtrar por servicio
		    AND os.shift_id = $3 -- Filtrar por turno
		    AND os.day_of_week = $4 -- Filtrar por día de la semana
		    AND os.status_id = 1 -- Solo horarios activos
		ORDER BY os.time_start ASC;
    `

	rows, err := ar.storage.DbPool.Query(ctx, query, date, serviceID, shiftID, dayOfWeek)
	if err != nil {
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
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return schedules, nil
}
