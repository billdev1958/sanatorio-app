package usecases

import (
	"context"
	"errors"
	"sanatorioApp/internal/domain/appointment"
	"sanatorioApp/internal/domain/appointment/http/models"
	"sanatorioApp/internal/domain/catalogs"
	"time"
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

func (u *usecase) GetAvaliableSchedules(ctx context.Context, params models.SchedulesAppointmentRequest) ([]models.OfficeScheduleResponse, error) {
	if params.AppointmentDate == "" {
		return nil, errors.New("appointmentDate is required")
	}

	appointmentDate, err := time.Parse("2006-01-02", params.AppointmentDate)
	if err != nil {
		return nil, errors.New("invalid appointmentDate format, expected YYYY-MM-DD")
	}

	dayOfWeek := int(appointmentDate.Weekday())

	schedules, err := u.repo.GetAvaliableSchedules(ctx, params.AppointmentDate, dayOfWeek, params.Service, params.Shift)
	if err != nil {
		return nil, err
	}

	var response []models.OfficeScheduleResponse
	for _, schedule := range schedules {
		response = append(response, models.OfficeScheduleResponse{
			ID:           schedule.ID,
			TimeStart:    schedule.TimeStart.Format("15:04:05"),
			TimeEnd:      schedule.TimeEnd.Format("15:04:05"),
			TimeDuration: schedule.TimeDuration.String(),
			OfficeName:   schedule.OfficeName,
			StatusID:     schedule.StatusID,
		})
	}

	return response, nil
}
