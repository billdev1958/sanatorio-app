package cites

import (
	"context"
	"sanatorioApp/internal/domain/cites/http/models"
)

type Usecase interface {
	RegisterSpecialty(ctx context.Context, sp models.RegisterSpecialtyRequest) (string, error)

	RegisterOfficeSchedule(ctx context.Context, request models.RegisterOfficeScheduleRequest) (string, error)

	RegisterOffice(ctx context.Context, office models.RegisterOfficeRequest) (string, error)

	RegisterAppointment(ctx context.Context, appointment models.RegisterAppointmentRequest) (string, error)

	GetSchedules(ctx context.Context) ([]models.OfficeScheduleResponse, error)
}
