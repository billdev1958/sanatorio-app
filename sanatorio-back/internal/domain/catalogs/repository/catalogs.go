package repository

import (
	"context"
	"fmt"

	"sanatorioApp/internal/domain/catalogs"
	"sanatorioApp/internal/domain/catalogs/models"
	postgres "sanatorioApp/internal/infraestructure/db"
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
