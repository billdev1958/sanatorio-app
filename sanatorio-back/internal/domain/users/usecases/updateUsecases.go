package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/auth"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
)

func (u *usecase) UpdateUser(ctx context.Context, userUpdate models.UpdateUser) (string, error) {
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return "", fmt.Errorf("unauthorized: no claims found in context")
	}

	updateAccount := entities.Account{
		Email:    userUpdate.Email,
		Password: userUpdate.Password,
	}

	updateUser := entities.PatientUser{
		FirstName: userUpdate.Name,
		LastName1: userUpdate.Lastname1,
		LastName2: userUpdate.Lastname2,
		Curp:      userUpdate.Curp,
	}

	message, err := u.repo.UpdatePatient(ctx, updateAccount, updateUser)
	if err != nil {
		return "", fmt.Errorf("failed to update user: %w", err)
	}

	return message, nil
}
