package repository

import (
	"context"
	"fmt"
	"log"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	storage *postgres.PgxStorage
}

func NewUserRepository(storage *postgres.PgxStorage) user.Repository {
	return &userRepository{storage: storage}
}

func (pr *userRepository) registerPatient(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, pt entities.PatientUser) error {
	pt.AccountID = accountID

	// Verificar que los campos obligatorios estén presentes
	if pt.AccountID == uuid.Nil || pt.Curp == "" {
		return fmt.Errorf("invalid input: missing required fields")
	}

	// Preparar la consulta para insertar el tipo de doctor
	query := "INSERT INTO patient (account_id, medical_history_id, first_name, last_name1, last_name2, curp, sex) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	// Ejecutar la consulta dentro de la transacción
	_, err := tx.Exec(ctx, query, accountID, pt.MedicalHistoryID, pt.FirstName, pt.LastName1, pt.LastName2, pt.Curp, pt.Sex)
	if err != nil {
		return fmt.Errorf("insert into patient table: %w", err)
	}

	return nil
}

func (pr *userRepository) registerAccount(ctx context.Context, tx pgx.Tx, ru entities.Account) (uuid.UUID, error) {

	log.Printf("Repository - Inserting account with ID: %s, dependency_id: %d, phone: %s, email: %s, role_id: %d", ru.ID, ru.AfiliationID, ru.PhoneNumber, ru.Email, ru.Rol)

	var accountID uuid.UUID
	query := "INSERT INTO account (id, dependency_id, phone, email, password, role_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := tx.QueryRow(ctx, query, ru.ID, ru.AfiliationID, ru.PhoneNumber, ru.Email, ru.Password, ru.Rol).Scan(&accountID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert account: %w", err)
	}
	return accountID, nil
}

func (pr *userRepository) registerMedicalHistory(ctx context.Context, tx pgx.Tx, medicalHistoryID string, pt entities.PatientUser) error {
	query := "INSERT INTO medical_history (id, medical_history_id, patient_name, lastname_1, lastname_2, curp, gender, status_md) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := tx.Exec(ctx, query, uuid.New(), medicalHistoryID, pt.FirstName, pt.LastName1, pt.LastName2, pt.Curp, pt.Sex, false)
	if err != nil {
		return fmt.Errorf("insert into medical_history table: %w", err)
	}
	return nil
}

func (pr *userRepository) registerMedicalHistoryB(ctx context.Context, tx pgx.Tx, medicalHistoryID string, bu entities.BeneficiaryUser) error {
	query := "INSERT INTO medical_history (id, medical_history_id, patient_name, curp, gender) VALUES ($1, $2, $3, $4, $5)"
	_, err := tx.Exec(ctx, query, uuid.New(), medicalHistoryID, bu.Firstname, bu.Lastname1, bu.Lastname2, bu.Curp, bu.Sex)
	if err != nil {
		return fmt.Errorf("insert into medical_history table: %w", err)
	}
	return nil
}

// TODO REFACTORIZAR ENTIDADES YA QUE AL REGISTRAR MEDICALHISTORY SE USAN DISTINTAS STRUCTS A PESAR DE QUE TIENEN LOS MISMOS PARAMETROS SEPARAR LOS CAMPOS QUE COINCIDEN Y HACER UNA NUEVA STRUCT
