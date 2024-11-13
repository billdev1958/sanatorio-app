package usecase

import (
	"context"
	"fmt"
	"log"
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
		ID:           uuid.New(), // Asignar un nuevo UUID
		AfiliationID: request.AfiliationID,
		Email:        request.Email,
		Password:     hashedPassword,
		Rol:          entities.Doctor,
	}

	// Crear la entidad PatientUser con los datos de la solicitud
	registerDoctor := entities.DoctorUser{
		MedicalLicense:   request.MedicalLicense,
		SpecialtyLicense: request.SpecialtyLicense,
		FirstName:        request.Name,
		LastName1:        request.Lastname1,
		LastName2:        request.Lastname2,
		Sex:              request.Sex,
	}

	// Intentar registrar al paciente en una transacción
	doctorResponse, err := u.repo.RegisterDoctorTransaction(ctx, registerAccount, registerDoctor)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register doctor: %w", err)
	}

	// Retornar los datos del paciente registrado
	return models.UserData{
		Name: doctorResponse.FirstName,
	}, nil
}

func (u *usecase) GetDoctorByID(ctx context.Context, userID int) (models.DoctorRequest, error) {
	// Llamar al repositorio para obtener el doctor por ID
	doctorEntity, err := u.repo.GetDoctorByID(ctx, userID)
	if err != nil {
		return models.DoctorRequest{}, err
	}

	// Crear el objeto DoctorRequest con los datos del doctor
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
