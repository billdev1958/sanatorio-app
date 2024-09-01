package postgres

import (
	"context"
	"fmt"
	password "sanatorioApp/pkg/pass"
	"time"

	"github.com/google/uuid"
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

func (storage *PgxStorage) SeedAdminUser(ctx context.Context) (err error) {
	// Verificar si ya existe un usuario con el correo electrónico especificado
	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM account WHERE email = $1", "bilxd1958@gmail.com").Scan(&count)
	if err != nil {
		return fmt.Errorf("count admin user: %w", err)
	}

	if count > 0 {
		fmt.Println("El usuario administrador ya existe")
		return nil
	}

	// Hashear la contraseña
	hashedPassword, err := password.HashPassword("1a2s3d4f")
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	// Generar UUID para el account_id
	accountID := uuid.New()

	// Insertar usuario en la tabla users
	var userID int
	queryUser := "INSERT INTO users (name, lastname1, lastname2, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err = storage.DbPool.QueryRow(ctx, queryUser, "Billy", "Rivera", "Salinas", time.Now()).Scan(&userID)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	// Insertar cuenta en la tabla account
	queryAccount := "INSERT INTO account (id, user_id, email, password, rol, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = storage.DbPool.Exec(ctx, queryAccount, accountID, userID, "bilxd1958@gmail.com", hashedPassword, 1, time.Now())
	if err != nil {
		return fmt.Errorf("insert account: %w", err)
	}

	// Insertar en la tabla super_user
	querySuperUser := "INSERT INTO super_user (account_id, curp, created_at) VALUES ($1, $2, $3)"
	_, err = storage.DbPool.Exec(ctx, querySuperUser, accountID, "RISB010314HMCVLLA0", time.Now())
	if err != nil {
		return fmt.Errorf("insert super_user: %w", err)
	}

	fmt.Println("Usuario administrador insertado correctamente")
	return nil
}
