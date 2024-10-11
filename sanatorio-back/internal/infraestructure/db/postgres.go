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
	rolesValues := [4]string{"SuperAdmin", "Admin", "Doctor", "Patient"}

	var count int

	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM cat_rol").Scan(&count)
	if err != nil {
		return fmt.Errorf("count roles: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla cat_rol ya contiene datos")
		return nil
	}

	query := "INSERT INTO cat_rol (name, created_at) VALUES($1, $2)"
	for _, value := range rolesValues {
		_, err = storage.DbPool.Exec(ctx, query, value, time.Now())
		if err != nil {
			return fmt.Errorf("insert roles: %w", err)
		}
	}

	fmt.Println("Valores insertados correctamente en cat_rol")
	return nil
}

func (storage *PgxStorage) SeedPermissions(ctx context.Context) (err error) {
	permissionsValues := [16]string{"CreateUsers", "ViewUsers", "EditUsers", "DeleteUsers", "CreateSchedule", "ViewSchedule", "EditSchedule", "DeleteSchedule", "CreateAppointment", "ViewAppointment", "EditAppointment", "DeleteAppointment", "CreateMedicalHistory", "ViewMedicalHistory", "EditMedicalHistory", "DeleteMedicalHistory"}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM permissions").Scan(&count)
	if err != nil {
		return fmt.Errorf("count permissions: %w", err)
	}
	if count > 0 {
		fmt.Println("La tabla permissions ya contiene datos")
		return nil
	}

	query := "INSERT INTO permissions (name)"
	for _, value := range permissionsValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert permissions: %w", err)
		}
	}

	fmt.Println("Valores insertados correctamente en permissions")
	return nil

}

func (storage *PgxStorage) SeedRolePermissions(ctx context.Context) (err error) {
	// superAdmin = 1
	superAdminPermissions := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	}

	adminPermissions := []int{
		1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	}

	doctorPermissions := []int{
		10, 11, 12, 13, 14, 15, 16,
	}

	patientPermissions := []int{
		9, 10, 11, 12, 14,
	}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM role_permissions").Scan(&count)
	if err != nil {
		return fmt.Errorf("count role_permissions %w: ", err)
	}

	if count > 0 {
		fmt.Println("La tabla role_permissions ya contiene datos")
		return nil
	}

	err = insertPermissions(ctx, storage, 1, superAdminPermissions)
	if err != nil {
		return err
	}

	err = insertPermissions(ctx, storage, 2, adminPermissions)
	if err != nil {
		return err
	}

	err = insertPermissions(ctx, storage, 3, doctorPermissions)
	if err != nil {
		return err
	}

	err = insertPermissions(ctx, storage, 4, patientPermissions)
	if err != nil {
		return err
	}

	fmt.Println("permisos insertados correctamente en la tabla role_permissions")

	return nil
}

func insertPermissions(ctx context.Context, storage *PgxStorage, roleID int, permissions []int) error {
	query := "INSERT INTO role_permissions (role_id, permission_id) VALUES($1, $2)"

	for _, permission := range permissions {
		_, err := storage.DbPool.Exec(ctx, query, roleID, permission)
		if err != nil {
			return fmt.Errorf("error inserting role %d, permission %d: %w", roleID, permission, err)
		}
	}
	return nil
}

func (storage *PgxStorage) SeedOfficeStatus(ctx context.Context) (err error) {
	statusValues := [3]string{"Disponible", "No disponible", "No asignado"}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM office_status").Scan(&count)
	if err != nil {
		return fmt.Errorf("count office_status: %w", err)
	}
	if count > 0 {
		fmt.Println("La tabla office_status ya contiene datos")
		return nil
	}

	query := "INSERT INTO office_status (name, created_at) VALUES($1, $2)"
	for _, value := range statusValues {
		_, err = storage.DbPool.Exec(ctx, query, value, time.Now())
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

func (storage *PgxStorage) SeedSpecialties(ctx context.Context) (err error) {
	specialtiesValues := [5]string{"Medicina General", "Cardiologo", "Dermatologo", "Pediatra", "Ginecologia"}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM cat_specialty").Scan(&count)
	if err != nil {
		return fmt.Errorf("count specialties: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla cat_specialty ya contiene datos")
		return nil
	}

	query := "INSERT INTO cat_specialty (name, created_at) VALUES($1, $2)"
	for _, value := range specialtiesValues {
		_, err = storage.DbPool.Exec(ctx, query, value, time.Now())
		if err != nil {
			return fmt.Errorf("insert specialties: %w", err)
		}
	}

	fmt.Println("Especialidades insertadas correctamente en cat_specialty")
	return nil
}

func (storage *PgxStorage) SeedAppointmentStatus(ctx context.Context) (err error) {
	// Estados de la cita que vamos a insertar
	statusValues := [2]string{"Confirmada", "Cancelada"}

	// Verificar si ya hay datos en la tabla appointment_status
	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM appointment_status").Scan(&count)
	if err != nil {
		return fmt.Errorf("count appointment status: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla appointment_status ya contiene datos")
		return nil
	}

	// Query para insertar los estados
	query := "INSERT INTO appointment_status (name, created_at) VALUES($1, $2)"
	for _, value := range statusValues {
		_, err = storage.DbPool.Exec(ctx, query, value, time.Now())
		if err != nil {
			return fmt.Errorf("insert appointment status: %w", err)
		}
	}

	fmt.Println("Estados de citas insertados correctamente en appointment_status")
	return nil
}
