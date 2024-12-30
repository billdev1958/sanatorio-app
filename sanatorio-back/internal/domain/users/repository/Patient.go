package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/users/entities"
)

// Función para registrar un usuario
func (ur *userRepository) RegisterPatientTransaction(ctx context.Context, account entities.Account, pu entities.PatientUser) (entities.PatientUser, error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return pu, fmt.Errorf("begin transaction: %w", err)
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
		return pu, fmt.Errorf("failed to register account: %w", err)
	}

	err = ur.registerMedicalHistory(ctxTx, tx, pu.MedicalHistoryID, pu)
	if err != nil {
		return pu, fmt.Errorf("failed to register medical history: %w", err)
	}

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerPatient(ctxTx, tx, accountID, pu)
	if err != nil {
		return pu, fmt.Errorf("failed to register patient: %w", err)
	}

	return pu, nil
}

func (ur *userRepository) UpdatePatient(ctx context.Context, d entities.PatientUser) (message string, err error) {
	// Consulta SQL con COALESCE
	query := `
		UPDATE patient
		SET 
			first_name = COALESCE($1, first_name),
			last_name1 = COALESCE($2, last_name1),
			last_name2 = COALESCE($3, last_name2),
			curp = COALESCE($4, CURP),
			sex = COALESCE($5, sex),
			updated_at = CURRENT_TIMESTAMP
		WHERE account_id = $6
	`

	log.Printf("Starting update for patient with account_id: %s", d.AccountID)

	tag, err := ur.storage.DbPool.Exec(ctx, query,
		d.FirstName,
		d.LastName1,
		d.LastName2,
		d.Curp,
		d.Sex,
		d.AccountID,
	)

	if err != nil {
		log.Printf("Error updating patient with account_id: %s. Error: %v", d.AccountID, err)
		return "", fmt.Errorf("failed to update patient: %w", err)
	}

	if tag.RowsAffected() == 0 {
		log.Printf("No rows updated for patient with account_id: %s", d.AccountID)
		return "No patient record updated", nil
	}

	log.Printf("Successfully updated patient with account_id: %s. Rows affected: %d", d.AccountID, tag.RowsAffected())

	return "patient updated successfully", nil
}

func (ur *userRepository) SoftDeleteUserPatient(ctx context.Context, a entities.Account) (bool, error) {

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
    UPDATE patient 
    SET deleted_at = NOW() 
    WHERE account_id = $1
    `

	_, err = tx.Exec(ctx, query, a.ID)
	if err != nil {
		return false, err // Retorna false y el error si la ejecución falla
	}

	return true, nil // Retorna true y nil si todo sale bien
}

func (ur *userRepository) RegisterBeneficiary(ctx context.Context, bu entities.BeneficiaryUser) (message string, err error) {
	// Crear un contexto con cancelación
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Iniciar la transacción
	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return "", fmt.Errorf("begin transaction: %w", err)
	}

	// Definir el defer para manejo de commit/rollback
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

	err = ur.registerMedicalHistoryB(ctxTx, tx, bu.MedicalHistoryID, bu)
	if err != nil {
		return "", fmt.Errorf("failed to register medical history: %w", err)
	}

	query := "INSERT INTO beneficiary (id, account_holder, medical_history_id, first_name, last_name1, last_name2, curp, sex) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err = tx.Exec(ctxTx, query, bu.ID, bu.AccountHolder, bu.MedicalHistoryID, bu.Firstname, bu.Lastname1, bu.Lastname2, bu.Curp, bu.Sex)
	if err != nil {
		return "", fmt.Errorf("error al insertar en la tabla beneficiary: %w", err)
	}

	return "Beneficiario registrado exitosamente", nil
}
