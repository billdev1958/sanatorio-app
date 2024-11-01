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
	superAdminResponse, err := u.repo.RegisterAdminTransaction(ctx, registerAccount, registerSuperAdmin)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register patient: %w", err)
	}

	// Retornar los datos del paciente registrado
	return models.UserData{
		Name: superAdminResponse.FirstName,
	}, nil
}
