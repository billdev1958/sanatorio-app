package usecase

import (
	"context"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"

	"github.com/google/uuid"
)

type usecase struct {
	repo user.Repository
}

func NewUsecase(repo user.Repository) user.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterUser(ctx context.Context, request models.RegisterUserRequest) (models.Response, error) {

	registerUser := entities.RegisterUser{
		User: entities.User{
			Name:      request.Name,
			Lastname1: request.Lastname1,
			Lastname2: request.Lastname2,
		},
		Account: entities.Account{
			AccountID: uuid.New(),
			Email:     request.Email,
			Password:  request.Password,
			Rol:       request.Rol,
		},
		DocumentID: request.Curp,
	}

	userResponse, err := u.repo.RegisterUserTransaction(ctx, registerUser)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to register user",
			Errors:  map[string]string{"register": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data: models.UserData{
			Name:  userResponse.Name,
			Email: userResponse.Email,
		},
	}, nil

}

func (u *usecase) RegisterDoctor(ctx context.Context, request models.RegisterDoctorRequest) (models.Response, error) {
	registerDoctor := entities.RegisterDoctor{
		User: entities.User{
			Name:      request.Name,
			Lastname1: request.Lastname1,
			Lastname2: request.Lastname2,
		},
		Account: entities.Account{
			AccountID: uuid.New(),
			Email:     request.Email,
			Password:  request.Password,
			Rol:       request.Rol,
		},
		DocumentID:  request.MedicalLicense,
		SpecialtyID: request.Specialty,
	}

	doctorResponse, err := u.repo.RegisterDoctorTransaction(ctx, registerDoctor)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to register user",
			Errors:  map[string]string{"register": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data: models.UserData{
			Name:  doctorResponse.Name,
			Email: doctorResponse.Email,
		},
	}, nil

}
