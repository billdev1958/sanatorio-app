package usecases

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/catalogs"
	schedule "sanatorioApp/internal/domain/schedules"
	"sanatorioApp/internal/domain/schedules/http/models"
)

// Get offices *
// Get services *
// Get doctors *
type usecase struct {
	catalogRepo catalogs.CatalogsRepository
}

func NewUsecase(catalogRepo catalogs.CatalogsRepository) schedule.OfficeSchedule {
	return &usecase{
		catalogRepo: catalogRepo,
	}
}
func (u *usecase) GetInfoOfficeSchedule(ctx context.Context) (models.GetInfoOfficeSchedule, error) {
	days, err := u.catalogRepo.GetDays(ctx)
	if err != nil {
		return models.GetInfoOfficeSchedule{}, fmt.Errorf("error al obtener días: %w", err)
	}

	services, err := u.catalogRepo.GetServices(ctx)
	if err != nil {
		return models.GetInfoOfficeSchedule{}, fmt.Errorf("error al obtener services: %w", err)

	}

	shifts, err := u.catalogRepo.GetShifts(ctx)
	if err != nil {
		return models.GetInfoOfficeSchedule{}, fmt.Errorf("error al obtener shifts: %w", err)

	}

	doctors, err := u.catalogRepo.GetDoctors(ctx)
	if err != nil {
		return models.GetInfoOfficeSchedule{}, fmt.Errorf("error al obtener doctores: %w", err)
	}

	response := models.GetInfoOfficeSchedule{
		CatDays:     days,
		CatShifts:   shifts,
		CatServices: services,
		Doctors:     doctors,
	}

	return response, nil

}
