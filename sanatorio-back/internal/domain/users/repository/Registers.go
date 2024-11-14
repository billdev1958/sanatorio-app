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

// Función para registrar un usuario
func (ur *userRepository) RegisterPatientTransaction(ctx context.Context, account entities.Account, pu entities.PatientUser) (entities.PatientUser, error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return pu, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctxTx); rbErr != nil {
				log.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			err = tx.Commit(ctxTx)
			if err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	// Registrar la cuenta utilizando el userID generado
	accountID, err := ur.registerAccount(ctxTx, tx, account)
	if err != nil {
		return pu, fmt.Errorf("failed to register account: %w", err)
	}

	err = ur.registerMedicalHistory(ctxTx, tx, pu.MedicalHistoryID, pu)
	if err != nil {
		return pu, fmt.Errorf("failed to register medical history: %w", err)
	}

	// Registrar al paciente en su tabla correspondiente (ignorar valores no necesarios)
	err = ur.registerPatient(ctxTx, tx, accountID, pu)
	if err != nil {
		return pu, fmt.Errorf("failed to register patient: %w", err)
	}

	return pu, nil
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
	query := "INSERT INTO medical_history (id, medical_history_id, patient_name, curp, gender) VALUES ($1, $2, $3, $4, $5)"
	_, err := tx.Exec(ctx, query, uuid.New(), medicalHistoryID, pt.FirstName, pt.Curp, pt.Sex)
	if err != nil {
		return fmt.Errorf("insert into medical_history table: %w", err)
	}
	return nil
}

func (pr *userRepository) registerMedicalHistoryB(ctx context.Context, tx pgx.Tx, medicalHistoryID string, bu entities.BeneficiaryUser) error {
	query := "INSERT INTO medical_history (id, medical_history_id, patient_name, curp, gender) VALUES ($1, $2, $3, $4, $5)"
	_, err := tx.Exec(ctx, query, uuid.New(), medicalHistoryID, bu.Firstname, bu.Curp, bu.Sex)
	if err != nil {
		return fmt.Errorf("insert into medical_history table: %w", err)
	}
	return nil
}

func (ur *userRepository) RegisterBeneficiary(ctx context.Context, bu entities.BeneficiaryUser) (message string, err error) {
	// Crear un contexto con cancelación
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Iniciar la transacción
	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return "", fmt.Errorf("begin transaction: %w", err)
	}

	// Definir el defer para manejo de commit/rollback
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctxTx); rbErr != nil {
				log.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			err = tx.Commit(ctxTx)
			if err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	err = ur.registerMedicalHistoryB(ctxTx, tx, bu.MedicalHistoryID, bu)
	if err != nil {
		return "", fmt.Errorf("failed to register medical history: %w", err)
	}

	query := "INSERT INTO beneficiary (id, account_holder, medical_history_id, first_name, last_name1, last_name2, curp, sex) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err = tx.Exec(ctxTx, query, bu.ID, bu.AccountHolder, bu.MedicalHistoryID, bu.Firstname, bu.Lastname1, bu.Lastname2, bu.Curp, bu.Sex)
	if err != nil {
		return "", fmt.Errorf("error al insertar en la tabla beneficiary: %w", err)
	}

	return "Beneficiario registrado exitosamente", nil
}

// TODO REFACTORIZAR ENTIDADES YA QUE AL REGISTRAR MEDICALHISTORY SE USAN DISTINTAS STRUCTS A PESAR DE QUE TIENEN LOS MISMOS PARAMETROS SEPARAR LOS CAMPOS QUE COINCIDEN Y HACER UNA NUEVA STRUCT
