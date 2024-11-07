package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/cites"
	"sanatorioApp/internal/domain/cites/entities"
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

	specialty := entities.Specialty{
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
		Name:      request.Name,
		ServiceID: request.ServiceID,
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

	timeDuration, err := time.ParseDuration(request.TimeDuration)
	if err != nil {
		return "", fmt.Errorf("invalid duration format for TimeDuration: %w", err)
	}

	// Crear la entidad Schedule a partir del request
	schedule := entities.Schedule{
		DayOfWeek:    request.DayOfWeek,
		TimeStart:    startTime,
		TimeEnd:      endTime,
		TimeDuration: timeDuration,
	}

	// Crear la entidad OfficeSchedule a partir del request
	officeSchedule := entities.OfficeSchedule{
		OfficeID:  request.OfficeID,
		ShiftID:   request.ShiftID,
		ServiceID: request.ServiceID,
		DoctorID:  request.DoctorID,
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
