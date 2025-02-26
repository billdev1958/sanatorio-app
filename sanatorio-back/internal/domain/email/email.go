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

// Interfaz del servicio de email
type EmailS interface {
	SendEmail(ctx context.Context, dd *model.DestinataryData) (bool, error)
	ConfirmAccount(ctx context.Context, cr models.ConfirmRequest) (bool, error)
}

// Implementación del servicio de email
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
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// Enviar email
func (e *EmailService) SendEmail(ctx context.Context, dd *model.DestinataryData) (bool, error) {
	// 🔹 Validar que el email y el token no estén vacíos
	if dd.Email == "" || dd.Token == "" {
		return false, fmt.Errorf("❌ Error: `Email` y `Token` son obligatorios para enviar el correo")
	}

	// 🔹 Asignar URL de confirmación
	dd.LinkConfirmacion = fmt.Sprintf("https://cms.ax01.dev/confirmation/%s", dd.Token)

	// 🔹 Cargar plantilla HTML
	tmpl, err := LoadTemplate("/app/email/plantillaConfirmacion.html")
	if err != nil {
		return false, fmt.Errorf("error cargando plantilla: %w", err)
	}

	// 🔹 Crear mensaje de correo
	m := mail.NewMsg()
	if err := m.From(e.Username); err != nil {
		return false, fmt.Errorf("falló al asignar remitente: %w", err)
	}

	m.To(dd.Email)
	m.Subject("Confirmación de cuenta")

	m.EmbedFile("/app/email/logo.png")
	m.EmbedFile("/app/email/logo_cms.png")

	// Asignar manualmente los CIDs en la plantilla
	dd.Logo1CID = "logo.png"
	dd.Logo2CID = "logo_cms.png"

	// 🔹 Establecer el cuerpo del correo con la plantilla HTML
	if err := m.SetBodyHTMLTemplate(tmpl, dd); err != nil {
		return false, fmt.Errorf("error ejecutando plantilla: %w", err)
	}

	// 🔹 Configurar cliente SMTP
	client, err := mail.NewClient(
		e.SmtpHost,
		mail.WithPort(e.SmtpPort),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(e.Username),
		mail.WithPassword(e.Password),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return false, fmt.Errorf("falló al crear cliente SMTP: %w", err)
	}

	// 🔹 Enviar correo
	if err := client.DialAndSend(m); err != nil {
		return false, fmt.Errorf("falló al enviar el correo: %w", err)
	}

	log.Printf("📧 Correo enviado con éxito a %s", dd.Email)
	return true, nil
}

// Confirmar cuenta a través del token
func (e *EmailService) ConfirmAccount(ctx context.Context, cr models.ConfirmRequest) (bool, error) {
	_, err := auth.ValidateJWTConfirmation(cr.Token)
	if err != nil {
		log.Printf("❌ Error validando token de confirmación: %v", err)
		return false, err
	}

	log.Println("✅ Cuenta confirmada con éxito")
	return true, nil
}
