package postgres

import (
	"context"
	"fmt"
	password "sanatorioApp/pkg/pass"

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

	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM cat_role").Scan(&count)
	if err != nil {
		return fmt.Errorf("count roles: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla cat_role ya contiene datos")
		return nil
	}

	query := "INSERT INTO cat_role (name) VALUES($1)"
	for _, value := range rolesValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert roles: %w", err)
		}
	}

	fmt.Println("Valores insertados correctamente en cat_role")
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

	query := "INSERT INTO permissions (name) VALUES ($1)"
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
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM role_permission").Scan(&count)
	if err != nil {
		return fmt.Errorf("count role_permissions %w: ", err)
	}

	if count > 0 {
		fmt.Println("La tabla role_permission ya contiene datos")
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

	fmt.Println("permisos insertados correctamente en la tabla role_permission")

	return nil
}

func insertPermissions(ctx context.Context, storage *PgxStorage, roleID int, permissions []int) error {
	query := "INSERT INTO role_permission (role_id, permission_id) VALUES($1, $2)"

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
	// Verificar si ya existe un usuario con el correo electrónico especificado en la tabla account
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

	// Insertar cuenta en la tabla account
	queryAccount := `
		INSERT INTO account (id, dependency_id, phone, email, password, role_id) 
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = storage.DbPool.Exec(ctx, queryAccount, accountID, 1, "1234567890", "bilxd1958@gmail.com", hashedPassword, 1) // 1 es el ID del rol de administrador
	if err != nil {
		return fmt.Errorf("insert account: %w", err)
	}

	// Insertar en la tabla super_admin
	querySuperAdmin := `
		INSERT INTO super_admin (account_id, first_name, last_name1, last_name2, curp) 
		VALUES ($1, $2, $3, $4, $5)`
	_, err = storage.DbPool.Exec(ctx, querySuperAdmin, accountID, "Billy", "Rivera", "Salinas", "RISB010314HMCVLLA0")
	if err != nil {
		return fmt.Errorf("insert super_admin: %w", err)
	}

	// Insertar en la tabla user_roles para asignar el rol de administrador al usuario
	queryUserRole := `
		INSERT INTO user_roles (account_id, role_id) 
		VALUES ($1, $2)`
	_, err = storage.DbPool.Exec(ctx, queryUserRole, accountID, 1) // 1 es el ID del rol de administrador
	if err != nil {
		return fmt.Errorf("insert user_roles: %w", err)
	}

	fmt.Println("Usuario administrador insertado correctamente")
	return nil
}

func (storage *PgxStorage) SeedSpecialties(ctx context.Context) (err error) {
	specialtiesValues := [5]string{"Medicina General", "Cardiologo", "Dermatologo", "Pediatra", "Ginecologia"}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM specialty").Scan(&count)
	if err != nil {
		return fmt.Errorf("count specialties: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla specialty ya contiene datos")
		return nil
	}

	query := "INSERT INTO specialty (name) VALUES($1)"
	for _, value := range specialtiesValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert specialties: %w", err)
		}
	}

	fmt.Println("Especialidades insertadas correctamente en specialty")
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
	query := "INSERT INTO appointment_status (name) VALUES($1)"
	for _, value := range statusValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert appointment status: %w", err)
		}
	}

	fmt.Println("Estados de citas insertados correctamente en appointment_status")
	return nil
}

func (storage *PgxStorage) SeedDependencies(ctx context.Context) (err error) {
	dependenciesValues := [4]string{"Administrativo", "FAAPA", "SUTES", "Estudiante"}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM cat_dependencies").Scan(&count)
	if err != nil {
		return fmt.Errorf("count cat_dependencies status: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla cat_dependencies ya contiene datos")
		return nil
	}

	query := "INSERT INTO cat_dependencies (name) VALUES($1)"
	for _, value := range dependenciesValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert cat_dependencies status: %w", err)

		}

	}
	fmt.Println("Dependencias insertadas correctamente en cat_dependencies")
	return nil
}
