package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/cites/entities"
	"sanatorioApp/internal/domain/cites/http/models"
	"time"
)

func (u *usecase) RegisterOffice(ctx context.Context, request models.RegisterOfficeRequest) (string, error) {
	office := entities.Office{
		Name: request.Name,
	}

	message, err := u.repo.RegisterOffice(ctx, office)
	if err != nil {
		return "", fmt.Errorf("failed to register office %w:", err)
	}

	return message, nil
}

func (u *usecase) GetOffices(ctx context.Context) ([]models.OfficeResponse, error) {
	offices, err := u.repo.GetOffices(ctx)
	if err != nil {
		return nil, fmt.Errorf("error al obtener oficinas: %w", err)
	}

	var responses []models.OfficeResponse

	for _, office := range offices {
		response := models.OfficeResponse{
			OfficeID:   office.ID,
			OfficeName: office.Name,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (u *usecase) UpdateOffice(ctx context.Context, request models.UpdateOfficeRequest) (string, error) {
	office := entities.Office{
		ID:        request.OfficeID,
		Name:      request.OfficeName,
		UpdatedAt: time.Now(),
	}

	update, err := u.repo.UpdateOffice(ctx, office)
	if err != nil {
		return "", fmt.Errorf("failed to update office: %w", err)
	}

	return update, nil
}
