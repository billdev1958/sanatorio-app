package repository

import (
	"context"
	"fmt"
	"log"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	postgres "sanatorioApp/internal/infraestructure/db"
	"time"

	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	storage *postgres.PgxStorage
}

func NewUserRepository(storage *postgres.PgxStorage) user.Repository {
	return &userRepository{storage: storage}
}

// func (pr userRepository) RegisterPatient()

// Función para registrar un usuario y devolver el userID generado
func (pr *userRepository) RegisterUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUserByAdmin) (int, string, error) {
	var userID int
	var name string
	query := "INSERT INTO users (name, lastname1, lastname2, created_at) VALUES ($1, $2, $3, $4) RETURNING id, name"
	err := tx.QueryRow(ctx, query, ru.Name, ru.Lastname1, ru.Lastname2, time.Now()).Scan(&userID, &name)
	if err != nil {
		return 0, "", fmt.Errorf("insert user: %w", err)
	}
	return userID, name, nil
}

func (pr *userRepository) RegisterPatient(ctx context.Context, tx pgx.Tx, ru entities.PatientUser) (int, string, error) {
	var userID int
	var name string
	query := "INSERT INTO users (name, lastname1, lastname2, created_at) VALUES ($1, $2, $3, $4) RETURNING id, name"
	err := tx.QueryRow(ctx, query, ru.Name, ru.Lastname1, ru.Lastname2, time.Now()).Scan(&userID, &name)
	if err != nil {
		return 0, "", fmt.Errorf("insert user: %w", err)
	}
	return userID, name, nil
}

// Función para registrar la cuenta utilizando el userID generado
func (pr *userRepository) RegisterAccount(ctx context.Context, tx pgx.Tx, ru entities.RegisterUserByAdmin, userID int) (string, error) {
	var email string
	query := "INSERT INTO account (id, user_id, email, password, rol, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING email"
	err := tx.QueryRow(ctx, query, ru.AccountID, userID, ru.Email, ru.Password, ru.Rol, time.Now()).Scan(&email)
	if err != nil {
		return "", fmt.Errorf("insert account: %w", err)
	}
	return email, nil
}

func (pr *userRepository) RegisterAccountPatient(ctx context.Context, tx pgx.Tx, ru entities.PatientUser, userID int) (string, error) {
	var email string
	query := "INSERT INTO account (id, user_id, email, password, rol, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING email"
	err := tx.QueryRow(ctx, query, ru.AccountID, userID, ru.Email, ru.Password, ru.Rol, time.Now()).Scan(&email)
	if err != nil {
		return "", fmt.Errorf("insert account: %w", err)
	}
	return email, nil
}

func (ur *userRepository) RegisterPatientTransaction(ctx context.Context, rp entities.PatientUser) (response entities.UserResponse, err error) {
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

	// Registrar el usuario y obtener el userID generado
	userID, name, err := ur.RegisterPatient(ctxTx, tx, rp)
	if err != nil {
		return response, fmt.Errorf("failed to register user: %w", err)
	}

	// Registrar la cuenta utilizando el userID generado
	email, err := ur.RegisterAccountPatient(ctxTx, tx, rp, userID)
	if err != nil {
		return response, fmt.Errorf("failed to register account: %w", err)
	}

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	_, _, err = ur.RegisterPatient(ctxTx, tx, rp)
	if err != nil {
		return response, fmt.Errorf("failed to register patient: %w", err)
	}

	// Construir la respuesta
	response = entities.UserResponse{
		Name:  name,
		Email: email,
	}

	return response, nil
}
