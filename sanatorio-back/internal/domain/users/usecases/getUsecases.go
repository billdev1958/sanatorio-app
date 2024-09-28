package usecase

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"
)

func (u *usecase) GetSuperAdminByID(ctx context.Context, userID int) (models.UserRequest, error) {
	// Llama al repositorio para obtener el usuario por ID
	userEntity, err := u.repo.GetSuperUserByID(ctx, userID)
	if err != nil {
		return models.UserRequest{}, err
	}

	// Crear el objeto UserRequest con los datos del usuario
	userData := models.UserRequest{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Lastname1: userEntity.Lastname1,
		Lastname2: userEntity.Lastname2,
		Email:     userEntity.Email,
		Curp:      userEntity.Curp,
		AccountID: userEntity.AccountID,
	}

	return userData, nil
}

func (u *usecase) GetSuperAdmins(ctx context.Context) ([]models.UserRequest, error) {
	// Obtener los super administradores desde el repositorio
	users, err := u.repo.GetSuperAdmins(ctx)
	if err != nil {
		return nil, err
	}

	// Formatear la respuesta
	var userData []models.UserRequest
	for _, user := range users {
		userData = append(userData, models.UserRequest{
			AccountID:  user.AccountID,
			ID:         user.ID,
			Name:       user.Name,
			Lastname1:  user.Lastname1,
			Lastname2:  user.Lastname2,
			Email:      user.Email,
			Curp:       user.Curp,
			Created_At: user.Created_At.Format("2006-01-02 15:04:05"),
		})
	}

	return userData, nil
}

func (u *usecase) GetDoctorByID(ctx context.Context, userID int) (models.DoctorRequest, error) {
	// Llamar al repositorio para obtener el doctor por ID
	doctorEntity, err := u.repo.GetDoctorByID(ctx, userID)
	if err != nil {
		return models.DoctorRequest{}, err
	}

	// Crear el objeto DoctorRequest con los datos del doctor
	doctorData := models.DoctorRequest{
		ID:             doctorEntity.ID,
		Name:           doctorEntity.Name,
		Lastname1:      doctorEntity.Lastname1,
		Lastname2:      doctorEntity.Lastname2,
		Email:          doctorEntity.Email,
		MedicalLicense: doctorEntity.MedicalLicense,
		SpecialtyID:    int(doctorEntity.SpecialtyID),
		AccountID:      doctorEntity.AccountID,
	}

	return doctorData, nil
}

func (u *usecase) GetDoctors(ctx context.Context) ([]models.DoctorRequest, error) {
	// Obtener los doctores desde el repositorio
	doctors, err := u.repo.GetDoctors(ctx)
	if err != nil {
		return nil, err
	}

	// Formatear la respuesta
	var doctorData []models.DoctorRequest
	for _, doctor := range doctors {
		doctorData = append(doctorData, models.DoctorRequest{
			ID:             doctor.ID,
			Name:           doctor.Name,
			Lastname1:      doctor.Lastname1,
			Lastname2:      doctor.Lastname2,
			Email:          doctor.Email,
			MedicalLicense: doctor.MedicalLicense,
			SpecialtyID:    int(doctor.SpecialtyID),
			AccountID:      doctor.AccountID,
		})
	}

	return doctorData, nil
}
