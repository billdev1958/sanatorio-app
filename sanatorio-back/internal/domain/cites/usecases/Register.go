package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/cites"
	"sanatorioApp/internal/domain/cites/entities"
	users "sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/pkg"

	"sanatorioApp/internal/domain/cites/http/models"
	"time"
)

type usecase struct {
	repo cites.CitesRepository
}

func NewUsecase(repo cites.CitesRepository) cites.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterOfficeSchedule(ctx context.Context, request models.RegisterOfficeScheduleRequest) (string, error) {
	// Definir el formato de hora
	layout := "15:04"

	// Parsear TimeStart desde el request
	startTime, err := time.Parse(layout, request.TimeStart)
	if err != nil {
		return "", fmt.Errorf("invalid time format for TimeStart: %w", err)
	}

	// Parsear TimeEnd desde el request
	endTime, err := time.Parse(layout, request.TimeEnd)
	if err != nil {
		return "", fmt.Errorf("invalid time format for TimeEnd: %w", err)
	}

	// Calcula la duracion de las horas ingresadas
	timeDuration := endTime.Sub(startTime)

	// Crear la entidad Schedule a partir del request
	schedule := entities.Schedule{
		DayOfWeek:    request.DayOfWeek,
		TimeStart:    startTime,
		TimeEnd:      endTime,
		TimeDuration: timeDuration,
	}

	// Crear la entidad OfficeSchedule a partir del request
	officeSchedule := entities.OfficeSchedule{
		Office: entities.Office{
			ID: request.OfficeID,
		},
		ShiftID: request.ShiftID,
		Services: entities.Services{
			ID: request.ServiceID,
		},
		DoctorUser: users.DoctorUser{
			AccountID: request.DoctorID,
		},
		OfficeStatus: entities.OfficeStatus{
			ID: int(entities.OfficeStatusAvailable),
		},
	}

	// Llama al repositorio para registrar el horario
	message, err := u.repo.RegisterOfficeSchedule(ctx, schedule, officeSchedule)
	if err != nil {
		return "", fmt.Errorf("failed to register schedule: %w", err)
	}
	return message, nil
}

func (u *usecase) RegisterAppointment(ctx context.Context, request models.RegisterAppointmentRequest) (string, error) {

	appointment := entities.Appointment{
		PatientAccountID: request.PatientAccountID,
		OfficeID:         request.OfficeID,
		Date:             request.Date,
		ScheduleID:       request.ScheduleID,
		StatusID:         int(entities.AppointmentStatusPendiente),
	}

	message, err := u.repo.RegisterAppointment(ctx, appointment)
	if err != nil {
		return "", fmt.Errorf("failed to register appointment: %w", err)
	}
	return message, nil
}

func (u *usecase) GetSchedules(ctx context.Context, filtersRequest models.OfficeSCheduleFiltersRequest) ([]models.OfficeScheduleResponse, error) {
	// Convertir la estructura en filtros dinámicos
	filters, err := pkg.Filter(filtersRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to process filters: %w", err)
	}

	// Llamar al repositorio con los filtros
	schedules, err := u.repo.GetSchedules(ctx, filters)
	if err != nil {
		return nil, err
	}

	// Crear un slice para almacenar los modelos transformados
	var responses []models.OfficeScheduleResponse

	// Transformar cada entidad en un modelo de respuesta
	for _, schedule := range schedules {
		response := models.OfficeScheduleResponse{
			OfficeScheduleID: schedule.ID,
			ServiceID:        schedule.Services.ID,
			ScheduleID:       schedule.Schedule.ID,
			OfficeID:         schedule.Office.ID,
			OfficeStatusID:   schedule.Services.ID,
			ShiftID:          schedule.ShiftID,
			DoctorID:         schedule.DoctorUser.AccountID,
			ServiceName:      schedule.Services.Name,
			DaySchedule:      schedule.Schedule.DayOfWeek,
			TimeStart:        schedule.Schedule.TimeStart.Format("15:04"),
			TimeEnd:          schedule.Schedule.TimeEnd.Format("15:04"),
			TimeDuration:     schedule.Schedule.TimeDuration.String(),
			OfficeName:       schedule.Office.Name,
			OfficeStatusName: schedule.OfficeStatus.Name,
			ShiftName:        schedule.ShiftName,
			DoctorName:       schedule.DoctorUser.FirstName,
			DoctorLastName1:  schedule.DoctorUser.LastName1,
			DoctorLastName2:  schedule.DoctorUser.LastName2,
			MedicalLicense:   schedule.DoctorUser.MedicalLicense,
		}

		// Agregar al slice de respuestas
		responses = append(responses, response)
	}

	return responses, nil
}
