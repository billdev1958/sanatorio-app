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

func (u *usecase) RegisterReceptionist(ctx context.Context, request models.RegisterReceptionistRequest) (models.UserData, error) {

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
		Rol:          entities.Receptionist,
		IsVerified:   true,
	}

	// Crear la entidad PatientUser con los datos de la solicitud
	registerReceptionist := entities.ReceptionistUser{
		FirstName: request.Name,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,
		Curp:      request.Curp, // Asignar el CURP al paciente
		Sex:       request.Sex,
	}

	// Intentar registrar al paciente en una transacción
	receptionistResponse, err := u.repo.RegisterReceptionistTransaction(ctx, registerAccount, registerReceptionist)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register patient: %w", err)
	}

	// Retornar los datos del paciente registrado
	return models.UserData{
		Name: receptionistResponse.FirstName,
	}, nil
}

func (u *usecase) UpdatedReceptionist(ctx context.Context, request models.UpdateUser) (message string, err error) {
	update := entities.ReceptionistUser{
		AccountID: request.AccountID,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,
		Curp:      request.Curp,
		Sex:       request.Sex,
	}

	if request.Sex == catalogs.Male || request.Sex == catalogs.Female {
		update.Sex = request.Sex
	}

	message, err = u.repo.UpdateReceptionist(ctx, update)
	if err != nil {
		log.Printf("Failed to update receptionist with account_id: %s. Error: %v", request.AccountID, err)
		return "", fmt.Errorf("failed to update receptionist with account_id %s: %w", request.AccountID, err)
	}

	log.Printf("Successfully updated receptionist with account_id: %s", request.AccountID)
	return message, nil
}

func (u *usecase) SoftDeleteReceptionist(ctx context.Context, accountID uuid.UUID) (message string, err error) {
	delete := entities.Account{
		ID: accountID,
	}

	_, err = u.repo.SoftDeleteUserReceptionist(ctx, delete)
	if err != nil {
		log.Printf("Failed to delete receptionist with account_id: %s. Error: %v", accountID, err)
		return "", fmt.Errorf("failed to delete receptionist with account_id %s: %w", accountID, err)
	}

	log.Printf("Successfully delete receptionist with account_id: %s", accountID)
	return message, nil
}
