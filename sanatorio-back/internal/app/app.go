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

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DSN")
	port := os.Getenv("PORT")

	db, err := setupDatabase(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to setup database: %w", err)
	}

	if err := seedDatabase(db); err != nil {
		return nil, fmt.Errorf("failed to seed database: %w", err)
	}

	router := http.NewServeMux()

	return &App{
		DB:     db,
		port:   port,
		router: router,
	}, nil

}

func (app *App) Run() error {
	middleware.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	loggingMiddleware := middleware.NewLoggingMiddleware(app.router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization", "Accept", "Success", "OK"},
		AllowCredentials: true,
	})

	handler := c.Handler(loggingMiddleware)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.port),
		Handler: handler,
	}

	if err := StartService(context.Background(), app.DB, app.router); err != nil {
		return fmt.Errorf("failed to start user service: %w", err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-quit
		log.Println("server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		server.Shutdown(ctx)
		app.DB.Close()
	}()

	log.Printf("starting server on port %s...", app.port)
	return server.ListenAndServe()
}

func setupDatabase(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return dbPool, nil
}

func seedDatabase(dbPool *pgxpool.Pool) error {
	storage := postgres.NewPgxStorage(dbPool)

	err := storage.SeedRoles(context.Background())
	if err != nil {
		log.Fatalf("error seeding roles %v\n", err)
	}

	err = storage.SeedOfficeStatus(context.Background())
	if err != nil {
		log.Fatalf("error seeding office status %v\n", err)
	}

	err = storage.SeedAdminUser(context.Background())
	if err != nil {
		log.Fatalf("error seeding user  %v\n", err)
	}

	return nil
}
