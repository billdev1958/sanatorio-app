package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxStorage struct {
	DbPool *pgxpool.Pool
}

func NewPgxStorage(dbPool *pgxpool.Pool) *PgxStorage {
	return &PgxStorage{DbPool: dbPool}
}

func (storage *PgxStorage) SeedRoles(ctx context.Context) (err error) {
	rolesValues := [3]string{"Super Usuario", "Doctor", "Paciente"}

	var count int

	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM cat_rol").Scan(&count)
	if err != nil {
		return fmt.Errorf("count roles: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla cat_rol ya contiene datos")
		return nil
	}

	query := "INSERT INTO cat_rol (name) VALUES($1)"
	for _, value := range rolesValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert roles: %w", err)
		}
	}

	fmt.Println("Valores insertados correctamente en cat_rol")
	return nil
}

func (storage *PgxStorage) SeedOfficeStatus(ctx context.Context) (err error) {
	statusValues := [3]string{"Disponible", "No disponible", "Mantenimiento"}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM office_status").Scan(&count)
	if err != nil {
		return fmt.Errorf("count office_status: %w", err)
	}
	if count > 0 {
		fmt.Println("La tabla office_status ya contiene datos")
		return nil
	}

	query := "INSERT INTO office_status (name) VALUES($1)"
	for _, value := range statusValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert office status: %w", err)
		}
	}
	fmt.Println("Valores insertados correctamente en office_status")
	return nil
}
