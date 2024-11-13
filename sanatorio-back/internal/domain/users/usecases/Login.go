package usecase

import (
	"context"
	"errors"
	"fmt"
	"sanatorioApp/internal/auth"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
	validation "sanatorioApp/pkg"
	password "sanatorioApp/pkg/pass"
	"strconv"
)

func (u *usecase) LoginUser(ctx context.Context, request models.LoginUser) (models.LoginResponse, error) {

	// Validar los datos de login
	err := validation.ValidateLoginData(request.Email, request.Password)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("validación fallida: %w", err)
	}

	// Buscar el usuario por identificador (username o email)
	account, err := u.repo.GetUserByIdentifier(ctx, request.Email)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("usuario no encontrado: %w", err)
	}

	// Verificar la contraseña
	if !password.CheckPasswordHash(request.Password, account.Password) {
		return models.LoginResponse{}, errors.New("contraseña incorrecta")
	}

	token, err := auth.GenerateJWT(account.ID, int(account.Rol))
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("error al generar el token: %w", err)
	}

	// Crear y devolver la respuesta del login
	return models.LoginResponse{
		AccountID: account.ID,
		Role:      int(account.Rol),
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
