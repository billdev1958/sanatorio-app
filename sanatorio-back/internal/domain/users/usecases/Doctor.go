package usecase

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/catalogs"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
	password "sanatorioApp/pkg/pass"

	"github.com/google/uuid"
)

func (u *usecase) RegisterDoctor(ctx context.Context, request models.RegisterDoctorRequest) (models.UserData, error) {

	log.Printf("Usecase - Received AfiliationID: %d", request.AfiliationID)

	// Hashear la contraseña del nuevo paciente
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to hash password: %w", err)
	}

	registerAccount := entities.Account{
		ID:           uuid.New(),
		AfiliationID: request.AfiliationID,
		Email:        request.Email,
		Password:     hashedPassword,
		Rol:          entities.Doctor,
	}

	registerDoctor := entities.DoctorUser{
		MedicalLicense:   request.MedicalLicense,
		SpecialtyLicense: request.SpecialtyLicense,
		FirstName:        request.Name,
		LastName1:        request.Lastname1,
		LastName2:        request.Lastname2,
		Sex:              request.Sex,
	}

	doctorResponse, err := u.repo.RegisterDoctorTransaction(ctx, registerAccount, registerDoctor)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register doctor: %w", err)
	}

	return models.UserData{
		Name: doctorResponse.FirstName,
	}, nil
}

func (u *usecase) GetDoctorByID(ctx context.Context, userID int) (models.DoctorRequest, error) {
	doctorEntity, err := u.repo.GetDoctorByID(ctx, userID)
	if err != nil {
		return models.DoctorRequest{}, err
	}

	doctorData := models.DoctorRequest{
		ID:               doctorEntity.AccountID,
		Name:             doctorEntity.FirstName,
		Lastname1:        doctorEntity.LastName1,
		Lastname2:        doctorEntity.LastName2,
		MedicalLicense:   doctorEntity.MedicalLicense,
		SpecialtyLicense: doctorEntity.SpecialtyLicense,
		AccountID:        doctorEntity.AccountID,
	}

	return doctorData, nil
}

func (u *usecase) UpdatedDoctor(ctx context.Context, request models.DoctorUpdateRequest) (message string, err error) {
	update := entities.DoctorUser{
		AccountID:      request.AccountID,
		MedicalLicense: request.MedicalLicense,
		FirstName:      request.Firstname,
		LastName1:      request.Lastname1,
		LastName2:      request.Lastname2,
	}

	if request.Sex == catalogs.Male || request.Sex == catalogs.Female {
		update.Sex = request.Sex
	}

	message, err = u.repo.UpdateDoctor(ctx, update)
	if err != nil {
		log.Printf("Failed to update doctor with account_id: %s. Error: %v", request.AccountID, err)
		return "", fmt.Errorf("failed to update doctor with account_id %s: %w", request.AccountID, err)
	}

	log.Printf("Successfully updated doctor with account_id: %s", request.AccountID)
	return message, nil
}

func (u *usecase) SoftDeleteDoctor(ctx context.Context, accountID uuid.UUID) (message string, err error) {
	delete := entities.Account{
		ID: accountID,
	}

	_, err = u.repo.SoftDeleteUserDoctor(ctx, delete)
	if err != nil {
		log.Printf("Failed to delete doctor with account_id: %s. Error: %v", accountID, err)
		return "", fmt.Errorf("failed to delete doctor with account_id %s: %w", accountID, err)
	}

	log.Printf("Successfully delete doctor with account_id: %s", accountID)
	return message, nil
}
