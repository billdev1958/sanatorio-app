package usecase

import (
	"context"
	"fmt"
)

func (u *usecase) DeleteUser(ctx context.Context, accountID string) (string, error) {
	message, err := u.repo.DeleteUser(ctx, accountID)
	if err != nil {
		return "", fmt.Errorf("failed to delete user: %w", err)
	}
	return message, nil
}

// SoftDeleteUser marca un usuario como eliminado sin borrar su información.
func (u *usecase) SoftDeleteUser(ctx context.Context, accountID string) (string, error) {
	message, err := u.repo.SoftDeleteUser(ctx, accountID)
	if err != nil {
		return "", fmt.Errorf("failed to soft delete user: %w", err)
	}
	return message, nil
}
