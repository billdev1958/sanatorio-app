package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/jackc/pgx/v5"
)

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
