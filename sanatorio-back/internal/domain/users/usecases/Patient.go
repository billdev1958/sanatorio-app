package usecase

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sanatorioApp/internal/domain/auth"
	"sanatorioApp/internal/domain/catalogs"
	"sanatorioApp/internal/domain/email"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
	password "sanatorioApp/pkg/pass"

	model "sanatorioApp/internal/domain/email/models"

	"github.com/google/uuid"
)

type usecase struct {
	repo  user.Repository
	email email.EmailS
}

func NewUsecase(repo user.Repository) user.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterPatient(ctx context.Context, request models.RegisterPatientRequest) (models.UserData, error) {
	log.Printf("Usecase - Received AfiliationID: %d", request.AfiliationID)

	// 🔹 Validar datos de entrada
	if request.Email == "" || request.Password == "" {
		return models.UserData{}, fmt.Errorf("❌ Error: Email y Password son obligatorios")
	}

	// 🔹 Hashear la contraseña del nuevo paciente
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// 🔹 Crear ID único para la historia clínica del paciente
	patientMedicalHistory := patient{
		FirstName: request.Name,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,
		LegacyID:  rand.Intn(900000) + 100000,
	}

	medicalHistoryID, err := createMedicalHistoryID(patientMedicalHistory)
	if err != nil {
		log.Printf("❌ Error creando medicalHistoryID: %v", err)
		return models.UserData{}, fmt.Errorf("error creating medical history ID: %w", err)
	}

	// 🔹 Crear la cuenta del usuario
	registerAccount := entities.Account{
		ID:           uuid.New(), // Generar UUID
		AfiliationID: request.AfiliationID,
		Email:        request.Email,
		Password:     hashedPassword,
		Rol:          entities.Patient,
		IsVerified:   false,
	}

	// 🔹 Crear la entidad de paciente
	registerPatient := entities.PatientUser{
		MedicalHistoryID: medicalHistoryID,
		FirstName:        request.Name,
		LastName1:        request.Lastname1,
		LastName2:        request.Lastname2,
		Curp:             request.Curp,
		Sex:              request.Sex,
	}

	// 🔹 Intentar registrar al paciente en la base de datos
	patientResponse, err := u.repo.RegisterPatientTransaction(ctx, registerAccount, registerPatient)
	if err != nil {
		log.Printf("❌ Error registrando paciente: %v", err)
		return models.UserData{}, fmt.Errorf("failed to register patient: %w", err)
	}

	// 🔹 Generar token de confirmación
	token, err := auth.GenerateJWTConfirmation(registerAccount.ID)
	if err != nil {
		log.Printf("❌ Error generando token: %v", err)
		return models.UserData{}, fmt.Errorf("error al generar el token: %w", err)
	}

	// 🔹 Concatenar el nombre completo del usuario
	fullName := fmt.Sprintf("%s %s %s", registerPatient.FirstName, registerPatient.LastName1, registerPatient.LastName2)

	// 🔹 Preparar datos para el correo
	dd := model.DestinataryData{
		FullName: fullName,
		Email:    registerAccount.Email,
		Token:    token,
	}

	// 🔹 Enviar email de confirmación
	if u.email == nil {
		log.Printf("❌ Error: `u.email` es nil, el servicio de email no está inicializado")
		return models.UserData{}, fmt.Errorf("email service not initialized")
	}

	if _, err := u.email.SendEmail(ctx, &dd); err != nil {
		log.Printf("❌ Error al enviar el correo a %s: %v", dd.Email, err)
		return models.UserData{}, fmt.Errorf("error sending confirmation email: %w", err)
	}

	// 🔹 Retornar los datos del paciente registrado
	return models.UserData{
		Name: patientResponse.FirstName,
	}, nil
}

func (u *usecase) RegisterBeneficiary(ctx context.Context, request models.RegisterBeneficiaryRequest) (message string, err error) {
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return "", fmt.Errorf("unauthorized: no claims found in context")
	}

	beneficiaryMedicalHistory := patient{
		FirstName: request.Firstname,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,
		LegacyID:  rand.Intn(900000) + 100000,
	}

	medicalHistoryID, err := createMedicalHistoryID(beneficiaryMedicalHistory)
	if err != nil {
		log.Printf("Error creating medical history ID: %v", err)
		return "models.UserData{}", err // Manejar el error devolviendo un valor vacío o adecuado
	}

	registerBeneficiary := entities.BeneficiaryUser{
		ID:               uuid.New(),
		AccountHolder:    claims.AccountID,
		MedicalHistoryID: medicalHistoryID,
		Firstname:        request.Firstname,
		Lastname1:        request.Lastname1,
		Lastname2:        request.Lastname2,
	}

	message, err = u.repo.RegisterBeneficiary(ctx, registerBeneficiary)
	if err != nil {
		log.Printf("Error registering beneficiary: %v", err)
		return "", fmt.Errorf("failed to register beneficiary: %w", err)
	}

	return message, nil
}

func (u *usecase) UpdatedPatient(ctx context.Context, request models.UpdateUser) (message string, err error) {
	update := entities.PatientUser{
		AccountID: request.AccountID,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,
		Curp:      request.Curp,
		Sex:       request.Sex,
	}

	if request.Sex == catalogs.Male || request.Sex == catalogs.Female {
		update.Sex = request.Sex
	}

	message, err = u.repo.UpdatePatient(ctx, update)
	if err != nil {
		log.Printf("Failed to update patient with account_id: %s. Error: %v", request.AccountID, err)
		return "", fmt.Errorf("failed to update patient with account_id %s: %w", request.AccountID, err)
	}

	log.Printf("Successfully updated patient with account_id: %s", request.AccountID)
	return message, nil
}

func (u *usecase) SoftDeletePatient(ctx context.Context, accountID uuid.UUID) (message string, err error) {
	delete := entities.Account{
		ID: accountID,
	}

	_, err = u.repo.SoftDeleteUserPatient(ctx, delete)
	if err != nil {
		log.Printf("Failed to delete patient with account_id: %s. Error: %v", accountID, err)
		return "", fmt.Errorf("failed to delete patient with account_id %s: %w", accountID, err)
	}

	log.Printf("Successfully delete patient with account_id: %s", accountID)
	return message, nil
}
