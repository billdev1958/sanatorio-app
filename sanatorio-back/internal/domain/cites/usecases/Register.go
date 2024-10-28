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
		Name:        request.Name,
		SpecialtyID: request.SpecialtyID,
	}

	message, err := u.repo.RegisterOffice(ctx, office)
	if err != nil {
		return "", fmt.Errorf("failed to register office %w:", err)
	}

	return message, nil
}

func (u *usecase) RegisterSchedule(ctx context.Context, request models.RegisterScheduleRequest) (string, error) {
	layout := "15:04"

	startTime, err := time.Parse(layout, request.TimeStart)
	if err != nil {
		return "", fmt.Errorf("invalid time format for TimeStart: %w", err)
	}

	endTime, err := time.Parse(layout, request.TimeEnd)
	if err != nil {
		return "", fmt.Errorf("invalid time format for TimeEnd: %w", err)
	}

	schedule := entities.Schedule{
		OfficeID:  request.OfficeID,
		DayOfWeek: request.DayOfWeek,
		TimeStart: startTime,
		TimeEnd:   endTime,
	}

	// Llama al repositorio para registrar el horario
	message, err := u.repo.RegisterSchedule(ctx, schedule)
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
