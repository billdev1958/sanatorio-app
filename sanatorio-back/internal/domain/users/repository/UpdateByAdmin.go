package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"
	"strings"

	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) UpdateUser(ctx context.Context, userUpdate entities.UpdateUser) (string, error) {
	// Iniciar la transacción
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx) // Asegura que la transacción se revierta si no se confirma

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctx, tx, userUpdate.AccountAdminID, userUpdate.AdminPassword)
	if err != nil {
		return "", fmt.Errorf("failed to authenticate admin: %w", err)
	}
	if !isValid {
		return "", fmt.Errorf("authentication failed: invalid credentials")
	}

	// Verificar que el usuario existe a partir del account_id
	var existingUserID int
	err = tx.QueryRow(ctx, "SELECT user_id FROM account WHERE id = $1", userUpdate.AccountID).Scan(&existingUserID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("user with account ID %s not found", userUpdate.AccountID)
		}
		return "", fmt.Errorf("failed to check if user exists: %w", err)
	}

	// Actualización condicional de los campos en la tabla 'users'
	var setClauses []string
	var args []interface{}
	argIndex := 1

	if userUpdate.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, userUpdate.Name)
		argIndex++
	}
	if userUpdate.Lastname1 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname1 = $%d", argIndex))
		args = append(args, userUpdate.Lastname1)
		argIndex++
	}
	if userUpdate.Lastname2 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname2 = $%d", argIndex))
		args = append(args, userUpdate.Lastname2)
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
	if userUpdate.Password != "" {
		accountSetClauses := []string{"updated_at = NOW()"}
		accountArgs := []interface{}{}
		accountIndex := 1

		if userUpdate.Password != "" {
			accountSetClauses = append(accountSetClauses, fmt.Sprintf("password = $%d", accountIndex))
			accountArgs = append(accountArgs, userUpdate.Password)
			accountIndex++
		}

		accountQuery := fmt.Sprintf("UPDATE account SET %s WHERE id = $%d", strings.Join(accountSetClauses, ", "), accountIndex)
		accountArgs = append(accountArgs, userUpdate.AccountID)

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

func (ur *userRepository) UpdateDoctor(ctx context.Context, doctorUpdate entities.UpdateDoctor) (string, error) {
	// Iniciar la transacción
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctx, tx, doctorUpdate.AccountAdminID, doctorUpdate.AdminPassword)
	if err != nil {
		return "", fmt.Errorf("failed to authenticate admin: %w", err)
	}
	if !isValid {
		return "", fmt.Errorf("authentication failed: invalid credentials")
	}

	var setClauses []string
	var args []interface{}
	argIndex := 1

	if doctorUpdate.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, doctorUpdate.Name)
		argIndex++
	}
	if doctorUpdate.Lastname1 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname1 = $%d", argIndex))
		args = append(args, doctorUpdate.Lastname1)
		argIndex++
	}
	if doctorUpdate.Lastname2 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname2 = $%d", argIndex))
		args = append(args, doctorUpdate.Lastname2)
		argIndex++
	}

	setClauses = append(setClauses, "updated_at = NOW()")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = (SELECT user_id FROM account WHERE id = $%d) RETURNING name",
		strings.Join(setClauses, ", "), argIndex)
	args = append(args, doctorUpdate.AccountID)

	var updatedName string
	err = tx.QueryRow(ctx, query, args...).Scan(&updatedName)
	if err != nil {
		return "", fmt.Errorf("failed to update doctor user: %w", err)
	}

	if doctorUpdate.Password != "" || doctorUpdate.MedicalLicense != "" || doctorUpdate.SpecialtyID != 0 {
		accountSetClauses := []string{"updated_at = NOW()"}
		accountArgs := []interface{}{}
		accountIndex := 1

		if doctorUpdate.Password != "" {
			accountSetClauses = append(accountSetClauses, fmt.Sprintf("password = $%d", accountIndex))
			accountArgs = append(accountArgs, doctorUpdate.Password)
			accountIndex++
		}

		accountQuery := fmt.Sprintf("UPDATE account SET %s WHERE id = $%d", strings.Join(accountSetClauses, ", "), accountIndex)
		accountArgs = append(accountArgs, doctorUpdate.AccountID)

		_, err := tx.Exec(ctx, accountQuery, accountArgs...)
		if err != nil {
			return "", fmt.Errorf("failed to update account: %w", err)
		}

		doctorSetClauses := []string{"updated_at = NOW()"}
		doctorArgs := []interface{}{}
		doctorIndex := 1

		if doctorUpdate.MedicalLicense != "" {
			doctorSetClauses = append(doctorSetClauses, fmt.Sprintf("medical_license = $%d", doctorIndex))
			doctorArgs = append(doctorArgs, doctorUpdate.MedicalLicense)
			doctorIndex++
		}
		if doctorUpdate.SpecialtyID != 0 {
			doctorSetClauses = append(doctorSetClauses, fmt.Sprintf("id_specialty = $%d", doctorIndex))
			doctorArgs = append(doctorArgs, doctorUpdate.SpecialtyID)
			doctorIndex++
		}

		doctorQuery := fmt.Sprintf("UPDATE doctor_user SET %s WHERE account_id = $%d",
			strings.Join(doctorSetClauses, ", "), doctorIndex)
		doctorArgs = append(doctorArgs, doctorUpdate.AccountID)

		_, err = tx.Exec(ctx, doctorQuery, doctorArgs...)
		if err != nil {
			return "", fmt.Errorf("failed to update doctor_user: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("Doctor %s updated successfully", updatedName), nil
}
