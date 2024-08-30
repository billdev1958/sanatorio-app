package repository

import (
	"context"
	"fmt"
	"log"
	user "sanatorioApp/internal/domain/users"
	postgres "sanatorioApp/internal/infraestructure/db"
	"time"
)

type userRepository struct {
	storage *postgres.PgxStorage
}

func NewUserRepository(storage *postgres.PgxStorage) user.Repository {
	return &userRepository{storage: storage}
}

func (pr *userRepository) RegisterUser(ctx context.Context) (message string, err error) {
	// Crear un contexto con timeout de 5 segundos
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Consulta simple para verificar la conexión y devolver "Hola, Mundo!"
	var result string
	err = pr.storage.DbPool.QueryRow(ctx, "SELECT 'Hola, Mundo!'").Scan(&result)
	if err != nil {
		log.Printf("error executing query: %v", err)
		return "", fmt.Errorf("execute query: %w", err)
	}

	return result, nil
}
