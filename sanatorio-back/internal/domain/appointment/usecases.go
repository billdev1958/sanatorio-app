package appointment

import (
	"context"
	"sanatorioApp/internal/domain/appointment/http/models"
)

type Usecase interface {
	GetParamsForAppointments(ctx context.Context) (models.Response, error)
	GetSchedulesForAppointment(ctx context.Context, filtersRequest models.SchedulesAppointmentRequest) ([]models.OfficeScheduleResponse, error)
}
