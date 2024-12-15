package schedule

import (
	"context"
	"sanatorioApp/internal/domain/schedules/http/models"
)

type OfficeSchedule interface {
	GetInfoOfficeSchedule(ctx context.Context) (models.GetInfoOfficeSchedule, error)
}
