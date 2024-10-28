package cites

import (
	"context"
	"sanatorioApp/internal/domain/cites/http/models"
)

type Usecase interface {
	RegisterSpecialty(ctx context.Context, sp models.RegisterSpecialtyRequest) (string, error)

	RegisterOffice(ctx context.Context, office models.RegisterOfficeRequest) (string, error)

	RegisterSchedule(ctx context.Context, sh models.RegisterScheduleRequest) (string, error)

	RegisterAppointment(ctx context.Context, appointment models.RegisterAppointmentRequest) (string, error)
}
