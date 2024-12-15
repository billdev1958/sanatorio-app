package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/cites/entities"
	"sanatorioApp/pkg"

	"github.com/jackc/pgx/v5"
)

func (cr *citesRepository) RegisterOfficeSchedule(ctx context.Context, sc entities.Schedule, os entities.OfficeSchedule) (string, error) {
	// Inicia la transacción
	tx, err := cr.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Inserta el nuevo horario en la tabla `schedule` y obtén el `scheduleID`
	scheduleID, err := cr.insertSchedule(ctx, tx, sc)
	if err != nil {
		return "", fmt.Errorf("failed to insert schedule: %w", err)
	}

	// Inserta el nuevo registro en la tabla `office_schedule`, incluyendo el `scheduleID`
	os.Schedule.ID = scheduleID
	if err := cr.insertOfficeSchedule(ctx, tx, os); err != nil {
		return "", fmt.Errorf("failed to insert office schedule: %w", err)
	}

	// Confirma la transacción
	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("Horario registrado con éxito para la oficina '%d' el día '%d'", os.Office.ID, sc.DayOfWeek), nil
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
		INSERT INTO office_schedule (schedule_id, office_id, shift_id, service_id, doctor_id, status_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (office_id, shift_id, service_id, doctor_id)
		DO NOTHING
	`
	_, err := tx.Exec(ctx, queryInsert, sc.Schedule.ID, sc.Office.ID, sc.ShiftID, sc.Services.ID, sc.DoctorUser.AccountID, sc.OfficeStatus.ID)
	if err != nil {
		log.Printf("error al registrar la programación de la oficina '%d' con turno '%d' y servicio '%d': %v", sc.Office.ID, sc.ShiftID, sc.Services.ID, err)
		return err
	}
	return nil
}

// Filters for servicesID, OfficeID, StatusID, ShiftID, DoctorID, DayOfWeek,
func (cr *citesRepository) GetSchedules(ctx context.Context, filters map[string]interface{}) ([]entities.OfficeSchedule, error) {
	// Definir el mapeo de nombres JSON a columnas de la base de datos
	columnMapping := map[string]string{
		"serviceID":      "os.service_id",
		"doctorID":       "os.doctor_id",
		"officeID":       "os.office_id",
		"officeStatusID": "of_status.id",
		"shiftID":        "os.shift_id",
		"dayOfWeek":      "sc.day_of_week",
	}

	// Usar el método del paquete para traducir los filtros
	dbFilters, err := pkg.MapFiltersToColumns(filters, columnMapping)
	if err != nil {
		return nil, fmt.Errorf("error al traducir filtros: %w", err)
	}

	// Construir la cláusula WHERE dinámica
	whereClause, args, err := pkg.BuildWhereClause(dbFilters)
	if err != nil {
		return nil, fmt.Errorf("error construyendo la cláusula WHERE: %w", err)
	}

	// Construir la consulta completa
	query := fmt.Sprintf(`
	SELECT
		os.id AS office_schedule_id,
		os.service_id,
		os.schedule_id,
		os.office_id,
		os.status_id AS office_schedule_status_id, -- Status de office_schedule
		os.shift_id,
		os.doctor_id,
		s.name AS service_name,
		sc.day_of_week AS day_schedule,
		sc.time_start,
		sc.time_end,
		sc.time_duration,
		o.name AS office_name,
		sh.name AS shift_name,
		d.first_name AS doctor_name,
		d.last_name1 AS doctor_lastname1,
		d.last_name2 AS doctor_lastname2,
		d.medical_license
	FROM office_schedule os
	INNER JOIN office o ON os.office_id = o.id
	INNER JOIN services s ON os.service_id = s.id
	INNER JOIN schedule sc ON os.schedule_id = sc.id
	INNER JOIN cat_shift sh ON os.shift_id = sh.id
	INNER JOIN doctor d ON os.doctor_id = d.account_id
	%s
`, whereClause)

	// Ejecutar la consulta con los argumentos generados
	rows, err := cr.storage.DbPool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Crear un slice para almacenar los resultados
	var responses []entities.OfficeSchedule

	// Iterar sobre las filas
	for rows.Next() {
		var response entities.OfficeSchedule

		err := rows.Scan(
			&response.ID,
			&response.Services.ID,
			&response.Schedule.ID,
			&response.Office.ID,
			&response.OfficeStatus.ID,
			&response.ShiftID,
			&response.DoctorUser.AccountID,
			&response.Services.Name,
			&response.Schedule.DayOfWeek,
			&response.Schedule.TimeStart,
			&response.Schedule.TimeEnd,
			&response.Schedule.TimeDuration,
			&response.Office.Name,
			&response.OfficeStatus.Name,
			&response.ShiftName,
			&response.DoctorUser.FirstName,
			&response.DoctorUser.LastName1,
			&response.DoctorUser.LastName2,
			&response.DoctorUser.MedicalLicense,
		)
		if err != nil {
			return nil, err
		}

		// Agregar el resultado al slice de respuestas
		responses = append(responses, response)
	}

	// Verificar errores durante la iteración
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return responses, nil
}

/*
	&response.ID,

	si
	&response.Services.ID,

	no
	&response.Schedule.ID,

	si OfficeID debe ser una con status no asignado
	&response.Office.ID,
	&response.Office.StatusID,

	si agregar validacion de horas pertenezcan a horario matutino o vespertino
	&response.ShiftID,

	si y no debe tener un horario que coincida con las horas del horario asignado
	&response.DoctorUser.AccountID,

	&response.Services.Name,

	si, validaciones de timestart y timeEnd
	&response.Schedule.DayOfWeek,
	&response.Schedule.TimeStart,
	&response.Schedule.TimeEnd,
	&response.Schedule.TimeDuration,

	no
	&response.Office.Name,
	no
	&response.StatusName,
	no
	&response.ShiftName,

	// esto no
	&response.DoctorUser.FirstName,
	&response.DoctorUser.LastName1,
	&response.DoctorUser.LastName2,
	&response.DoctorUser.MedicalLicense,
*/
