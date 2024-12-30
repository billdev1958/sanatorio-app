package appointment

import (
	"context"
	"sanatorioApp/internal/domain/appointment/http/models"
)

type Usecase interface {
	GetSchedulesForAppointment(ctx context.Context, filtersRequest models.SchedulesAppointmentRequest) ([]models.OfficeScheduleResponse, error)
}
