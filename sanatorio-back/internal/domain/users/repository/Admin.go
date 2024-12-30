package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) RegisterAdminTransaction(ctx context.Context, account entities.Account, au entities.AdminUser) (entities.AdminUser, error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return au, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctxTx); rbErr != nil {
				log.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			err = tx.Commit(ctxTx)
			if err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	// Registrar la cuenta utilizando el userID generado
	accountID, err := ur.registerAccount(ctxTx, tx, account)
	if err != nil {
		return au, fmt.Errorf("failed to register account: %w", err)
	}

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerAdmin(ctxTx, tx, accountID, au)
	if err != nil {
		return au, fmt.Errorf("failed to register admin: %w", err)
	}

	return au, nil
}

func (pr *userRepository) registerAdmin(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, au entities.AdminUser) error {
	au.AccountID = accountID

	if au.AccountID == uuid.Nil || au.Curp == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	query := "INSERT INTO admin (account_id, first_name, last_name1, last_name2, curp, sex) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := tx.Exec(ctx, query, accountID, au.FirstName, au.LastName1, au.LastName2, au.Curp, au.Sex)
	if err != nil {
		return fmt.Errorf("insert into admin table: %w", err)
	}

	return nil
}

func (ur *userRepository) UpdateAdmin(ctx context.Context, d entities.AdminUser) (message string, err error) {
	// Consulta SQL con COALESCE
	query := `
		UPDATE admin
		SET 
			first_name = COALESCE($1, first_name),
			last_name1 = COALESCE($2, last_name1),
			last_name2 = COALESCE($3, last_name2),
			curp = COALESCE($4, CURP),
			sex = COALESCE($5, sex),
			updated_at = CURRENT_TIMESTAMP
		WHERE account_id = $6
	`

	log.Printf("Starting update for admin with account_id: %s", d.AccountID)

	tag, err := ur.storage.DbPool.Exec(ctx, query,
		d.FirstName,
		d.LastName1,
		d.LastName2,
		d.Curp,
		d.Sex,
		d.AccountID,
	)

	if err != nil {
		log.Printf("Error updating admin with account_id: %s. Error: %v", d.AccountID, err)
		return "", fmt.Errorf("failed to update admin: %w", err)
	}

	if tag.RowsAffected() == 0 {
		log.Printf("No rows updated for admin with account_id: %s", d.AccountID)
		return "No admin record updated", nil
	}

	log.Printf("Successfully updated admin with account_id: %s. Rows affected: %d", d.AccountID, tag.RowsAffected())

	return "admin updated successfully", nil
}

func (ur *userRepository) SoftDeleteUserAdmin(ctx context.Context, a entities.Account) (bool, error) {

	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return false, err // Retorna false y el error
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctxTx); rbErr != nil {
				log.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			err = tx.Commit(ctxTx)
			if err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	var exists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM account WHERE id = $1)"
	err = tx.QueryRow(ctx, checkQuery, a.ID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence in account: %w", err)
	}

	if !exists {
		return false, fmt.Errorf("user with account ID %s not found", a.ID)
	}

	query := `
    UPDATE admin 
    SET deleted_at = NOW() 
    WHERE account_id = $1
    `

	_, err = tx.Exec(ctx, query, a.ID)
	if err != nil {
		return false, err // Retorna false y el error si la ejecución falla
	}

	return true, nil // Retorna true y nil si todo sale bien
}
