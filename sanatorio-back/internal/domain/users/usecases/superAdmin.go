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

func (u *usecase) RegisterSuperAdmin(ctx context.Context, request models.RegisterSuperAdminRequest) (models.UserData, error) {

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
		Rol:          entities.SuperAdmin,
		IsVerified:   true,
	}

	// Crear la entidad PatientUser con los datos de la solicitud
	registerSuperAdmin := entities.SuperAdminUser{
		FirstName: request.Name,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,
		Curp:      request.Curp, // Asignar el CURP al paciente
		Sex:       request.Sex,
	}

	// Intentar registrar al paciente en una transacción
	superAdminResponse, err := u.repo.RegisterSuperAdminTransaction(ctx, registerAccount, registerSuperAdmin)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register patient: %w", err)
	}

	// Retornar los datos del paciente registrado
	return models.UserData{
		Name: superAdminResponse.FirstName,
	}, nil
}

func (u *usecase) UpdatedSuperAdmin(ctx context.Context, request models.UpdateUser) (message string, err error) {
	update := entities.SuperAdminUser{
		AccountID: request.AccountID,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,
		Curp:      request.Curp,
		Sex:       request.Sex,
	}

	if request.Sex == catalogs.Male || request.Sex == catalogs.Female {
		update.Sex = request.Sex
	}

	message, err = u.repo.UpdateSuperAdmin(ctx, update)
	if err != nil {
		log.Printf("Failed to update super_admin with account_id: %s. Error: %v", request.AccountID, err)
		return "", fmt.Errorf("failed to update super_admin with account_id %s: %w", request.AccountID, err)
	}

	log.Printf("Successfully updated super_admin with account_id: %s", request.AccountID)
	return message, nil
}

func (u *usecase) SoftDeleteSuperAdmin(ctx context.Context, accountID uuid.UUID) (message string, err error) {
	delete := entities.Account{
		ID: accountID,
	}

	_, err = u.repo.SoftDeleteUserSuperAdmin(ctx, delete)
	if err != nil {
		log.Printf("Failed to delete super_admin with account_id: %s. Error: %v", accountID, err)
		return "", fmt.Errorf("failed to delete admin with account_id %s: %w", accountID, err)
	}

	log.Printf("Successfully delete super_admin with account_id: %s", accountID)
	return message, nil
}
