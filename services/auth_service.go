package services

import (
	"context"
	"errors"
	"rent-book/middlewares"
	"rent-book/models"
	"rent-book/repositories"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, userLogin models.LoginRequest) (models.LoginResponse, error)
}

type AuthService struct {
	authRepository repositories.AuthRepositoryInterface
}

func NewAuthService(authRepo repositories.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, userLogin models.LoginRequest) (models.LoginResponse, error) {

	if userLogin.Email == "" {
		return models.LoginResponse{}, errors.New("Email is Required")
	}

	user, err := as.authRepository.Login(ctx, userLogin.Email)
	if err != nil {
		return models.LoginResponse{}, err
	}

	token, err := middlewares.CreateToken(user.UserId)
	if err != nil {
		return models.LoginResponse{}, err
	}

	loginResponse := models.LoginResponse{
		Token:  token,
		UserId: user.UserId,
	}

	return loginResponse, err
}
