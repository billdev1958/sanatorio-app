package repository

import (
	"context"
	"fmt"
	"log"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	storage *postgres.PgxStorage
}

func NewUserRepository(storage *postgres.PgxStorage) user.Repository {
	return &userRepository{storage: storage}
}

func (ur *userRepository) RegisterUserTransaction(ctx context.Context, ru entities.RegisterUser) (response entities.UserResponse, err error) {

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

	// Registrar el usuario
	userID, name, err := ur.RegisterUser(ctxTx, tx, ru)
	if err != nil {
		return response, err
	}

	// Registrar la cuenta
	ru.User.ID = userID // Asigna el ID del usuario al campo UserID en la cuenta
	email, err := ur.RegisterAccount(ctxTx, tx, ru)
	if err != nil {
		return response, err
	}

	// Manejar el tipo de usuario según el rol
	if ru.Rol == entities.SuperUsuario {
		err = ur.RegisterTypeUser(ctxTx, tx, ru)
		if err != nil {
			return response, err
		}
	} else if ru.Rol == entities.Patient {
		err = ur.RegisterTypeUser(ctxTx, tx, ru)
		if err != nil {
			return response, err
		}
	} else {
		return response, fmt.Errorf("unknown role: %d", ru.Rol)
	}

	// Construir la respuesta
	response = entities.UserResponse{
		Name:  name,
		Email: email,
	}

	return response, nil
}

func (ur *userRepository) RegisterDoctorTransaction(ctx context.Context, rd entities.RegisterDoctor) (response entities.UserResponse, err error) {

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

	// Registrar el usuario como Doctor
	err = ur.RegisterTypeDoctor(ctxTx, tx, rd)
	if err != nil {
		return response, err
	}

	// Construir la respuesta
	response = entities.UserResponse{
		Name:  rd.Name,
		Email: rd.Email,
	}

	return response, nil
}

func (pr *userRepository) RegisterUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (userID int, name string, err error) {
	query := "INSERT INTO users (name, lastname1, lastname2, created_at) VALUES ($1, $2, $3, $4) RETURNING id, name"
	err = tx.QueryRow(ctx, query, ru.Name, ru.Lastname1, ru.Lastname2).Scan(&userID, &name)
	if err != nil {
		return 0, "", fmt.Errorf("insert user: %w", err)
	}
	return userID, name, nil
}

func (pr *userRepository) RegisterAccount(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (email string, err error) {
	query := "INSERT INTO account (user_id, email, password, rol, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING email"
	err = tx.QueryRow(ctx, query, ru.AccountID, ru.Email, ru.Password, ru.Rol).Scan(&email)
	if err != nil {
		return "", fmt.Errorf("insert account: %w", err)
	}
	return email, nil
}

func (pr *userRepository) RegisterTypeUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) error {
	var query string
	var values []interface{}

	if ru.Rol == entities.SuperUsuario {
		query = "INSERT INTO super_user (user_id, account_id, curp) VALUES ($1, $2, $3)"
		values = []interface{}{ru.UserID, ru.AccountID, ru.DocumentID}
	} else if ru.Rol == entities.Patient {
		query = "INSERT INTO patient_user (user_id, account_id, curp) VALUES ($1, $2, $3)"
		values = []interface{}{ru.UserID, ru.AccountID, ru.DocumentID}
	} else {
		return fmt.Errorf("unknown role: %d", ru.Rol)
	}

	_, err := tx.Exec(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("insert into table: %w", err)
	}

	return nil
}

func (pr *userRepository) RegisterTypeDoctor(ctx context.Context, tx pgx.Tx, rd entities.RegisterDoctor) error {
	query := "INSERT INTO doctor_user (user_id, account_id, specialty_id, medical_license) VALUES ($1, $2, $3, $4)"
	_, err := tx.Exec(ctx, query, rd.UserID, rd.AccountID, rd.SpecialtyID, rd.DocumentID)
	if err != nil {
		return fmt.Errorf("insert into table: %w", err)
	}
	return nil
}
