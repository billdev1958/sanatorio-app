package repository

import (
	"context"
	"fmt"
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
