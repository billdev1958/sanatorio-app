package app

import (
	"context"
	"net/http"
	"sanatorioApp/internal/domain/email"
	v1 "sanatorioApp/internal/domain/email/http"
)

func EmailService(ctx context.Context, router *http.ServeMux, username, password, smtpHost string, smtpPort int) error {
	e := email.NewEmailService(username, password, smtpHost, smtpPort)
	h := v1.NewHandler(e)
	h.EmailRoutes(router)

	return nil
}
