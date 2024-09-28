package repository

import (
	"context"
	"fmt"
	"log"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	postgres "sanatorioApp/internal/infraestructure/db"
)

type userRepository struct {
	storage *postgres.PgxStorage
}

func NewUserRepository(storage *postgres.PgxStorage) user.Repository {
	return &userRepository{storage: storage}
}

func (ur *userRepository) RegisterSuperUserTransaction(ctx context.Context, ad entities.AdminData, su entities.SuperUser) (entities.SuperUser, error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return su, fmt.Errorf("begin transaction: %w", err)
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

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctxTx, tx, ad.AccountID, ad.PasswordAdmin)
	if err != nil {
		return su, fmt.Errorf("failed to authenticate admin: %w", err)
	}
	if !isValid {
		return su, fmt.Errorf("failed to authenticate admin: invalid password")
	}

	// Registrar el usuario y obtener el userID generado
	userID, err := ur.registerUser(ctxTx, tx, su.User)
	if err != nil {
		return su, fmt.Errorf("failed to register user: %w", err)
	}

	// Registrar la cuenta utilizando el userID generado
	err = ur.registerAccount(ctxTx, tx, su.Account, userID)
	if err != nil {
		return su, fmt.Errorf("failed to register account: %w", err)
	}

	_, err = ur.registerSuperAdmin(ctxTx, tx, su)
	if err != nil {
		return su, fmt.Errorf("failed to register super_usuario: %w", err)
	}

	return su, nil
}

func (ur *userRepository) RegisterDoctorTransaction(ctx context.Context, ad entities.AdminData, du entities.DoctorUser) (entities.DoctorUser, error) {
	// Iniciar la transacción
	ctxTx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := ur.storage.DbPool.Begin(ctxTx)
	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return du, fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctxTx); rbErr != nil {
				log.Printf("error rolling back transaction: %v", rbErr)
			}
		} else {
			if err = tx.Commit(ctxTx); err != nil {
				log.Printf("error committing transaction: %v", err)
			}
		}
	}()

	// Verificar la contraseña del administrador antes de proceder
	isValid, err := ur.CheckAdminPassword(ctxTx, tx, ad.AccountID, ad.PasswordAdmin)
	if err != nil {
		return du, fmt.Errorf("failed to authenticate admin: %w", err)
	}
	if !isValid {
		return du, fmt.Errorf("authentication failed: invalid credentials")
	}

	// Registrar el usuario y obtener el userID generado
	userID, err := ur.registerUser(ctxTx, tx, du.User)
	if err != nil {
		return du, fmt.Errorf("failed to register doctor user: %w", err)
	}

	// Registrar la cuenta utilizando el userID generado
	err = ur.registerAccount(ctxTx, tx, du.Account, userID)
	if err != nil {
		return du, fmt.Errorf("failed to register account: %w", err)
	}

	err = ur.registerDoctor(ctxTx, tx, du)
	if err != nil {
		return du, fmt.Errorf("failed to register user type: %w", err)
	}

	return du, nil
}
