package usecases

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/appointment"
	"sanatorioApp/internal/domain/appointment/http/models"
	"sanatorioApp/internal/domain/catalogs"
	"sanatorioApp/pkg"
)

type usecase struct {
	repo        appointment.AppointmentRepository
	catalogRepo catalogs.CatalogsRepository
}

func NewUsecase(repo appointment.AppointmentRepository, catalogRepo catalogs.CatalogsRepository) appointment.Usecase {
	return &usecase{repo: repo, catalogRepo: catalogRepo}
}

func (u *usecase) GetParamsForAppointments(ctx context.Context) (models.Response, error) {
	services, err := u.catalogRepo.GetServices(ctx)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Error al obtener los servicios",
			Errors:  err.Error(),
		}, nil
	}

	shifts, err := u.catalogRepo.GetShifts(ctx)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Error al obtener los turnos",
			Errors:  err.Error(),
		}, nil
	}

	return models.Response{
		Status: "success",
		Data: map[string]interface{}{
			"services": services,
			"shifts":   shifts,
		},
	}, nil
}

func (u *usecase) GetSchedulesForAppointment(ctx context.Context, filtersRequest models.SchedulesAppointmentRequest) ([]models.OfficeScheduleResponse, error) {
	filters, err := pkg.Filter(filtersRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to process filters: %w", err)
	}

	schedules, err := u.catalogRepo.GetSchedulesForAppointment(ctx, filters)
	if err != nil {
		return nil, err
	}

	var responses []models.OfficeScheduleResponse
	for _, schedule := range schedules {
		response := models.OfficeScheduleResponse{
			ID:           schedule.ID,
			OfficeID:     schedule.OfficeID,
			ShiftID:      schedule.ShiftID,
			DoctorID:     schedule.DoctorID,
			StatusID:     schedule.StatusID,
			DayOfWeek:    schedule.DayOfWeek,
			TimeStart:    schedule.TimeStart.Format("15:04"),
			TimeEnd:      schedule.TimeEnd.Format("15:04"),
			TimeDuration: schedule.TimeDuration.String(),
		}
		responses = append(responses, response)
	}

	return responses, nil
}
