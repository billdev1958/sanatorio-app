package cites

import (
	"context"
	"sanatorioApp/internal/domain/cites/http/models"
)

type Usecase interface {
	RegisterOfficeSchedule(ctx context.Context, request models.RegisterOfficeScheduleRequest) (string, error)

	RegisterOffice(ctx context.Context, office models.RegisterOfficeRequest) (string, error)

	UpdateOffice(ctx context.Context, request models.UpdateOfficeRequest) (string, error)

	RegisterAppointment(ctx context.Context, appointment models.RegisterAppointmentRequest) (string, error)

	GetSchedules(ctx context.Context, filtersRequest models.OfficeSCheduleFiltersRequest) ([]models.OfficeScheduleResponse, error)

	GetOffices(ctx context.Context) ([]models.OfficeResponse, error)
}
