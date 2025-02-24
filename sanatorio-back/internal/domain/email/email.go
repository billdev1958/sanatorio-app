package email

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"sanatorioApp/internal/domain/auth"
	"sanatorioApp/internal/domain/email/http/models"
	model "sanatorioApp/internal/domain/email/models"

	"github.com/wneessen/go-mail"
)

type EmailS interface {
	SendEmail(ctx context.Context, dd *model.DestinataryData) (bool, error)
	ConfirmAccount(ctx context.Context, cr models.ConfirmRequest) (bool, error)
}

type EmailService struct {
	Username string
	Password string
	SmtpHost string
	SmtpPort int
}

func NewEmailService(username, password, smtpHost string, smtPort int) *EmailService {
	return &EmailService{
		Username: username,
		Password: password,
		SmtpHost: smtpHost,
		SmtpPort: smtPort,
	}
}

func LoadTemplate(filePath string) (*template.Template, error) {
	template, err := template.ParseFiles(filePath)
	if err != nil {
		return nil, err
	}
	return template, nil
}

func (e *EmailService) SendEmail(ctx context.Context, dd *model.DestinataryData) (bool, error) {
	tmpl, err := LoadTemplate("./plantillaConfirmacion.html")
	if err != nil {
		return false, fmt.Errorf("error loading template: %w", err)
	}

	confirmationURL := fmt.Sprintf("https://cms.ax01.dev/v1/patient/confirm/%s", dd.Token)

	dd.LinkConfirmacion = confirmationURL

	m := mail.NewMsg()
	if err := m.From(e.Username); err != nil {
		return false, fmt.Errorf("failed to set From address: %w", err)
	}

	m.To(dd.Email)
	m.Subject("Confirmacion de cuenta")

	m.EmbedFile("./logo.png")
	m.EmbedFile("./logo_cms.png")

	dd.Logo1CID = "logo.png"
	dd.Logo2CID = "logo_cms.png"

	if err := m.SetBodyHTMLTemplate(tmpl, dd); err != nil {
		return false, fmt.Errorf("error executing template: %w", err)
	}

	client, err := mail.NewClient(
		e.SmtpHost,
		mail.WithPort(e.SmtpPort),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(e.Username),
		mail.WithPassword(e.Password),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return false, fmt.Errorf("failed to create mail client: %w", err)
	}

	if err := client.DialAndSend(m); err != nil {
		return false, fmt.Errorf("failed to send email: %w", err)
	}

	return true, nil
}

func (e *EmailService) ConfirmAccount(ctx context.Context, cr models.ConfirmRequest) (bool, error) {
	_, err := auth.ValidateJWTConfirmation(cr.Token)
	if err != nil {
		log.Printf("Error validando token de confirmación: %v", err)
		return false, err
	}

	log.Println("Cuenta confirmada con éxito")
	return true, nil
}
