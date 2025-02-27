package usecase

import (
	"context"
	"sanatorioApp/internal/domain/auth"
	"sanatorioApp/internal/domain/users/http/models"
)

func (u *usecase) AccountVerification(ctx context.Context, token string) (models.Response, error) {
	claims, err := auth.ValidateJWTConfirmation(token)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Token de confirmación inválido",
			Errors:  err.Error(),
		}, nil
	}

	accountID := claims.AccountID

	updated, err := u.repo.AccountVerification(ctx, accountID, true)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Error al verificar la cuenta",
			Errors:  err.Error(),
		}, nil
	}

	if !updated {
		return models.Response{
			Status:  "error",
			Message: "No se pudo verificar la cuenta. No se actualizó ninguna fila.",
		}, nil
	}

	return models.Response{
		Status:  "success",
		Message: "La cuenta ha sido verificada exitosamente",
	}, nil
}
