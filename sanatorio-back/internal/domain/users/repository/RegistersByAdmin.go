package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/users/entities"

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
			err = tx.Commit(ctxTx)
			if err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctxTx, tx, rd.AccountAdminID, rd.AdminPassword)
	if err != nil || !isValid {
		return response, fmt.Errorf("failed to authenticate admin: %w", err)
	}

	// Registrar el usuario y obtener el userID generado
	userID, name, err := ur.RegisterUser(ctxTx, tx, entities.RegisterUserByAdmin{
		User:       rd.User,
		Account:    rd.Account,
		DocumentID: rd.DocumentID,
	})
	if err != nil {
		return response, err
	}

	// Registrar la cuenta utilizando el userID generado
	email, err := ur.RegisterAccount(ctxTx, tx, entities.RegisterUserByAdmin{
		User:       rd.User,
		Account:    rd.Account,
		DocumentID: rd.DocumentID,
	}, userID)
	if err != nil {
		return response, err
	}

	// Registrar al doctor en su tabla correspondiente
	err = ur.RegisterTypeDoctor(ctxTx, tx, rd)
	if err != nil {
		return response, err
	}

	// Construir la respuesta
	response = entities.UserResponse{
		Name:  name,
		Email: email,
	}

	return response, nil
}

// Función para registrar el doctor en su tabla correspondiente
func (pr *userRepository) RegisterTypeDoctor(ctx context.Context, tx pgx.Tx, rd entities.RegisterDoctorByAdmin) error {
	query := "INSERT INTO doctor_user (user_id, account_id, specialty_id, medical_license) VALUES ($1, $2, $3, $4)"
	_, err := tx.Exec(ctx, query, rd.AccountID, rd.SpecialtyID, rd.DocumentID)
	if err != nil {
		return fmt.Errorf("insert into table: %w", err)
	}
	return nil
}

// Función para registrar el tipo de usuario en la tabla correspondiente
func (pr *userRepository) RegisterTypeUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUserByAdmin) error {
	var query string
	var values []interface{}

	if ru.Rol == entities.SuperUsuario {
		query = "INSERT INTO super_user (account_id, curp, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())"
		values = []interface{}{ru.AccountID, ru.DocumentID}
	} else if ru.Rol == entities.Patient {
		query = "INSERT INTO patient_user (account_id, curp, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())"
		values = []interface{}{ru.AccountID, ru.DocumentID}
	} else {
		return fmt.Errorf("unknown role: %d", ru.Rol)
	}

	_, err := tx.Exec(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("insert into table: %w", err)
	}

	return nil
}
