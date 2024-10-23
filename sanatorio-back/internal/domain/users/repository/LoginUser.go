package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"
	password "sanatorioApp/pkg/pass"

	"github.com/jackc/pgx/v5"
)

func (pr *userRepository) LoginUser(ctx context.Context, lu entities.Account) (entities.Account, error) {
	var account entities.Account

	// Asegúrate de que estás recuperando el UUID y el rol correctamente
	query := "SELECT id, role_id, password FROM account WHERE email = $1"
	err := pr.storage.DbPool.QueryRow(ctx, query, lu.Email).Scan(&account.ID, &account.Rol, &account.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return account, fmt.Errorf("user not found")
		}
		return account, fmt.Errorf("error querying user: %w", err)
	}

	// Verificar si la contraseña es válida
	if !password.CheckPasswordHash(lu.Password, account.Password) {
		return account, fmt.Errorf("invalid password")
	}

	// Limpiar la contraseña del objeto antes de devolverlo por seguridad
	account.Password = ""

	// Si la autenticación es exitosa, retornar la cuenta
	return account, nil
}
