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

	return du, nil
}

func (pr *userRepository) registerDoctor(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, du entities.DoctorUser) error {
	du.AccountID = accountID

	if du.AccountID == uuid.Nil || du.MedicalLicense == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	query := "INSERT INTO doctor (account_id, first_name, last_name1, last_name2, specialty_license, medical_license, sex) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := tx.Exec(ctx, query, accountID, du.FirstName, du.LastName1, du.LastName2, du.SpecialtyLicense, du.MedicalLicense, du.Sex)
	if err != nil {
		return fmt.Errorf("insert into doctor table: %w", err)
	}

	return nil
}

func (ur *userRepository) UpdateDoctor(ctx context.Context, d entities.DoctorUser) (message string, err error) {
	// Consulta SQL con COALESCE
	query := `
		UPDATE doctor
		SET 
			first_name = COALESCE($1, first_name),
			last_name1 = COALESCE($2, last_name1),
			last_name2 = COALESCE($3, last_name2),
			specialty_license = COALESCE($4, specialty_license),
			medical_license = COALESCE($5, medical_license),
			sex = COALESCE($6, sex),
			updated_at = CURRENT_TIMESTAMP
		WHERE account_id = $7
	`

	log.Printf("Starting update for doctor with account_id: %s", d.AccountID)

	tag, err := ur.storage.DbPool.Exec(ctx, query,
		d.FirstName,
		d.LastName1,
		d.LastName2,
		d.SpecialtyLicense,
		d.MedicalLicense,
		d.Sex,
		d.AccountID,
	)

	if err != nil {
		log.Printf("Error updating doctor with account_id: %s. Error: %v", d.AccountID, err)
		return "", fmt.Errorf("failed to update doctor: %w", err)
	}

	if tag.RowsAffected() == 0 {
		log.Printf("No rows updated for doctor with account_id: %s", d.AccountID)
		return "No doctor record updated", nil
	}

	log.Printf("Successfully updated doctor with account_id: %s. Rows affected: %d", d.AccountID, tag.RowsAffected())

	return "Doctor updated successfully", nil
}
