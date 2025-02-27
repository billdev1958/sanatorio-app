package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/auth"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
	password "sanatorioApp/pkg/pass"
	"sanatorioApp/pkg/validation"
	"strconv"
)

func (u *usecase) LoginUser(ctx context.Context, request models.LoginUser) (models.Response, error) {
	// 1. Validar los datos de login
	if err := validation.ValidateLoginData(request.Email, request.Password); err != nil {
		return models.Response{
			Status:  "error",
			Message: "Validación fallida",
			Errors:  err.Error(),
		}, nil
	}

	// 2. Buscar el usuario por identificador (username o email)
	account, err := u.repo.GetUserByIdentifier(ctx, request.Email)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Usuario no encontrado",
			Errors:  err.Error(),
		}, nil
	}

	// 3. Verificar si está verificado
	if !account.IsVerified {
		return models.Response{
			Status:  "error",
			Message: "Usuario no verificado",
			Errors:  "El usuario debe verificar su cuenta antes de iniciar sesión",
		}, nil
	}

	// 4. Verificar la contraseña
	if !password.CheckPasswordHash(request.Password, account.Password) {
		return models.Response{
			Status:  "error",
			Message: "Contraseña incorrecta",
			Errors:  "La contraseña proporcionada no coincide",
		}, nil
	}

	// 5. Generar el token JWT usando account.ID y account.Rol
	token, err := auth.GenerateJWT(account.ID, int(account.Rol), account.IsVerified)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Error al generar el token",
			Errors:  err.Error(),
		}, nil
	}

	// 6. Crear la respuesta de login y retornarla dentro de Response
	loginResp := models.LoginResponse{
		AccountID: account.ID,
		Role:      int(account.Rol),
		Token:     token,
	}

	return models.Response{
		Status:  "success",
		Message: "Login exitoso",
		Data:    loginResp,
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
