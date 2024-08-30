package usecase

import (
	"context"
	user "sanatorioApp/internal/domain/users"
)

type usecase struct {
	repo user.Repository
}

func NewUsecase(repo user.Repository) user.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterUser(ctx context.Context) (string, error) {
	// Llamar al método RegisterUser del repositorio que simplemente devuelve "Hola Mundo"
	message, err := u.repo.RegisterUser(ctx)
	if err != nil {
		return "", err
	}
	return message, nil
}
