package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) RegisterAdminTransaction(ctx context.Context, account entities.Account, su entities.SuperAdminUser) (entities.SuperAdminUser, error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return su, fmt.Errorf("begin transaction: %w", err)
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
		return su, fmt.Errorf("failed to register account: %w", err)
	}

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerAdmin(ctxTx, tx, accountID, su)
	if err != nil {
		return su, fmt.Errorf("failed to register super_admin: %w", err)
	}

	// registrar la cuenta de doctor en tabla user_roles
	queryUserRole := `
		INSERT INTO user_roles (account_id, role_id) 
		VALUES ($1, $2)`
	_, err = tx.Exec(ctxTx, queryUserRole, accountID, entities.SuperAdmin)
	if err != nil {
		return su, fmt.Errorf("failed to assign super_admin role in user_roles: %w", err)
	}

	return su, nil
}

func (pr *userRepository) registerAdmin(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, su entities.SuperAdminUser) error {
	su.AccountID = accountID

	// Verificar que los campos obligatorios estén presentes
	if su.AccountID == uuid.Nil || su.Curp == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	// Preparar la consulta para insertar el tipo de doctor
	query := "INSERT INTO super_admin (account_id, first_name, last_name1, last_name2, curp, sex) VALUES ($1, $2, $3, $4, $5, $6)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, accountID, su.FirstName, su.LastName1, su.LastName2, su.Curp, su.Sex)
	if err != nil {
		return fmt.Errorf("insert into super_admin table: %w", err)
	}

	return nil
}
