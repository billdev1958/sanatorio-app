package appointment

import (
	"context"
	"sanatorioApp/internal/domain/appointment/http/models"

	"github.com/google/uuid"
)

type Usecase interface {
	GetParamsForAppointments(ctx context.Context, accountID uuid.UUID) (models.Response, error)
	GetAvaliableSchedules(ctx context.Context, params models.SchedulesAppointmentRequest) ([]models.OfficeScheduleResponse, error)
	RegisterAppointment(ctx context.Context, accountID uuid.UUID, request models.RegisterAppointmentRequest) (message string, err error)
	GetAppointmentForPatient(ctx context.Context, patientID uuid.UUID) (models.Response, error)
}
