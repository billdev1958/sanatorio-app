package catalogs

import (
	"context"
	"sanatorioApp/internal/domain/catalogs/models"
)

type CatalogsRepository interface {
	GetServices(ctx context.Context) ([]models.Services, error)
	GetShifts(ctx context.Context) ([]models.CatShift, error)
	GetDays(ctx context.Context) ([]models.DayOfWeek, error)
	GetDoctors(ctx context.Context) ([]models.Doctor, error)
	GetOffices(ctx context.Context) ([]models.Office, error)
}
