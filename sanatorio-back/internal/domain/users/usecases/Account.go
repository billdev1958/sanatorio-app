package usecase

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/auth"
	"sanatorioApp/internal/domain/users/http/models"
	"sanatorioApp/pkg/generate"
	"time"
)

func (u *usecase) AccountVerification(ctx context.Context, token string) (models.Response, error) {
	claims, err := auth.ValidateJWTConfirmation(token)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Token de confirmación inválido",
			Errors:  err.Error(),
		}, nil
	}

	accountID := claims.AccountID

	updated, err := u.repo.AccountVerification(ctx, accountID, true)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Error al verificar la cuenta",
			Errors:  err.Error(),
		}, nil
	}

	if !updated {
		return models.Response{
			Status:  "error",
			Message: "No se pudo verificar la cuenta. No se actualizó ninguna fila.",
		}, nil
	}

	return models.Response{
		Status:  "success",
		Message: "La cuenta ha sido verificada exitosamente",
	}, nil
}

func (u *usecase) SendEmailVerification(ctx context.Context, email string) (string, error) {
	code, err := generate.GenerateCode(6)
	if err != nil {
		return "", fmt.Errorf("generated code verification failed: %w", err)
	}

	if email == "" {
		log.Printf("❌ Error: `u.email` esta vacio, el servicio de email no está inicializado")
		return "", fmt.Errorf("email service not initialized")
	}

	var expiredAt = time.Now().Add(2 * time.Hour)

	if err := u.repo.SaveCodeVerification(ctx, email, code, expiredAt); err != nil {
		log.Printf("❌ Error al guardar el código de verificación para %s: %v", email, err)
		return "", fmt.Errorf("save verification code failed")
	}

	if _, err := u.email.ForwardEmail(ctx, email, code); err != nil {
		log.Printf("❌ Error al enviar el correo a %s: %v", email, err)
		return "", fmt.Errorf("error sending confirmation email: %w", err)
	}

	return "Email enviado correctamente", nil
}

func (u *usecase) CodeVerification(ctx context.Context, cd models.ConfirmationData) (string, error) {
	verification, err := u.repo.VerifyCode(ctx, cd.Code, cd.Email)
	if err != nil {
		return "", fmt.Errorf("verify code from repository failed: %w", err)
	}

	if !verification {
		return "", fmt.Errorf("verification failed: invalid or expired code")
	}

	return "Verificación exitosa", nil
}
