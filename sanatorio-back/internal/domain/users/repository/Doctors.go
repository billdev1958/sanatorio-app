package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) RegisterDoctorTransaction(ctx context.Context, account entities.Account, du entities.DoctorUser) (entities.DoctorUser, error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return du, fmt.Errorf("begin transaction: %w", err)
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
		return du, fmt.Errorf("failed to register account: %w", err)
	}

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerDoctor(ctxTx, tx, accountID, du)
	if err != nil {
		return du, fmt.Errorf("failed to register doctor: %w", err)
	}

	// registrar la cuenta de doctor en tabla user_roles
	queryUserRole := `
		INSERT INTO user_roles (account_id, role_id) 
		VALUES ($1, $2)`
	_, err = tx.Exec(ctxTx, queryUserRole, accountID, entities.Doctor)
	if err != nil {
		return du, fmt.Errorf("failed to assign doctor role in user_roles: %w", err)
	}

	return du, nil
}

func (pr *userRepository) registerDoctor(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, du entities.DoctorUser) error {
	du.AccountID = accountID

	// Verificar que los campos obligatorios estén presentes
	if du.AccountID == uuid.Nil || du.MedicalLicense == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	// Preparar la consulta para insertar el tipo de doctor
	query := "INSERT INTO doctor (account_id, first_name, last_name1, last_name2, specialty_license, medical_license, sex) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, accountID, du.FirstName, du.LastName1, du.LastName2, du.SpecialtyLicense, du.MedicalLicense, du.Sex)
	if err != nil {
		return fmt.Errorf("insert into doctor table: %w", err)
	}

	return nil
}
