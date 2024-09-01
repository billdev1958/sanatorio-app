package repository

import (
	"context"
	"fmt"
	password "sanatorioApp/pkg/pass"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (pr *userRepository) CheckAdminPassword(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, pass string) (bool, error) {
	var hashedPassword string
	var roleID int

	query := `
		SELECT a.password, a.rol 
		FROM account a 
		WHERE a.id = $1`

	// Aquí debes asignar el resultado de la consulta al hashedPassword y roleID
	err := tx.QueryRow(ctx, query, accountID).Scan(&hashedPassword, &roleID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("user not found")
		}
		return false, err
	}

	// Verificar si el roleID es igual a 1 (asumiendo que 1 es el ID del rol de administrador)
	if roleID != 1 {
		return false, fmt.Errorf("user is not an admin")
	}

	// Ahora, verifica la contraseña usando el hash recuperado de la base de datos
	isValid := password.CheckPasswordHash(pass, hashedPassword)
	if !isValid {
		return false, nil
	}

	return true, nil
}
