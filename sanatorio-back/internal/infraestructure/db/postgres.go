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
	rolesValues := [5]string{"SuperAdmin", "Admin", "Doctor", "Receptionist", "Patient"}

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
	permissionsValues := [40]string{
		"CreateSuperAdmin", "CreateAdmin", "CreateDoctor", "CreateReceptionist", "CreatePatient", "CreateBeneficairy",
		"ViewUsers", "EditSuperAdmin", "EditAdmin", "EditDoctor", "EditReceptionist", "EditPatient",
		"DeleteSuperAdmin", "DeleteAdmin", "DeleteDoctor", "DeleteReceptionist", "DeletePatient",
		"CreateSchedule", "ViewSchedule", "EditSchedule", "DeleteSchedule",
		"CreateAppointment", "ViewAppointment", "EditAppointment", "DeleteAppointment", "CancelAppointment",
		"CreateLaboratory", "EditLaboratory", "ViewLaboratory", "DeleteLaboratory",
		"CreateMedicalHistory", "ViewMedicalHistory", "EditMedicalHistory",
		"CreateEvolutionNote", "ViewEvolutionNote", "EditEvolutionNote",
		"CreatePrescription", "CreateIncapacity", "ViewIncapacity", "EditIncapacity",
	}

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
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26,
	}

	adminPermissions := []int{
		2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 32,
	}

	doctorPermissions := []int{
		23, 24, 25, 26, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	}

	receptionistPermissions := []int{
		18, 19, 20, 21, 23,
	}

	patientPermissions := []int{
		6, 22, 23, 24, 26,
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

	err = insertPermissions(ctx, storage, 4, receptionistPermissions)
	if err != nil {
		return err
	}

	err = insertPermissions(ctx, storage, 5, patientPermissions)
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

// SeedDays inserta los días de la semana en la tabla days
func (storage *PgxStorage) SeedDays(ctx context.Context) error {
	nameDays := []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}

	for dayOfWeek, name := range nameDays {
		if err := insertDay(ctx, storage, dayOfWeek, name); err != nil {
			return fmt.Errorf("error inserting day %s: %w", name, err)
		}
	}

	return nil
}

// insertDay inserta un único día en la tabla days
func insertDay(ctx context.Context, storage *PgxStorage, dayOfWeek int, name string) error {
	query := "INSERT INTO days (day_of_week, name) VALUES ($1, $2) ON CONFLICT (day_of_week) DO NOTHING"

	_, err := storage.DbPool.Exec(ctx, query, dayOfWeek, name)
	if err != nil {
		return fmt.Errorf("failed to insert day %s: %w", name, err)
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
		INSERT INTO account (id, dependency_id, phone, email, password, role_id, is_verified) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = storage.DbPool.Exec(ctx, queryAccount, accountID, 1, "1234567890", "bilxd1958@gmail.com", hashedPassword, 1, true) // 1 es el ID del rol de administrador
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

	fmt.Println("Usuario administrador insertado correctamente")
	return nil
}

func (storage *PgxStorage) SeedServices(ctx context.Context) (err error) {
	// Lista de nombres de servicios
	services := []string{
		"Audiometría",
		"Acupuntura",
		"Densitometría",
		"Educación especial",
		"Electrocardiografía",
		"Evaluación psicológica infantil",
		"Gerontología",
		"Ginecología",
		"Laboratorio clínico",
		"Mastografías",
		"Rayos X",
		"Medicina física y rehabilitación",
		"Medicina general matutino",
		"Medicina general vespertino",
		"Nutrición matutino",
		"Nutrición vespertino",
		"Odontología Matutino",
		"Odontología Vespertino",
		"Otorrinolaringología",
		"Prueba de detección por PCR de SARS-CoV-2",
		"Prueba de detección PCR de SARS-CoV-2 e Influenza",
		"Prueba PCR diagnóstico de Infecciones Urogenitales",
		"Prueba rápida de antígeno SARS-CoV-2 e Influenza",
		"Psicología clínica infantil",
		"Psicología matutino",
		"Psicología sabatino",
		"Psicología vespertino",
		"Terapia ocupacional",
		"Terapia física matutino",
		"Terapia física vespertino",
	}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM services").Scan(&count)
	if err != nil {
		return fmt.Errorf("count services: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla services ya contiene datos")
		return nil
	}

	query := "INSERT INTO services (name) VALUES($1)"
	for _, service := range services {
		_, err = storage.DbPool.Exec(ctx, query, service)
		if err != nil {
			return fmt.Errorf("insert services: %w", err)
		}
	}

	fmt.Println("Servicios insertados correctamente en la tabla services")
	return nil
}

func (storage *PgxStorage) SeedAppointmentStatus(ctx context.Context) (err error) {
	// Estados de la cita que vamos a insertar
	statusValues := [5]string{"Pendiente", "Confirmada", "Iniciada", "Concluida", "Cancelada"}

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
	dependenciesValues := [5]string{"Administrativo", "FAAPA", "SUTES", "Estudiante", "Externo"}

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

func (storage *PgxStorage) SeedMedicalInstitution(ctx context.Context) (err error) {
	mi := [6]string{"IMMS", "ISSSTE", "ISSEMYM", "SEDENA", "PEMEX", "NINGUNO"}

	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM cat_medical_institutions").Scan(&count)
	if err != nil {
		return fmt.Errorf("count cat_medical_institutions status: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla cat_medical_institutions ya contiene datos")
		return nil
	}

	query := "INSERT INTO cat_medical_institutions (name) VALUES($1)"
	for _, value := range mi {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert cat_medical_institutions status: %w", err)

		}

	}
	fmt.Println("Dependencias insertadas correctamente en cat_medical_institutions")
	return nil
}

// Seed de horarios y consultorios

func (storage *PgxStorage) SeedShifts(ctx context.Context) (err error) {
	// Estados de la cita que vamos a insertar
	statusValues := [2]string{"Matutino", "Vespertino"}

	// Verificar si ya hay datos en la tabla appointment_status
	var count int
	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM cat_shift").Scan(&count)
	if err != nil {
		return fmt.Errorf("count cat_shift: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla cat_shift ya contiene datos")
		return nil
	}

	// Query para insertar los estados
	query := "INSERT INTO cat_shift (name) VALUES($1)"
	for _, value := range statusValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert cat_shift: %w", err)
		}
	}

	fmt.Println("Estados de citas insertados correctamente en cat_shift")
	return nil
}

func (storage *PgxStorage) SeedOffices(ctx context.Context) (err error) {
	rolesValues := [48]string{
		"Consultorio: Audiometría",
		"Consultorio: Acupuntura",
		"Consultorio: Densitometría",
		"Consultorio: Educación especial",
		"Consultorio: Electrocardiografía",
		"Consultorio: Evaluación psicológica infantil",
		"Consultorio: Gerontología",
		"Consultorio: Ginecología",
		"Consultorio: Laboratorio clínico",
		"Consultorio: Mastografías",
		"Consultorio: Rayos X",
		"Consultorio: Medicina física y rehabilitación",
		"Consultorio: Nutrición matutino",
		"Consultorio: Nutrición vespertino",
		"Consultorio: Odontología Matutino",
		"Consultorio: Odontología Vespertino",
		"Consultorio: Otorrinolaringología",
		"Consultorio: Prueba de detección por PCR de SARS-CoV-2",
		"Consultorio: Prueba de detección PCR de SARS-CoV-2 e Influenza",
		"Consultorio: Prueba PCR diagnóstico de Infecciones Urogenitales",
		"Consultorio: Prueba rápida de antígeno SARS-CoV-2 e Influenza",
		"Consultorio: Psicología clínica infantil",

		"Consultorio 1: Medicina general matutino (Respiratorio)",
		"Consultorio 2: Medicina general matutino (Respiratorio)",
		"Consultorio 3: Medicina general matutino",
		"Consultorio 4: Medicina general matutino",
		"Consultorio 5: Medicina general matutino",
		"Consultorio 3: Medicina general vespertino",
		"Consultorio 4: Medicina general vespertino",
		"Consultorio 5: Medicina general vespertino (Respiratorio)",

		"Consultorio 2: Psicología Matutino ",
		"Consultorio 3: Psicología Matutino ",
		"Consultorio 4: Psicología Matutino ",
		"Consultorio 5: Psicología Matutino ",
		"Consultorio 6: Psicología Matutino ",

		"Consultorio 1: Psicología Sabatino ",
		"Consultorio 2: Psicología Sabatino ",
		"Consultorio 3: Psicología Sabatino ",

		"Consultorio 1: Psicología Vespertino ",
		"Consultorio 3: Psicología Vespertino ",
		"Consultorio 4: Psicología Vespertino ",
		"Consultorio 5: Psicología Vespertino ",

		"Consultorio 1: Terapia Ocupacional ",
		"Consultorio 2: Terapia Ocupacional ",

		"Consultorio 1: Terapia Fisica Matutino",
		"Consultorio 2: Terapia Fisica Matutino",

		"Consultorio 1: Terapia Fisica Vespertino",
		"Consultorio 2: Terapia Fisica Vespertino",
	}

	var count int

	err = storage.DbPool.QueryRow(ctx, "SELECT COUNT(*) FROM office").Scan(&count)
	if err != nil {
		return fmt.Errorf("count office: %w", err)
	}

	if count > 0 {
		fmt.Println("La tabla office ya contiene datos")
		return nil
	}

	query := "INSERT INTO office (name) VALUES($1)"
	for _, value := range rolesValues {
		_, err = storage.DbPool.Exec(ctx, query, value)
		if err != nil {
			return fmt.Errorf("insert roles: %w", err)
		}
	}

	fmt.Println("Valores insertados correctamente en office")
	return nil
}
