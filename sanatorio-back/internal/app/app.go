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
)

type App struct {
	DB     *pgxpool.Pool
	port   string
	router *http.ServeMux
}

func NewApp(port string) (*App, error) {
	dsn := "postgres://root:secret@db:5432/university_db?sslmode=disable"

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
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.port),
		Handler: loggingMiddleware,
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
