package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/cites"
	"sanatorioApp/internal/domain/cites/entities"
	users "sanatorioApp/internal/domain/users/entities"

	"sanatorioApp/internal/domain/cites/http/models"
	"time"
)

type usecase struct {
	repo cites.CitesRepository
}

func NewUsecase(repo cites.CitesRepository) cites.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterSpecialty(ctx context.Context, request models.RegisterSpecialtyRequest) (string, error) {

	specialty := entities.Services{
		Name: request.Name,
	}

	message, err := u.repo.RegisterSpecialty(ctx, specialty)
	if err != nil {
		return "", fmt.Errorf("failed to register specialty %w: ", err)
	}

	return message, nil
}

func (u *usecase) RegisterOffice(ctx context.Context, request models.RegisterOfficeRequest) (string, error) {
	office := entities.Office{
		Name: request.Name,
	}

	message, err := u.repo.RegisterOffice(ctx, office)
	if err != nil {
		return "", fmt.Errorf("failed to register office %w:", err)
	}

	return message, nil
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

func (u *usecase) GetSchedules(ctx context.Context) ([]models.OfficeScheduleResponse, error) {
	// Obtener las entidades desde el repositorio
	schedules, err := u.repo.GetSchedules(ctx)
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
			OfficeStatus:     schedule.Office.StatusID,
			ShiftID:          schedule.ShiftID,
			DoctorID:         schedule.DoctorUser.AccountID,
			ServiceName:      schedule.Services.Name,
			DaySchedule:      schedule.Schedule.DayOfWeek,
			TimeStart:        schedule.Schedule.TimeStart.Format("15:04"), // Formato requerido
			TimeEnd:          schedule.Schedule.TimeEnd.Format("15:04"),   // Formato requerido
			TimeDuration:     schedule.Schedule.TimeDuration.String(),     // "HH:MM"
			OfficeName:       schedule.Office.Name,
			OfficeStatusName: schedule.StatusName,
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

// Get by various id for example AccountID, ServicesID, OfficeID, StatusID, ShiftID
