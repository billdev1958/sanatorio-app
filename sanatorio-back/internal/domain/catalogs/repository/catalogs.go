package repository

import (
	"context"
	"fmt"

	"sanatorioApp/internal/domain/catalogs"
	"sanatorioApp/internal/domain/catalogs/models"
	postgres "sanatorioApp/internal/infraestructure/db"
	"sanatorioApp/pkg"

	"github.com/google/uuid"
)

type catalogRepository struct {
	storage *postgres.PgxStorage
}

func NewCatalogRepository(storage *postgres.PgxStorage) catalogs.CatalogsRepository {
	return &catalogRepository{storage: storage}
}

func (cr *catalogRepository) GetServices(ctx context.Context) ([]models.Services, error) {
	var services []models.Services
	query := "SELECT id, name FROM services"

	rows, err := cr.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los servicios: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var service models.Services
		if err := rows.Scan(&service.ID, &service.Name); err != nil {
			return nil, fmt.Errorf("error al escanear servicio: %w", err)
		}
		services = append(services, service)
	}
	return services, nil
}

func (cr *catalogRepository) GetShifts(ctx context.Context) ([]models.CatShift, error) {
	var shifts []models.CatShift

	query := "SELECT id, name FROM cat_shift"

	rows, err := cr.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener shifts: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var shift models.CatShift
		if err := rows.Scan(&shift.ID, &shift.Name); err != nil {
			return nil, fmt.Errorf("error al escanear shifts: %w", err)
		}
		shifts = append(shifts, shift)
	}
	return shifts, nil
}

func (cr *catalogRepository) GetDays(ctx context.Context) ([]models.DayOfWeek, error) {
	var days []models.DayOfWeek
	query := "SELECT day_of_week, name FROM days"

	rows, err := cr.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener days: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var day models.DayOfWeek
		if err := rows.Scan(&day.ID, &day.Name); err != nil {
			return nil, fmt.Errorf("error al escanear days: %w", err)
		}
		days = append(days, day)
	}
	return days, nil
}

func (cr *catalogRepository) GetDoctors(ctx context.Context) ([]models.Doctor, error) {
	var doctors []models.Doctor
	query := "SELECT account_id, first_name, last_name1, last_name2 FROM doctor"

	rows, err := cr.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener days: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var doctor models.Doctor
		if err := rows.Scan(&doctor.AccountID, &doctor.FirstName, &doctor.LastName1, &doctor.LastName1); err != nil {
			return nil, fmt.Errorf("error al escanear days: %w", err)
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

func (cr *catalogRepository) GetOffices(ctx context.Context) ([]models.Office, error) {
	query := `
		SELECT 
			o.id AS office_id,
			o.name AS office_name
		FROM office o
	`

	rows, err := cr.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offices []models.Office

	for rows.Next() {
		var office models.Office

		err := rows.Scan(
			&office.ID,
			&office.Name,
		)
		if err != nil {
			return nil, err
		}

		offices = append(offices, office)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return offices, nil
}

func (cr *catalogRepository) GetSchedulesForAppointment(ctx context.Context, filters map[string]interface{}) ([]models.OfficeSchedule, error) {
	columnMapping := map[string]string{
		"shift":   "shift_id",
		"service": "service_id",
		"day":     "day_of_week",
	}

	dbFilters, err := pkg.MapFiltersToColumns(filters, columnMapping)
	if err != nil {
		return nil, fmt.Errorf("error translating filters: %w", err)
	}

	// Build dynamic WHERE clause
	whereClause, args, err := pkg.BuildWhereClause(dbFilters)
	if err != nil {
		return nil, fmt.Errorf("error building WHERE clause: %w", err)
	}

	query := fmt.Sprintf(`
		SELECT 
			id,
			office_id,
			shift_id,
			service_id,
			doctor_id,
			status_id,
			day_of_week,
			time_start,
			time_end,
			time_duration
		FROM office_schedule
		%s
	`, whereClause)

	rows, err := cr.storage.DbPool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// Parse query results
	var schedules []models.OfficeSchedule
	for rows.Next() {
		var schedule models.OfficeSchedule
		if err := rows.Scan(
			&schedule.ID,
			&schedule.OfficeID,
			&schedule.ShiftID,
			&schedule.ServiceID,
			&schedule.DoctorID,
			&schedule.StatusID,
			&schedule.DayOfWeek,
			&schedule.TimeStart,
			&schedule.TimeEnd,
			&schedule.TimeDuration,
		); err != nil {
			return nil, fmt.Errorf("error scanning schedule row: %w", err)
		}
		schedules = append(schedules, schedule)
	}

	// Check for iteration errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return schedules, nil
}

func (cr *catalogRepository) GetPatientAndBeneficiaries(ctx context.Context, accountID uuid.UUID) (models.PatientAndBenefeciaries, error) {
	query := `
        SELECT 
            b.id AS beneficiary_id,
            CONCAT(b.first_name, ' ', b.last_name1, ' ', b.last_name2) AS beneficiary_full_name,
            b.account_holder AS account_holder_id,
            CONCAT(p.first_name, ' ', p.last_name1, ' ', p.last_name2) AS account_holder_full_name
        FROM 
            beneficiary b
        JOIN 
            patient p ON b.account_holder = p.account_id
        WHERE 
            b.account_holder = $1
    `

	rows, err := cr.storage.DbPool.Query(ctx, query, accountID)
	if err != nil {
		return models.PatientAndBenefeciaries{}, fmt.Errorf("error fetching beneficiaries: %w", err)
	}
	defer rows.Close()

	var result models.PatientAndBenefeciaries
	var beneficiaries []models.Beneficiary

	for rows.Next() {
		var beneficiary models.Beneficiary
		var accountHolderID uuid.UUID
		var accountHolderFullName string

		if err := rows.Scan(
			&beneficiary.ID,
			&beneficiary.FullName,
			&accountHolderID,
			&accountHolderFullName,
		); err != nil {
			return models.PatientAndBenefeciaries{}, fmt.Errorf("error scanning row: %w", err)
		}

		if result.AccountHolderID == uuid.Nil {
			result.AccountHolderID = accountHolderID
			result.FullName = accountHolderFullName
		}

		beneficiaries = append(beneficiaries, beneficiary)
	}

	if err := rows.Err(); err != nil {
		return models.PatientAndBenefeciaries{}, fmt.Errorf("error iterating rows: %w", err)
	}

	result.Beneficiaries = beneficiaries

	return result, nil
}
