package repository

import (
	"context"
	"fmt"
	"log"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	postgres "sanatorioApp/internal/infraestructure/db"
	"time"

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
func (ur *userRepository) RegisterPatientTransaction(ctx context.Context, pu entities.PatientUser) (entities.PatientUser, error) {
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
	/*err = ur.registerAccount(ctxTx, tx, pu.Account, userID)
	if err != nil {
		return pu, fmt.Errorf("failed to register account: %w", err)
	}*/

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerPatient(ctxTx, tx, pu)
	if err != nil {
		return pu, fmt.Errorf("failed to register patient: %w", err)
	}

	return pu, nil
}

func (pr *userRepository) registerPatient(ctx context.Context, tx pgx.Tx, pt entities.PatientUser) error {
	// Verificar que los campos obligatorios estén presentes
	if pt.AccountID == uuid.Nil || pt.Curp == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	if pt.Rol != entities.Patient {
		return fmt.Errorf("invalid rol: required type patient")

	}

	// Preparar la consulta para insertar el tipo de doctor
	query := "INSERT INTO patient_user (account_id, curp, created_at) VALUES ($1, $2, $3, $4)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, pt.AccountID, pt.Curp, time.Now())
	if err != nil {
		return fmt.Errorf("insert into patient table: %w", err)
	}

	return nil
}

// Función para registrar la cuenta utilizando el userID generado
func (pr *userRepository) registerAccount(ctx context.Context, tx pgx.Tx, ru entities.Account, userID int) error {
	var email string
	query := "INSERT INTO account (id, user_id, email, password, rol, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING email"
	err := tx.QueryRow(ctx, query, ru.AccountID, userID, ru.Email, ru.Password, ru.Rol, time.Now()).Scan(&email)
	if err != nil {
		return fmt.Errorf("insert account: %w", err)
	}
	return nil
}
