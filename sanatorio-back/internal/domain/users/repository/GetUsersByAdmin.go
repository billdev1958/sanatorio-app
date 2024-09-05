package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) GetUserByID(ctx context.Context, userID int) (entities.Users, error) {
	var user entities.Users

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, su.curp, a.id AS account_id
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		LEFT JOIN super_user su ON a.id = su.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL AND u.id = $1
	`

	err := ur.storage.DbPool.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Name,
		&user.Lastname1,
		&user.Lastname2,
		&user.Email,
		&user.Rol,
		&user.Curp,
		&user.AccountID,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return user, fmt.Errorf("user with ID %v not found", userID)
		}
		return user, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (ur *userRepository) GetUsers(ctx context.Context) ([]entities.Users, error) {
	var users []entities.Users

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, 
		       su.curp, su.created_at
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
		var curp sql.NullString
		var createdAt sql.NullTime

		// Escanear los valores de las filas, incluidos los que pueden ser nulos
		err := rows.Scan(&user.ID, &user.Name, &user.Lastname1, &user.Lastname2, &user.Email, &user.Rol, &curp, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Asignar los valores, verificando si son válidos
		if curp.Valid {
			user.Curp = curp.String
		} else {
			user.Curp = "" // Asignar valor por defecto si es NULL
		}

		if createdAt.Valid {
			user.Created_At = createdAt.Time
		} else {
			user.Created_At = time.Time{} // Asignar valor por defecto si es NULL
		}

		users = append(users, user)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}

func (ur *userRepository) GetDoctorByID(ctx context.Context, userID int) (entities.Doctors, error) {
	var doctor entities.Doctors

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, a.id AS account_id, du.medical_license, du.id_specialty
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		INNER JOIN doctor_user du ON a.id = du.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL AND du.deleted_at IS NULL AND u.id = $1
	`

	err := ur.storage.DbPool.QueryRow(ctx, query, userID).Scan(
		&doctor.ID,
		&doctor.Name,
		&doctor.Lastname1,
		&doctor.Lastname2,
		&doctor.Email,
		&doctor.Rol,
		&doctor.MedicalLicense,
		&doctor.Specialty,
		&doctor.AccountID,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return doctor, fmt.Errorf("doctor with ID %v not found", userID)
		}
		return doctor, fmt.Errorf("failed to get doctor by ID: %w", err)
	}

	return doctor, nil
}

func (ur *userRepository) GetDoctors(ctx context.Context) ([]entities.Doctors, error) {
	var doctors []entities.Doctors

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, du.medical_license, du.specialty_id
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

func (ur *userRepository) GetAllUsers(ctx context.Context) ([]interface{}, error) {
	var users []interface{}

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, a.created_at, a.id AS account_id, du.medical_license, du.specialty_id
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		LEFT JOIN doctor_user du ON a.id = du.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL
	`

	rows, err := ur.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, lastname1, lastname2, email string
		var rol int
		var createdAt time.Time
		var accountID uuid.UUID
		var medicalLicense sql.NullString // Usamos NullString para manejar valores NULL
		var specialty sql.NullInt32       // Usamos NullInt32 para manejar valores NULL

		err := rows.Scan(&id, &name, &lastname1, &lastname2, &email, &rol, &createdAt, &accountID, &medicalLicense, &specialty)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		switch rol {
		case 1: // SuperUsuario
			superUser := entities.SuperUser{
				User: entities.User{
					ID:        id,
					Name:      name,
					Lastname1: lastname1,
					Lastname2: lastname2,
				},
				Email:      email,
				Rol:        rol,
				Created_At: createdAt,
				AccountID:  accountID,
			}
			users = append(users, superUser)

		case 2: // Doctor
			doctor := entities.Doctors{
				User: entities.User{
					ID:        id,
					Name:      name,
					Lastname1: lastname1,
					Lastname2: lastname2,
				},
				Email:          email,
				Rol:            rol,
				MedicalLicense: medicalLicense.String, // Asignamos el valor de NullString
				Specialty:      int(specialty.Int32),  // Asignamos el valor de NullInt32
				AccountID:      accountID,
			}
			users = append(users, doctor)

		case 3: // Paciente
			patient := entities.Users{
				User: entities.User{
					ID:        id,
					Name:      name,
					Lastname1: lastname1,
					Lastname2: lastname2,
				},
				Email:      email,
				Rol:        rol,
				Created_At: createdAt,
				AccountID:  accountID,
			}
			users = append(users, patient)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}
