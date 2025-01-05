package appointment

import (
	"context"
	"sanatorioApp/internal/domain/appointment/http/models"

	"github.com/google/uuid"
)

type Usecase interface {
	GetParamsForAppointments(ctx context.Context, accountID uuid.UUID) (models.Response, error)
	GetAvaliableSchedules(ctx context.Context, params models.SchedulesAppointmentRequest) ([]models.OfficeScheduleResponse, error)
}
