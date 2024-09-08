package usecase

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"
)

func (u *usecase) DeleteUser(ctx context.Context, accountID string) (models.Response, error) {
	message, err := u.repo.DeleteUser(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to delete user",
			Errors:  map[string]string{"delete": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}

// DeleteDoctor elimina un doctor y su cuenta asociada de la base de datos.
func (u *usecase) DeleteDoctor(ctx context.Context, accountID string) (models.Response, error) {
	message, err := u.repo.DeleteDoctor(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to delete doctor",
			Errors:  map[string]string{"delete": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}

// SoftDeleteUser marca un usuario como eliminado sin borrar su información.
func (u *usecase) SoftDeleteUser(ctx context.Context, accountID string) (models.Response, error) {
	message, err := u.repo.SoftDeleteUser(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to soft delete user",
			Errors:  map[string]string{"soft_delete": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}

// SoftDeleteDoctor marca un doctor como eliminado sin borrar su información.
func (u *usecase) SoftDeleteDoctor(ctx context.Context, accountID string) (models.Response, error) {
	message, err := u.repo.SoftDeleteDoctor(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to soft delete doctor",
			Errors:  map[string]string{"soft_delete": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}
