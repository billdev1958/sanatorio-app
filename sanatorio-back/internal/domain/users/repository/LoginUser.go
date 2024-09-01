package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"
	password "sanatorioApp/pkg/pass"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (pr *userRepository) LoginUser(ctx context.Context, lu entities.LoginUser) (entities.LoginResponse, error) {

	var accountID uuid.UUID
	var role int
	var hashedPassword string

	query := "SELECT id, rol, password FROM account WHERE email = $1"
	err := pr.storage.DbPool.QueryRow(ctx, query, lu.Email).Scan(&accountID, &role, &hashedPassword)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entities.LoginResponse{}, fmt.Errorf("user not found")
		}
		return entities.LoginResponse{}, fmt.Errorf("error querying user: %w", err)
	}

	if !password.CheckPasswordHash(lu.Password, hashedPassword) {
		return entities.LoginResponse{}, fmt.Errorf("invalid password")
	}

	response := entities.LoginResponse{
		AccountID: accountID,
		Role:      role,
	}

	return response, nil
}
