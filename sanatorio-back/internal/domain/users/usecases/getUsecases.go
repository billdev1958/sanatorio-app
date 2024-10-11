package usecase

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"
)

func (u *usecase) GetDoctorByID(ctx context.Context, userID int) (models.DoctorRequest, error) {
	// Llamar al repositorio para obtener el doctor por ID
	doctorEntity, err := u.repo.GetDoctorByID(ctx, userID)
	if err != nil {
		return models.DoctorRequest{}, err
	}

	// Crear el objeto DoctorRequest con los datos del doctor
	doctorData := models.DoctorRequest{
		ID:             doctorEntity.ID,
		Name:           doctorEntity.FirstName,
		Lastname1:      doctorEntity.LastName1,
		Lastname2:      doctorEntity.LastName2,
		Email:          doctorEntity.Email,
		MedicalLicense: doctorEntity.MedicalLicense,
		SpecialtyID:    int(doctorEntity.SpecialtyID),
		AccountID:      doctorEntity.AccountID,
	}

	return doctorData, nil
}
