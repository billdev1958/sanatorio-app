package repository

import (
	"context"
	"fmt"
	"log"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	storage *postgres.PgxStorage
}

func NewUserRepository(storage *postgres.PgxStorage) user.Repository {
	return &userRepository{storage: storage}
}

// TODO corregir registro
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

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerPatient(ctxTx, tx, accountID, pu)
	if err != nil {
		return pu, fmt.Errorf("failed to register patient: %w", err)
	}

	return pu, nil
}

func (pr *userRepository) registerPatient(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, pt entities.PatientUser) error {
	// Verificar que los campos obligatorios estén presentes
	if pt.AccountID == uuid.Nil || pt.Curp == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	// Preparar la consulta para insertar el tipo de doctor
	query := "INSERT INTO patient_user (account_id, medical_history_id, firstname, lastname1, lastname2, curp) VALUES ($1, $2, $3, $4, $5, $6)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, accountID, pt.MedicalHistoryID, pt.FirstName, pt.LastName1, pt.LastName2, pt.Curp)
	if err != nil {
		return fmt.Errorf("insert into patient table: %w", err)
	}

	return nil
}

// Función para registrar la cuenta utilizando el userID generado
func (pr *userRepository) registerAccount(ctx context.Context, tx pgx.Tx, ru entities.Account) (uuid.UUID, error) {
	var accountID uuid.UUID
	query := "INSERT INTO account (id, affiliation_id, phone, email, password, rol, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := tx.QueryRow(ctx, query, ru.AfiliationID, ru.ID, ru.PhoneNumber, ru.Email, ru.Password, ru.Rol).Scan(&accountID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert account: %w", err)
	}
	return accountID, nil
}
