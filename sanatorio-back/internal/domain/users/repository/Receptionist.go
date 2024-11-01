package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) RegisterReceptionistTransaction(ctx context.Context, account entities.Account, ru entities.ReceptionistUser) (entities.ReceptionistUser, error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return ru, fmt.Errorf("begin transaction: %w", err)
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
		return ru, fmt.Errorf("failed to register account: %w", err)
	}

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerReceptionist(ctxTx, tx, accountID, ru)
	if err != nil {
		return ru, fmt.Errorf("failed to register receptionist: %w", err)
	}

	// registrar la cuenta de doctor en tabla user_roles
	queryUserRole := `
		INSERT INTO user_roles (account_id, role_id) 
		VALUES ($1, $2)`
	_, err = tx.Exec(ctxTx, queryUserRole, accountID, entities.Receptionist)
	if err != nil {
		return ru, fmt.Errorf("failed to assign receptionist role in user_roles: %w", err)
	}

	return ru, nil
}

func (pr *userRepository) registerReceptionist(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, ru entities.ReceptionistUser) error {
	ru.AccountID = accountID

	// Verificar que los campos obligatorios estén presentes
	if ru.AccountID == uuid.Nil || ru.Curp == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	query := "INSERT INTO receptionist (account_id, first_name, last_name1, last_name2, curp, sex) VALUES ($1, $2, $3, $4, $5, $6)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, accountID, ru.FirstName, ru.LastName1, ru.LastName2, ru.Curp, ru.Sex)
	if err != nil {
		return fmt.Errorf("insert into receptionist table: %w", err)
	}

	return nil
}
