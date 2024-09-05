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

func (ur *userRepository) RegisterUserTransaction(ctx context.Context, ru entities.RegisterUserByAdmin) (response entities.UserResponse, err error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return response, fmt.Errorf("begin transaction: %w", err)
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

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctxTx, tx, ru.AccountAdminID, ru.AdminPassword)
	if err != nil {
		return response, fmt.Errorf("failed to authenticate admin: %w", err)
	}
	if !isValid {
		return response, fmt.Errorf("failed to authenticate admin: invalid password")
	}

	// Registrar el usuario y obtener el userID generado
	userID, name, err := ur.RegisterUser(ctxTx, tx, ru)
	if err != nil {
		return response, fmt.Errorf("failed to register user: %w", err)
	}

	// Registrar la cuenta utilizando el userID generado
	email, err := ur.RegisterAccount(ctxTx, tx, ru, userID)
	if err != nil {
		return response, fmt.Errorf("failed to register account: %w", err)
	}

	// Manejar el tipo de usuario según el rol
	switch ru.Rol {
	case entities.SuperUsuario, entities.Patient:
		err = ur.RegisterTypeUser(ctxTx, tx, ru)
		if err != nil {
			return response, fmt.Errorf("failed to register user type: %w", err)
		}
	default:
		return response, fmt.Errorf("unknown role: %d", ru.Rol)
	}

	// Construir la respuesta
	response = entities.UserResponse{
		Name:  name,
		Email: email,
	}

	return response, nil
}

func (ur *userRepository) RegisterDoctorTransaction(ctx context.Context, rd entities.RegisterDoctorByAdmin) (response entities.UserResponse, err error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return response, fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctxTx); rbErr != nil {
				log.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			if err = tx.Commit(ctxTx); err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctxTx, tx, rd.AccountAdminID, rd.AdminPassword)
	if err != nil {
		return response, fmt.Errorf("failed to authenticate admin: %w", err)
	}
	if !isValid {
		return response, fmt.Errorf("authentication failed: invalid credentials")
	}

	// Registrar el usuario y obtener el userID generado
	userID, name, err := ur.RegisterdDoctor(ctxTx, tx, rd)
	if err != nil {
		return response, fmt.Errorf("failed to register doctor user: %w", err)
	}

	// Registrar la cuenta utilizando el userID generado
	email, err := ur.RegisterAccountDoctor(ctxTx, tx, rd, userID)
	if err != nil {
		return response, fmt.Errorf("failed to register account: %w", err)
	}

	err = ur.RegisterTypeDoctor(ctxTx, tx, rd)
	if err != nil {
		return response, fmt.Errorf("failed to register user type: %w", err)
	}

	// Construir la respuesta
	response = entities.UserResponse{
		Name:  name,
		Email: email,
	}

	return response, nil
}

func (pr *userRepository) RegisterTypeDoctor(ctx context.Context, tx pgx.Tx, rd entities.RegisterDoctorByAdmin) error {
	// Verificar que los campos obligatorios estén presentes
	if rd.AccountID == uuid.Nil || rd.SpecialtyID == 0 || rd.DocumentID == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	// Preparar la consulta para insertar el tipo de doctor
	query := "INSERT INTO doctor_user (account_id, specialty_id, medical_license, created_at) VALUES ($1, $2, $3, $4)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, rd.AccountID, rd.SpecialtyID, rd.DocumentID, time.Now())
	if err != nil {
		return fmt.Errorf("insert into doctor_user table: %w", err)
	}

	return nil
}

// Función para registrar el tipo de usuario en la tabla correspondiente
func (pr *userRepository) RegisterTypeUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUserByAdmin) error {
	var query string
	var values []interface{}

	if ru.Rol == entities.SuperUsuario {
		query = "INSERT INTO super_user (account_id, curp, created_at) VALUES ($1, $2, $3)"
		values = []interface{}{ru.AccountID, ru.DocumentID, time.Now()}
	} else if ru.Rol == entities.Patient {
		query = "INSERT INTO patient_user (account_id, curp, created_at) VALUES ($1, $2, $3)"
		values = []interface{}{ru.AccountID, ru.DocumentID, time.Now()}
	} else {
		return fmt.Errorf("unknown role: %d", ru.Rol)
	}

	// Depuración antes de ejecutar la consulta
	log.Printf("Executing query: %s with values %v", query, values)

	_, err := tx.Exec(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("insert into table: %w", err)
	}

	return nil
}
