package services

import (
	"context"
	"errors"
	"rent-book/models"
	"rent-book/repositories"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, newUser models.NewUserRequest) error
	GetUserById(ctx context.Context, userId int) (models.UserResponse, error)
	GetAllUser(ctx context.Context) ([]models.UserResponse, error)
	DeleteUser(ctx context.Context, idToken int) error
	UpdateUser(ctx context.Context, updateUser models.NewUserRequest) (models.UserResponse, error)
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepo,
	}

}

func (us *UserService) CreateUser(ctx context.Context, newUser models.NewUserRequest) error {
	if newUser.Name == "" {
		return errors.New("Name is Required")
	}
	if newUser.Email == "" {
		return errors.New("Email is Required")
	}
	if newUser.Password == "" {
		return errors.New("Password is Required")
	}
	if newUser.PhoneNumber == "" {
		return errors.New("Phone Number is Required")
	}
	if newUser.Address == "" {
		return errors.New("Address is Required")
	}

	err := us.userRepository.CreateUser(ctx, newUser)
	return err
}

func (us *UserService) GetUserById(ctx context.Context, userId int) (models.UserResponse, error) {
	user, err := us.userRepository.GetUserById(ctx, userId)

	userResponse := models.UserResponse{
		UserId:      user.UserId,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return userResponse, err
}

func (us *UserService) GetAllUser(ctx context.Context) ([]models.UserResponse, error) {
	user, err := us.userRepository.GetAllUser(ctx)
	return user, err
}

func (us *UserService) DeleteUser(ctx context.Context, idToken int) error {
	err := us.userRepository.DeleteUser(ctx, idToken)
	return err
}

func (us *UserService) UpdateUser(ctx context.Context, updateUser models.NewUserRequest, idToken int) (models.UserResponse, error) {
	getUser, err := us.userRepository.GetUserById(ctx, idToken)
	if err != nil {
		return models.UserResponse{}, err
	}

	if updateUser.Name != "" {
		getUser.Name = updateUser.Name
	}

	if updateUser.Email != "" {
		getUser.Email = updateUser.Email
	}

	if updateUser.Password != "" {
		getUser.Password = updateUser.Password
	}

	if updateUser.PhoneNumber != "" {
		getUser.PhoneNumber = updateUser.PhoneNumber
	}

	if updateUser.Address != "" {
		getUser.Address = updateUser.Address
	}

	user, err := us.userRepository.UpdateUser(ctx, getUser, idToken)

	responseUpdate := models.UserResponse{
		UserId:      getUser.UserId,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
	return responseUpdate, err
}
