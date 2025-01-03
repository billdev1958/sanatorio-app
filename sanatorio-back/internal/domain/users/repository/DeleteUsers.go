package repository

import (
	"context"
	"fmt"
)

// DeleteUser elimina un usuario y su cuenta asociada de la base de datos.
func (ur *userRepository) DeleteUser(ctx context.Context, accountID string) (string, error) {
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Verificar si existe un registro en 'account' con el account_id proporcionado
	var exists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM account WHERE id = $1)"
	err = tx.QueryRow(ctx, checkQuery, accountID).Scan(&exists)
	if err != nil {
		return "", fmt.Errorf("failed to check existence in account: %w", err)
	}

	if !exists {
		return "", fmt.Errorf("user with account ID %s not found", accountID)
	}

	// Eliminar el usuario asociado en la tabla 'users' utilizando el user_id de la tabla 'account'
	userQuery := `
		DELETE FROM users 
		WHERE id = (SELECT user_id FROM account WHERE id = $1)
		RETURNING name
	`
	var deletedName string
	err = tx.QueryRow(ctx, userQuery, accountID).Scan(&deletedName)
	if err != nil {
		return "", fmt.Errorf("failed to delete user: %w", err)
	}

	// Eliminar el registro de la tabla 'account'
	accountQuery := "DELETE FROM account WHERE id = $1"
	_, err = tx.Exec(ctx, accountQuery, accountID)
	if err != nil {
		return "", fmt.Errorf("failed to delete account: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("User %s deleted successfully", deletedName), nil
}

// SoftDeleteUser marca un usuario como eliminado sin borrar su información.
func (ur *userRepository) SoftDeleteUser(ctx context.Context, accountID string) (string, error) {
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Verificar si existe un registro en 'account' con el account_id proporcionado
	var exists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM account WHERE id = $1)"
	err = tx.QueryRow(ctx, checkQuery, accountID).Scan(&exists)
	if err != nil {
		return "", fmt.Errorf("failed to check existence in account: %w", err)
	}

	if !exists {
		return "", fmt.Errorf("user with account ID %s not found", accountID)
	}

	// Marcar como eliminado el usuario en la tabla 'users' utilizando el user_id de la tabla 'account'
	query := `
		UPDATE patient 
		SET deleted_at = NOW() 
		WHERE id = (SELECT account_id FROM account WHERE id = $1)
		RETURNING name
	`
	var updatedName string
	err = tx.QueryRow(ctx, query, accountID).Scan(&updatedName)
	if err != nil {
		return "", fmt.Errorf("failed to soft delete user: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("User %s soft deleted successfully", updatedName), nil
}
