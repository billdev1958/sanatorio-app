package usecase

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/auth"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
	password "sanatorioApp/pkg/pass"
	"strconv"

	"math/rand"

	"github.com/google/uuid"
)

type usecase struct {
	repo user.Repository
}

func NewUsecase(repo user.Repository) user.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterPatient(ctx context.Context, request models.RegisterPatientRequest) (models.UserData, error) {

	log.Printf("Usecase - Received AfiliationID: %d", request.AfiliationID)

	// Hashear la contraseña del nuevo paciente
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to hash password: %w", err)
	}

	patientMedicalHistory := patient{
		FirstName: request.Name,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,
		LegacyID:  rand.Intn(900000) + 100000,
	}

	medicalHistoryID, err := createMedicalHistoryID(patientMedicalHistory)
	if err != nil {
		log.Printf("Error creating medical history ID: %v", err)
		return models.UserData{}, err // Manejar el error devolviendo un valor vacío o adecuado
	}

	registerAccount := entities.Account{
		ID:           uuid.New(), // Asignar un nuevo UUID
		AfiliationID: request.AfiliationID,
		Email:        request.Email,
		Password:     hashedPassword,
		Rol:          entities.Patient,
	}

	// Crear la entidad PatientUser con los datos de la solicitud
	registerPatient := entities.PatientUser{
		MedicalHistoryID: medicalHistoryID,
		FirstName:        request.Name,
		LastName1:        request.Lastname1,
		LastName2:        request.Lastname2,
		Curp:             request.Curp, // Asignar el CURP al paciente
		Sex:              request.Sex,
	}

	// Intentar registrar al paciente en una transacción
	patientResponse, err := u.repo.RegisterPatientTransaction(ctx, registerAccount, registerPatient)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register patient: %w", err)
	}

	// Retornar los datos del paciente registrado
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

func (u *usecase) LoginUser(ctx context.Context, request models.LoginUser) (models.LoginResponse, error) {
	// Crear la entidad de login
	loginUser := entities.Account{
		Email:    request.Email,
		Password: request.Password,
	}

	// Llamar al repositorio para autenticar el usuario
	loginResponse, err := u.repo.LoginUser(ctx, loginUser)
	if err != nil {
		return models.LoginResponse{}, err
	}

	// Generar el token JWT si el login fue exitoso
	token, err := auth.GenerateJWT(loginResponse.ID, int(loginResponse.Rol))
	if err != nil {
		return models.LoginResponse{}, err
	}

	// Retornar los datos crudos (LoginResponse) al handler
	return models.LoginResponse{
		AccountID: loginResponse.ID,
		Role:      int(loginResponse.Rol),
		Token:     token,
	}, nil
}

func getInitials(s string) string {
	return string(s[0])
}

type patient struct {
	LegacyID  int
	FirstName string
	LastName1 string
	LastName2 string
	Sex       string
	Rol       int
}

func createMedicalHistoryID(p patient) (medicalHistoryID string, err error) {
	nameDigit := getInitials(p.FirstName)
	last1Digit := getInitials(p.LastName1)
	last2Digit := getInitials(p.LastName2)
	rol := entities.Patient

	medicalHistoryID = nameDigit + last1Digit + last2Digit + p.Sex + strconv.Itoa(rol) + fmt.Sprintf("-%06d", p.LegacyID)

	return medicalHistoryID, nil
}
