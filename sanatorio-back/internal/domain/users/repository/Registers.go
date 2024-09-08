package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/users/entities"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

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

	// Registrar el usuario y obtener el userID generado
	userID, err := ur.registerUser(ctxTx, tx, pu.User)
	if err != nil {
		return pu, fmt.Errorf("failed to register user: %w", err)
	}

	// Registrar la cuenta utilizando el userID generado
	err = ur.registerAccount(ctxTx, tx, pu.Account, userID)
	if err != nil {
		return pu, fmt.Errorf("failed to register account: %w", err)
	}

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerPatient(ctxTx, tx, pu)
	if err != nil {
		return pu, fmt.Errorf("failed to register patient: %w", err)
	}

	return pu, nil
}

// Funciones auxiliares para registrar usuarios
func (pr *userRepository) registerSuperAdmin(ctx context.Context, tx pgx.Tx, su entities.SuperUser) (entities.SuperUser, error) {
	// Verificar que los campos obligatorios estén presentes
	if su.AccountID == uuid.Nil || su.Curp == "" {
		return su, fmt.Errorf("invalid input: missing required fields")
	}

	if su.Account.Rol != entities.SuperUsuario {
		return su, fmt.Errorf("invalid rol: required type super_usuario")
	}

	// Preparar la consulta para insertar el tipo de doctor
	query := "INSERT INTO super_user (account_id, curp, created_at) VALUES ($1, $2, $3)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, su.AccountID, su.Curp, time.Now())
	if err != nil {
		return su, fmt.Errorf("insert into super_user table: %w", err)
	}

	return su, nil
}

func (pr *userRepository) registerDoctor(ctx context.Context, tx pgx.Tx, du entities.DoctorUser) error {
	// Verificar que los campos obligatorios estén presentes
	if du.AccountID == uuid.Nil || du.SpecialtyID == 0 || du.MedicalLicense == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	// Preparar la consulta para insertar el tipo de doctor
	query := "INSERT INTO doctor_user (account_id, specialty_id, medical_license, created_at) VALUES ($1, $2, $3, $4)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, du.AccountID, du.SpecialtyID, du.MedicalLicense, time.Now())
	if err != nil {
		return fmt.Errorf("insert into doctor_user table: %w", err)
	}

	return nil
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

func (pr *userRepository) registerUser(ctx context.Context, tx pgx.Tx, ru entities.User) (int, error) {
	var userID int
	query := "INSERT INTO users (name, lastname1, lastname2, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := tx.QueryRow(ctx, query, ru.Name, ru.Lastname1, ru.Lastname2, time.Now()).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("insert user: %w", err)
	}
	return userID, nil
}
