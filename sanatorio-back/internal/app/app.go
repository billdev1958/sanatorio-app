package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	postgres "sanatorioApp/internal/infraestructure/db"
	"sanatorioApp/middleware"
	"strconv"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type App struct {
	DB     *pgxpool.Pool
	port   string
	router *http.ServeMux
}

func NewApp() (*App, error) {

	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		// Si ocurre un error al cargar el archivo .env, se detiene el programa y se imprime el error
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtener la DSN (cadena de conexión de la base de datos) y el puerto desde las variables de entorno
	dsn := os.Getenv("DSN")
	port := os.Getenv("PORT")

	// Configurar la base de datos utilizando la cadena DSN
	db, err := setupDatabase(dsn)
	if err != nil {
		// Si ocurre un error al configurar la base de datos, se devuelve un error
		return nil, fmt.Errorf("failed to setup database: %w", err)
	}

	// Sembrar datos iniciales en la base de datos
	if err := seedDatabase(db); err != nil {
		// Si ocurre un error al sembrar la base de datos, se devuelve un error
		return nil, fmt.Errorf("failed to seed database: %w", err)
	}

	// Crear un nuevo router para manejar las rutas HTTP
	router := http.NewServeMux()

	// Devolver una instancia de la aplicación con la base de datos configurada, el puerto y el router
	return &App{
		DB:     db,     // La conexión a la base de datos
		port:   port,   // El puerto en el que correrá el servidor
		router: router, // El enrutador HTTP
	}, nil
}

func (app *App) Run() error {

	// Configuración del logger para que use JSON en la salida estándar
	middleware.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Inicializa el middleware de logging que envuelve el enrutador
	loggingMiddleware := middleware.NewLoggingMiddleware(app.router)

	// Configuración de CORS para permitir peticiones desde cualquier origen
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                                                                                      // Permitir cualquier origen
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodOptions, http.MethodPut}, // Métodos permitidos
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization", "Accept", "Success", "OK"},                                     // Headers permitidos
		AllowCredentials: true,                                                                                                               // Permitir credenciales (cookies, autenticación HTTP)
	})

	// Aplica el middleware de CORS sobre el middleware de logging
	handler := c.Handler(loggingMiddleware)

	// Configuración del servidor HTTP
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.port), // Configura la dirección del servidor en base al puerto
		Handler: handler,                      // Usa el handler con CORS y logging aplicados
	}

	// variables para servicio de email
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	sPort, err := strconv.Atoi(smtpPort)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Inicia el servicio principal de la aplicación (base de datos, enrutador, etc.)
	if err := UserService(context.Background(), app.DB, app.router); err != nil {
		return fmt.Errorf("failed to start user service: %w", err) // Devuelve un error si no se pudo iniciar el servicio
	}

	if err := CitesService(context.Background(), app.DB, app.router); err != nil {
		return fmt.Errorf("failed to start user service: %w", err) // Devuelve un error si no se pudo iniciar el servicio
	}

	if err := ScheduleService(context.Background(), app.DB, app.router); err != nil {
		return fmt.Errorf("failed to start schedule service: %w", err) // Devuelve un error si no se pudo iniciar el servicio
	}

	if err := AppointmentService(context.Background(), app.DB, app.router); err != nil {
		return fmt.Errorf("failed to start appointment service: %w", err) // Devuelve un error si no se pudo iniciar el servicio
	}

	if err := EmailService(context.Background(), app.router, smtpUsername, smtpPassword, smtpHost, sPort); err != nil {
		return fmt.Errorf("failed to start user service: %w", err) // Devuelve un error si no se pudo iniciar el servicio
	}

	// Canal para escuchar señales del sistema (SIGTERM, SIGINT)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // Notifica al canal cuando el sistema recibe la señal de apagado

	// Goroutine que espera la señal de apagado
	go func() {
		<-quit                                    // Espera a que el canal reciba la señal
		log.Println("server is shutting down...") // Imprime un mensaje de que el servidor se está apagando

		// Contexto con un timeout de 10 segundos para cerrar el servidor
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() // Asegura la cancelación del contexto

		// Apaga el servidor de manera segura con un timeout
		server.Shutdown(ctx)
		app.DB.Close() // Cierra la conexión a la base de datos
	}()

	// Imprime en los logs que el servidor está iniciando en el puerto especificado
	log.Printf("starting server on port %s...", app.port)

	// Inicia el servidor HTTP y escucha peticiones
	return server.ListenAndServe()
}

// funcion para inicializar la db
func setupDatabase(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return dbPool, nil
}

// funcion para generar registros a la db
func seedDatabase(dbPool *pgxpool.Pool) error {
	storage := postgres.NewPgxStorage(dbPool)

	err := storage.SeedRoles(context.Background())
	if err != nil {
		log.Fatalf("error seeding roles %v\n", err)
	}

	err = storage.SeedDependencies(context.Background())
	if err != nil {
		log.Fatalf("error seeding dependencies  %v\n", err)
	}

	err = storage.SeedMedicalInstitution(context.Background())
	if err != nil {
		log.Fatalf("error seeding cat_medical_institutions %v\n", err)
	}

	err = storage.SeedOfficeStatus(context.Background())
	if err != nil {
		log.Fatalf("error seeding office status %v\n", err)
	}

	err = storage.SeedOffices(context.Background())
	if err != nil {
		log.Fatalf("error seeding user  %v\n", err)
	}

	err = storage.SeedServices(context.Background())
	if err != nil {
		log.Fatalf("error seeding user  %v\n", err)
	}

	err = storage.SeedPermissions(context.Background())
	if err != nil {
		log.Fatalf("error seeding user  %v\n", err)
	}

	err = storage.SeedRolePermissions(context.Background())
	if err != nil {
		log.Fatalf("error seeding user  %v\n", err)
	}

	err = storage.SeedAppointmentStatus(context.Background())
	if err != nil {
		log.Fatalf("error seeding user  %v\n", err)
	}

	err = storage.SeedAdminUser(context.Background())
	if err != nil {
		log.Fatalf("error seeding user  %v\n", err)
	}

	err = storage.SeedShifts(context.Background())
	if err != nil {
		log.Fatalf("error seeding days  %v\n", err)
	}

	err = storage.SeedDays(context.Background())
	if err != nil {
		log.Fatalf("error seeding days  %v\n", err)
	}
	return nil

}
