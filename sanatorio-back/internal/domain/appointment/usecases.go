package appointment

import (
	"context"
	"sanatorioApp/internal/domain/appointment/http/models"
)

type Usecase interface {
	GetParamsForAppointments(ctx context.Context) (models.Response, error)
	GetAvaliableSchedules(ctx context.Context, params models.SchedulesAppointmentRequest) ([]models.OfficeScheduleResponse, error)
}
