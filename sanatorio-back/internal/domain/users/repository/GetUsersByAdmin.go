package repository

/*func (ur *userRepository) GetSuperUserByID(ctx context.Context, superUserID int) (entities.SuperUser, error) {
	var user entities.SuperUser

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol, su.curp, a.id AS account_id
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		LEFT JOIN super_user su ON a.id = su.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL AND u.id = $1
	`

	err := ur.storage.DbPool.QueryRow(ctx, query, superUserID).Scan(
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
			return user, fmt.Errorf("user with ID %v not found", superUserID)
		}
		return user, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (ur *userRepository) GetSuperAdmins(ctx context.Context) ([]entities.SuperUser, error) {
	var users []entities.SuperUser

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
		var user entities.SuperUser

		// Escanear los valores directamente
		err := rows.Scan(&user.ID, &user.Name, &user.Lastname1, &user.Lastname2, &user.Email, &user.Rol, &user.Curp, &user.Created_At)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Agregar el usuario a la lista
		users = append(users, user)
	}

	// Comprobar si hubo errores durante la iteración
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}

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
		&doctor.ID,
		&doctor.Name,
		&doctor.Lastname1,
		&doctor.Lastname2,
		&doctor.Email,
		&doctor.Rol,
		&doctor.MedicalLicense,
		&doctor.SpecialtyID,
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

func (ur *userRepository) GetDoctors(ctx context.Context) ([]entities.DoctorUser, error) {
	var doctors []entities.DoctorUser

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
		var doctor entities.DoctorUser
		err := rows.Scan(&doctor.ID, &doctor.Name, &doctor.Lastname1, &doctor.Lastname2, &doctor.Email, &doctor.Rol, &doctor.MedicalLicense, &doctor.SpecialtyID)
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

func (ur *userRepository) GetPatients(ctx context.Context) ([]entities.PatientUser, error) {
	var users []entities.PatientUser

	query := `
		SELECT u.id, u.name, u.lastname1, u.lastname2, a.email, a.rol,
		       su.curp, su.created_at
		FROM users u
		INNER JOIN account a ON u.id = a.user_id
		LEFT JOIN patient su ON a.id = su.account_id
		WHERE u.deleted_at IS NULL AND a.deleted_at IS NULL
	`

	rows, err := ur.storage.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.PatientUser

		// Escanear los valores directamente
		err := rows.Scan(&user.ID, &user.Name, &user.Lastname1, &user.Lastname2, &user.Email, &user.Rol, &user.Curp, &user.Created_At)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Agregar el usuario a la lista
		users = append(users, user)
	}

	// Comprobar si hubo errores durante la iteración
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}
*/
