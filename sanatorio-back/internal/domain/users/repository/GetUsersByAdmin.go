package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/jackc/pgx/v5"
)

func (ur *userRepository) GetDoctorByID(ctx context.Context, doctorID int) (entities.DoctorUser, error) {
	var doctor entities.DoctorUser

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, a.id AS account_id, du.medical_license, du.id_specialty
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		INNER JOIN doctor_user du ON a.id = du.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL AND du.deleted_at IS NULL AND u.id = $1
	`

	err := ur.storage.DbPool.QueryRow(ctx, query, doctorID).Scan(
		&doctor.AccountID,
		&doctor.FirstName,
		&doctor.LastName1,
		&doctor.LastName2,
		&doctor.MedicalLicense,
		&doctor.SpecialtyLicense,
		&doctor.AccountID,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return doctor, fmt.Errorf("doctor with ID %v not found", doctorID)
		}
		return doctor, fmt.Errorf("failed to get doctor by ID: %w", err)
	}

	return doctor, nil
}
