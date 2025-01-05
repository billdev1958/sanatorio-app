package catalogs

import (
	"context"
	"sanatorioApp/internal/domain/catalogs/models"

	"github.com/google/uuid"
)

type CatalogsRepository interface {
	GetServices(ctx context.Context) ([]models.Services, error)
	GetShifts(ctx context.Context) ([]models.CatShift, error)
	GetDays(ctx context.Context) ([]models.DayOfWeek, error)
	GetDoctors(ctx context.Context) ([]models.Doctor, error)
	GetOffices(ctx context.Context) ([]models.Office, error)
	GetSchedulesForAppointment(ctx context.Context, filters map[string]interface{}) ([]models.OfficeSchedule, error)
	GetPatientAndBeneficiaries(ctx context.Context, accountID uuid.UUID) (models.PatientAndBenefeciaries, error)
}
