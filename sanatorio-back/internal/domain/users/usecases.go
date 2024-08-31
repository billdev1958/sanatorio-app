package user

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"
)

type Usecase interface {
	RegisterUser(ctx context.Context, request models.RegisterUserRequest) (models.Response, error)
	RegisterDoctor(ctx context.Context, request models.RegisterDoctorRequest) (models.Response, error)
}
