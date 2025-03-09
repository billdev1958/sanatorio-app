package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (ur *userRepository) SaveCodeVerification(ctx context.Context, email, code string, expired_at time.Time) error {
	query := `
		INSERT INTO verification_codes (email, code, expires_at) VALUES ($1, $2, $3)
	`
	_, err := ur.storage.DbPool.Exec(ctx, query, &email, &code, &expired_at)
	if err != nil {
		return fmt.Errorf("error updating account verification: %w", err)
	}

	return nil

}

func (ur *userRepository) VerifyCode(ctx context.Context, code, email string) (bool, error) {

	query1 := `
		SELECT id FROM verification_codes
		WHERE email = $1
		AND code = $2
		AND used = false
		AND expires_at > NOW();
	`

	query2 := `
		UPDATE account
		SET is_verified = true
		WHERE email = $1
	`

	query3 := `
		DELETE FROM verification_codes
		WHERE id = $1
	`

	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("transaction for verify code failed: %w", err)
	}

	defer tx.Rollback(ctx)

	var idVerifyCode int

	err = tx.QueryRow(ctx, query1, email, code).Scan(&idVerifyCode)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("code not found or expired")
		}
		return false, fmt.Errorf("database error: %w", err)
	}
	_, err = tx.Exec(ctx, query2, email)
	if err != nil {
		return false, fmt.Errorf("activate account failed: %w", err)
	}

	_, err = tx.Exec(ctx, query3, idVerifyCode)
	if err != nil {
		return false, fmt.Errorf("failed to delete verification code: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return false, fmt.Errorf("transaction failed: %w", err)
	}

	return true, nil
}
