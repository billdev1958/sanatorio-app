package usecases

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/appointment"
	"sanatorioApp/internal/domain/appointment/entities"
	"sanatorioApp/internal/domain/appointment/http/models"
	"sanatorioApp/internal/domain/catalogs"
	"time"

	"github.com/google/uuid"
)

type usecase struct {
	repo        appointment.AppointmentRepository
	catalogRepo catalogs.CatalogsRepository
}

func NewUsecase(repo appointment.AppointmentRepository, catalogRepo catalogs.CatalogsRepository) appointment.Usecase {
	return &usecase{repo: repo, catalogRepo: catalogRepo}
}

func (u *usecase) GetParamsForAppointments(ctx context.Context, accountID uuid.UUID) (models.Response, error) {
	patients, err := u.catalogRepo.GetPatientAndBeneficiaries(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Error al obtener los pacientes",
			Errors:  err.Error(),
		}, nil
	}

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
			"patients": patients,
			"services": services,
			"shifts":   shifts,
		},
	}, nil
}

func (u *usecase) GetAvaliableSchedules(ctx context.Context, params models.SchedulesAppointmentRequest) ([]models.OfficeScheduleResponse, error) {
	if params.AppointmentDate == "" {
		return nil, errors.New("appointmentDate is required")
	}

	appointmentDate, err := time.Parse(time.RFC3339, params.AppointmentDate)
	if err != nil {
		return nil, errors.New("invalid appointmentDate format, expected ISO 8601 (e.g., 2025-01-14T06:00:00.000Z)")
	}

	dayOfWeek := int(appointmentDate.Weekday())

	formattedDate := appointmentDate.Format("2006-01-02")

	schedules, err := u.repo.GetAvaliableSchedules(ctx, formattedDate, dayOfWeek, params.Service, params.Shift)
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

func (u *usecase) RegisterAppointment(ctx context.Context, accountID uuid.UUID, request models.RegisterAppointmentRequest) (message string, err error) {
	log.Printf("Inicio del registro de appointment. AccountID: %s, ScheduleID: %d", accountID, request.ScheduleID)

	if request.TimeStart.After(request.TimeEnd) {
		log.Printf("Error de validación: TimeStart (%s) es después de TimeEnd (%s)", request.TimeStart, request.TimeEnd)
		return "", fmt.Errorf("la hora de inicio no puede ser después de la hora de fin")
	}

	appointment := entities.Appointment{
		ID:            uuid.New(),
		AccountID:     accountID,
		ScheduleID:    request.ScheduleID,
		PatientID:     request.PatientID,
		BeneficiaryID: request.BeneficiaryID,
		TimeStart:     request.TimeStart,
		TimeEnd:       request.TimeEnd,
		StatusID:      int(appointment.AppointmentStatusPendiente),
	}
	log.Printf("Creado el objeto Appointment: %+v", appointment)

	consultation := entities.Consultation{
		Reason:   request.Reason,
		Symptoms: request.Symptoms,
	}
	log.Printf("Creado el objeto Consultation: %+v", consultation)

	log.Println("Llamando al repositorio para registrar el appointment.")
	success, err := u.repo.RegisterAppointment(ctx, appointment, consultation)
	if err != nil {
		log.Printf("Error al registrar el appointment en el repositorio: %v", err)
		return "", fmt.Errorf("error al registrar el appointment: %w", err)
	}

	if success {
		log.Println("El appointment se registró exitosamente.")
		return "El appointment se registró exitosamente.", nil
	}

	log.Println("No se pudo registrar el appointment por razones desconocidas.")
	return "", fmt.Errorf("no se pudo registrar el appointment por razones desconocidas")
}
