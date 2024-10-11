package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"
	"strings"

	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) UpdatePatient(ctx context.Context, patientUpdate entities.PatientUser) (string, error) {
	// Iniciar la transacción
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx) // Asegura que la transacción se revierta si no se confirma

	// Verificar que el usuario existe a partir del account_id
	var existingUserID int
	err = tx.QueryRow(ctx, "SELECT user_id FROM account WHERE id = $1", patientUpdate.AccountID).Scan(&existingUserID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("user with account ID %s not found", patientUpdate.AccountID)
		}
		return "", fmt.Errorf("failed to check if user exists: %w", err)
	}

	// Actualización condicional de los campos en la tabla 'users'
	var setClauses []string
	var args []interface{}
	argIndex := 1

	if patientUpdate.FirstName != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, patientUpdate.FirstName)
		argIndex++
	}
	if patientUpdate.LastName1 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname1 = $%d", argIndex))
		args = append(args, patientUpdate.LastName1)
		argIndex++
	}
	if patientUpdate.LastName2 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname2 = $%d", argIndex))
		args = append(args, patientUpdate.LastName2)
		argIndex++
	}

	setClauses = append(setClauses, "updated_at = NOW()")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d RETURNING name", strings.Join(setClauses, ", "), argIndex)
	args = append(args, existingUserID)

	var updatedName string
	err = tx.QueryRow(ctx, query, args...).Scan(&updatedName)
	if err != nil {
		return "", fmt.Errorf("failed to update user: %w", err)
	}

	// Actualizar la cuenta si es necesario
	if patientUpdate.Password != "" {
		accountSetClauses := []string{"updated_at = NOW()"}
		accountArgs := []interface{}{}
		accountIndex := 1

		if patientUpdate.Password != "" {
			accountSetClauses = append(accountSetClauses, fmt.Sprintf("password = $%d", accountIndex))
			accountArgs = append(accountArgs, patientUpdate.Password)
			accountIndex++
		}

		accountQuery := fmt.Sprintf("UPDATE account SET %s WHERE id = $%d", strings.Join(accountSetClauses, ", "), accountIndex)
		accountArgs = append(accountArgs, patientUpdate.AccountID)

		_, err := tx.Exec(ctx, accountQuery, accountArgs...)
		if err != nil {
			return "", fmt.Errorf("failed to update account: %w", err)
		}
	}

	// Confirmar la transacción
	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("User %s updated successfully", updatedName), nil
}
