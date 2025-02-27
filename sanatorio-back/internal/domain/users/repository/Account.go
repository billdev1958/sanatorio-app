package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (ur *userRepository) AccountVerification(ctx context.Context, accountID uuid.UUID, isValid bool) (bool, error) {
	query := `
        UPDATE account
        SET is_verified = $1
        WHERE id = $2
    `
	cmdTag, err := ur.storage.DbPool.Exec(ctx, query, isValid, accountID)
	if err != nil {
		return false, fmt.Errorf("error updating account verification: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return false, errors.New("no rows updated: account not found o nada que actualizar")
	}

	return true, nil
}
