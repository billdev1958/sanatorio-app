package repository

import (
	"context"
	"fmt"
	"log"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	postgres "sanatorioApp/internal/infraestructure/db"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	storage *postgres.PgxStorage
}

func NewUserRepository(storage *postgres.PgxStorage) user.Repository {
	return &userRepository{storage: storage}
}

func (ur *userRepository) RegisterUserTransaction(ctx context.Context, ru entities.RegisterUserByAdmin) (response entities.UserResponse, err error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return response, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctxTx); rbErr != nil {
				log.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			err = tx.Commit(ctxTx)
			if err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctxTx, tx, ru.AccountAdminID, ru.AdminPassword)
	if err != nil {
		return response, fmt.Errorf("failed to authenticate admin: %w", err)
	}
	if !isValid {
		return response, fmt.Errorf("failed to authenticate admin: invalid password")
	}

	// Registrar el usuario y obtener el userID generado
	userID, name, err := ur.RegisterUser(ctxTx, tx, ru)
	if err != nil {
		return response, fmt.Errorf("failed to register user: %w", err)
	}

	// Registrar la cuenta utilizando el userID generado
	email, err := ur.RegisterAccount(ctxTx, tx, ru, userID)
	if err != nil {
		return response, fmt.Errorf("failed to register account: %w", err)
	}

	// Manejar el tipo de usuario según el rol
	switch ru.Rol {
	case entities.SuperUsuario, entities.Patient:
		err = ur.RegisterTypeUser(ctxTx, tx, ru)
		if err != nil {
			return response, fmt.Errorf("failed to register user type: %w", err)
		}
	default:
		return response, fmt.Errorf("unknown role: %d", ru.Rol)
	}

	// Construir la respuesta
	response = entities.UserResponse{
		Name:  name,
		Email: email,
	}

	return response, nil
}

func (ur *userRepository) RegisterDoctorTransaction(ctx context.Context, rd entities.RegisterDoctorByAdmin) (response entities.UserResponse, err error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return response, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctxTx); rbErr != nil {
				log.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			err = tx.Commit(ctxTx)
			if err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctxTx, tx, rd.AccountAdminID, rd.AdminPassword)
	if err != nil || !isValid {
		return response, fmt.Errorf("failed to authenticate admin: %w", err)
	}

	// Registrar el usuario y obtener el userID generado
	userID, name, err := ur.RegisterUser(ctxTx, tx, entities.RegisterUserByAdmin{
		User:       rd.User,
		Account:    rd.Account,
		DocumentID: rd.DocumentID,
	})
	if err != nil {
		return response, err
	}

	// Registrar la cuenta utilizando el userID generado
	email, err := ur.RegisterAccount(ctxTx, tx, entities.RegisterUserByAdmin{
		User:       rd.User,
		Account:    rd.Account,
		DocumentID: rd.DocumentID,
	}, userID)
	if err != nil {
		return response, err
	}

	// Registrar al doctor en su tabla correspondiente
	err = ur.RegisterTypeDoctor(ctxTx, tx, rd)
	if err != nil {
		return response, err
	}

	// Construir la respuesta
	response = entities.UserResponse{
		Name:  name,
		Email: email,
	}

	return response, nil
}

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

// Función para registrar el tipo de usuario en la tabla correspondiente
func (pr *userRepository) RegisterTypeUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUserByAdmin) error {
	var query string
	var values []interface{}

	if ru.Rol == entities.SuperUsuario {
		query = "INSERT INTO super_user (account_id, curp, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())"
		values = []interface{}{ru.AccountID, ru.DocumentID}
	} else if ru.Rol == entities.Patient {
		query = "INSERT INTO patient_user (account_id, curp, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())"
		values = []interface{}{ru.AccountID, ru.DocumentID}
	} else {
		return fmt.Errorf("unknown role: %d", ru.Rol)
	}

	_, err := tx.Exec(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("insert into table: %w", err)
	}

	return nil
}

// Función para registrar el doctor en su tabla correspondiente
func (pr *userRepository) RegisterTypeDoctor(ctx context.Context, tx pgx.Tx, rd entities.RegisterDoctorByAdmin) error {
	query := "INSERT INTO doctor_user (user_id, account_id, specialty_id, medical_license) VALUES ($1, $2, $3, $4)"
	_, err := tx.Exec(ctx, query, rd.AccountID, rd.SpecialtyID, rd.DocumentID)
	if err != nil {
		return fmt.Errorf("insert into table: %w", err)
	}
	return nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, accountID string) (entities.Users, error) {
	var user entities.Users

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, su.curp
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		LEFT JOIN super_user su ON a.id = su.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL AND a.id = $1
	`

	err := ur.storage.DbPool.QueryRow(ctx, query, accountID).Scan(
		&user.ID,
		&user.Name,
		&user.Lastname1,
		&user.Lastname2,
		&user.Email,
		&user.Rol,
		&user.Curp,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return user, fmt.Errorf("user with account ID %s not found", accountID)
		}
		return user, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (ur *userRepository) GetUsers(ctx context.Context) ([]entities.Users, error) {
	var users []entities.Users

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, su.curp
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		LEFT JOIN super_user su ON a.id = su.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL
	`

	rows, err := ur.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.Users
		err := rows.Scan(&user.ID, &user.Name, &user.Lastname1, &user.Lastname2, &user.Email, &user.Rol, &user.Curp)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}

func (ur *userRepository) GetDoctorByID(ctx context.Context, accountID string) (entities.Doctors, error) {
	var doctor entities.Doctors

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, du.medical_license, du.id_specialty
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		INNER JOIN doctor_user du ON a.id = du.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL AND du.deleted_at IS NULL AND a.id = $1
	`

	err := ur.storage.DbPool.QueryRow(ctx, query, accountID).Scan(
		&doctor.ID,
		&doctor.Name,
		&doctor.Lastname1,
		&doctor.Lastname2,
		&doctor.Email,
		&doctor.Rol,
		&doctor.MedicalLicense,
		&doctor.Specialty,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return doctor, fmt.Errorf("doctor with account ID %s not found", accountID)
		}
		return doctor, fmt.Errorf("failed to get doctor by ID: %w", err)
	}

	return doctor, nil
}

func (ur *userRepository) GetDoctors(ctx context.Context) ([]entities.Doctors, error) {
	var doctors []entities.Doctors

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, du.medical_license, du.id_specialty
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		INNER JOIN doctor_user du ON a.id = du.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL AND du.deleted_at IS NULL
	`

	rows, err := ur.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var doctor entities.Doctors
		err := rows.Scan(&doctor.ID, &doctor.Name, &doctor.Lastname1, &doctor.Lastname2, &doctor.Email, &doctor.Rol, &doctor.MedicalLicense, &doctor.Specialty)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		doctors = append(doctors, doctor)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return doctors, nil
}

// UpdateUser actualiza los campos de un usuario en las tablas users y account.
// Solo se actualizan los campos que no están vacíos en la solicitud.
// La operación se realiza dentro de una transacción para garantizar la consistencia de los datos.
func (ur *userRepository) UpdateUser(ctx context.Context, userUpdate entities.UpdateUser) (string, error) {
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx) // Asegura que la transacción se revierta si no se confirma

	// Verificar que el usuario existe a partir del account_id
	var existingUserID int
	err = tx.QueryRow(ctx, "SELECT user_id FROM account WHERE id = $1", userUpdate.AccountID).Scan(&existingUserID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("user with account ID %s not found", userUpdate.AccountID)
		}
		return "", fmt.Errorf("failed to check if user exists: %w", err)
	}

	// Actualización condicional de los campos en la tabla 'users'
	var setClauses []string
	var args []interface{}
	argIndex := 1

	if userUpdate.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, userUpdate.Name)
		argIndex++
	}
	if userUpdate.Lastname1 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname1 = $%d", argIndex))
		args = append(args, userUpdate.Lastname1)
		argIndex++
	}
	if userUpdate.Lastname2 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname2 = $%d", argIndex))
		args = append(args, userUpdate.Lastname2)
		argIndex++
	}

	setClauses = append(setClauses, "updated_at = NOW()")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d RETURNING name", strings.Join(setClauses, ", "), argIndex)
	args = append(args, existingUserID)

	var updatedName string
	err = tx.QueryRow(ctx, query, args...).Scan(&updatedName)
	if err != nil {
		return "", fmt.Errorf("failed to update user: %w", err)
	}

	// Actualizar la cuenta si es necesario
	if userUpdate.Password != "" {
		accountSetClauses := []string{"updated_at = NOW()"}
		accountArgs := []interface{}{}
		accountIndex := 1

		if userUpdate.Password != "" {
			accountSetClauses = append(accountSetClauses, fmt.Sprintf("password = $%d", accountIndex))
			accountArgs = append(accountArgs, userUpdate.Password)
			accountIndex++
		}

		accountQuery := fmt.Sprintf("UPDATE account SET %s WHERE id = $%d", strings.Join(accountSetClauses, ", "), accountIndex)
		accountArgs = append(accountArgs, userUpdate.AccountID)

		_, err := tx.Exec(ctx, accountQuery, accountArgs...)
		if err != nil {
			return "", fmt.Errorf("failed to update account: %w", err)
		}
	}

	// Confirmar la transacción
	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("User %s updated successfully", updatedName), nil
}

// UpdateDoctor actualiza los campos de un doctor en las tablas users, account y doctor_user.
// La operación se realiza dentro de una transacción para garantizar la consistencia de los datos.
func (ur *userRepository) UpdateDoctor(ctx context.Context, doctorUpdate entities.UpdateDoctor) (string, error) {
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var setClauses []string
	var args []interface{}
	argIndex := 1

	if doctorUpdate.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, doctorUpdate.Name)
		argIndex++
	}
	if doctorUpdate.Lastname1 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname1 = $%d", argIndex))
		args = append(args, doctorUpdate.Lastname1)
		argIndex++
	}
	if doctorUpdate.Lastname2 != "" {
		setClauses = append(setClauses, fmt.Sprintf("lastname2 = $%d", argIndex))
		args = append(args, doctorUpdate.Lastname2)
		argIndex++
	}

	setClauses = append(setClauses, "updated_at = NOW()")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = (SELECT user_id FROM account WHERE id = $%d) RETURNING name",
		strings.Join(setClauses, ", "), argIndex)
	args = append(args, doctorUpdate.AccountID)

	var updatedName string
	err = tx.QueryRow(ctx, query, args...).Scan(&updatedName)
	if err != nil {
		return "", fmt.Errorf("failed to update doctor user: %w", err)
	}

	if doctorUpdate.Password != "" || doctorUpdate.MedicalLicense != "" || doctorUpdate.SpecialtyID != 0 {
		accountSetClauses := []string{"updated_at = NOW()"}
		accountArgs := []interface{}{}
		accountIndex := 1

		if doctorUpdate.Password != "" {
			accountSetClauses = append(accountSetClauses, fmt.Sprintf("password = $%d", accountIndex))
			accountArgs = append(accountArgs, doctorUpdate.Password)
			accountIndex++
		}

		accountQuery := fmt.Sprintf("UPDATE account SET %s WHERE id = $%d", strings.Join(accountSetClauses, ", "), accountIndex)
		accountArgs = append(accountArgs, doctorUpdate.AccountID)

		_, err := tx.Exec(ctx, accountQuery, accountArgs...)
		if err != nil {
			return "", fmt.Errorf("failed to update account: %w", err)
		}

		doctorSetClauses := []string{"updated_at = NOW()"}
		doctorArgs := []interface{}{}
		doctorIndex := 1

		if doctorUpdate.MedicalLicense != "" {
			doctorSetClauses = append(doctorSetClauses, fmt.Sprintf("medical_license = $%d", doctorIndex))
			doctorArgs = append(doctorArgs, doctorUpdate.MedicalLicense)
			doctorIndex++
		}
		if doctorUpdate.SpecialtyID != 0 {
			doctorSetClauses = append(doctorSetClauses, fmt.Sprintf("id_specialty = $%d", doctorIndex))
			doctorArgs = append(doctorArgs, doctorUpdate.SpecialtyID)
			doctorIndex++
		}

		doctorQuery := fmt.Sprintf("UPDATE doctor_user SET %s WHERE account_id = $%d",
			strings.Join(doctorSetClauses, ", "), doctorIndex)
		doctorArgs = append(doctorArgs, doctorUpdate.AccountID)

		_, err = tx.Exec(ctx, doctorQuery, doctorArgs...)
		if err != nil {
			return "", fmt.Errorf("failed to update doctor_user: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("Doctor %s updated successfully", updatedName), nil
}

// DeleteUser elimina un usuario y su cuenta asociada de la base de datos.
func (ur *userRepository) DeleteUser(ctx context.Context, accountID string) (string, error) {
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Verificar si existe un registro en 'account' con el account_id proporcionado
	var exists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM account WHERE id = $1)"
	err = tx.QueryRow(ctx, checkQuery, accountID).Scan(&exists)
	if err != nil {
		return "", fmt.Errorf("failed to check existence in account: %w", err)
	}

	if !exists {
		return "", fmt.Errorf("user with account ID %s not found", accountID)
	}

	// Eliminar el usuario asociado en la tabla 'users' utilizando el user_id de la tabla 'account'
	userQuery := `
		DELETE FROM users 
		WHERE id = (SELECT user_id FROM account WHERE id = $1)
		RETURNING name
	`
	var deletedName string
	err = tx.QueryRow(ctx, userQuery, accountID).Scan(&deletedName)
	if err != nil {
		return "", fmt.Errorf("failed to delete user: %w", err)
	}

	// Eliminar el registro de la tabla 'account'
	accountQuery := "DELETE FROM account WHERE id = $1"
	_, err = tx.Exec(ctx, accountQuery, accountID)
	if err != nil {
		return "", fmt.Errorf("failed to delete account: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("User %s deleted successfully", deletedName), nil
}

// DeleteDoctor elimina un doctor y su cuenta asociada de la base de datos.
func (ur *userRepository) DeleteDoctor(ctx context.Context, accountID string) (string, error) {
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Verificar si existe un registro en 'doctor_user' con el account_id proporcionado
	var exists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM doctor_user WHERE account_id = $1)"
	err = tx.QueryRow(ctx, checkQuery, accountID).Scan(&exists)
	if err != nil {
		return "", fmt.Errorf("failed to check existence in doctor_user: %w", err)
	}

	if !exists {
		return "", fmt.Errorf("doctor with account ID %s not found", accountID)
	}

	// Eliminar el registro de la tabla 'doctor_user' utilizando el account_id
	doctorQuery := `
		DELETE FROM doctor_user 
		WHERE account_id = $1
		RETURNING account_id
	`
	var deletedAccountID string
	err = tx.QueryRow(ctx, doctorQuery, accountID).Scan(&deletedAccountID)
	if err != nil {
		return "", fmt.Errorf("failed to delete doctor: %w", err)
	}

	// Eliminar el usuario asociado en la tabla 'users' utilizando el user_id de la tabla 'account'
	userQuery := `
		DELETE FROM users 
		WHERE id = (SELECT user_id FROM account WHERE id = $1)
		RETURNING name
	`
	var deletedName string
	err = tx.QueryRow(ctx, userQuery, accountID).Scan(&deletedName)
	if err != nil {
		return "", fmt.Errorf("failed to delete user: %w", err)
	}

	// Eliminar el registro de la tabla 'account'
	accountQuery := "DELETE FROM account WHERE id = $1"
	_, err = tx.Exec(ctx, accountQuery, accountID)
	if err != nil {
		return "", fmt.Errorf("failed to delete account: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("Doctor %s deleted successfully", deletedName), nil
}

// SoftDeleteUser marca un usuario como eliminado sin borrar su información.
func (ur *userRepository) SoftDeleteUser(ctx context.Context, accountID string) (string, error) {
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Verificar si existe un registro en 'account' con el account_id proporcionado
	var exists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM account WHERE id = $1)"
	err = tx.QueryRow(ctx, checkQuery, accountID).Scan(&exists)
	if err != nil {
		return "", fmt.Errorf("failed to check existence in account: %w", err)
	}

	if !exists {
		return "", fmt.Errorf("user with account ID %s not found", accountID)
	}

	// Marcar como eliminado el usuario en la tabla 'users' utilizando el user_id de la tabla 'account'
	query := `
		UPDATE users 
		SET deleted_at = NOW() 
		WHERE id = (SELECT user_id FROM account WHERE id = $1)
		RETURNING name
	`
	var updatedName string
	err = tx.QueryRow(ctx, query, accountID).Scan(&updatedName)
	if err != nil {
		return "", fmt.Errorf("failed to soft delete user: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("User %s soft deleted successfully", updatedName), nil
}

// SoftDeleteDoctor marca un doctor como eliminado sin borrar su información.
func (ur *userRepository) SoftDeleteDoctor(ctx context.Context, accountID string) (string, error) {
	tx, err := ur.storage.DbPool.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Verificar si existe un registro en 'doctor_user' con el account_id proporcionado
	var exists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM doctor_user WHERE account_id = $1)"
	err = tx.QueryRow(ctx, checkQuery, accountID).Scan(&exists)
	if err != nil {
		return "", fmt.Errorf("failed to check existence in doctor_user: %w", err)
	}

	if !exists {
		return "", fmt.Errorf("doctor with account ID %s not found", accountID)
	}

	// Marcar como eliminado el usuario en la tabla 'users' utilizando el user_id de la tabla 'account'
	query := `
		UPDATE users 
		SET deleted_at = NOW() 
		WHERE id = (SELECT user_id FROM account WHERE id = $1)
		RETURNING name
	`
	var updatedName string
	err = tx.QueryRow(ctx, query, accountID).Scan(&updatedName)
	if err != nil {
		return "", fmt.Errorf("failed to soft delete doctor: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return fmt.Sprintf("Doctor %s soft deleted successfully", updatedName), nil
}
