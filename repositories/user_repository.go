package repositories

import (
	"context"
	"database/sql"
	"errors"
	"rent-book/models"
	"time"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, newUser models.NewUserRequest) error
	GetUserById(ctx context.Context, userId int) (models.User, error)
	GetAllUser(ctx context.Context) ([]models.UserResponse, error)
	DeleteUser(ctx context.Context, idToken int) error
	UpdateUser(ctx context.Context, updateUser models.User, idToken int) (models.User, error)
}

type UserRepository struct {
	mysql *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		mysql: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, newUser models.NewUserRequest) error {
	query := "INSERT INTO users(name, email, password, phone_number, address, created_at) VALUES (?, ?, ?, ?, ?)"

	_, err := ur.mysql.ExecContext(ctx, query, newUser.Name, newUser.Email, newUser.Password, newUser.PhoneNumber, newUser.Address, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserById(ctx context.Context, userId int) (models.User, error) {
	var user models.User
	query := "SELECT userId, name, email, password, phone_number, address, created_at, updated_at FROM users WHERE userId = ?"

	err := ur.mysql.QueryRowContext(ctx, query, userId).Scan(&user.UserId, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("data not found")
		}
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetAllUser(ctx context.Context) ([]models.UserResponse, error) {
	query := "SELECT userId, name, email, phone_number, address, created_at, updated_at FROM users"

	rows, err := ur.mysql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.PhoneNumber, &user.Address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, idToken int) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := ur.mysql.ExecContext(ctx, query, idToken)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, updateUser models.User, idToken int) (models.User, error) {
	query := "UPDATE users SET name = ?, email = ?, password = ?, phone_number = ?, address = ?, updated_at = ? WHERE userId = ?"

	result, err := ur.mysql.ExecContext(ctx, query, updateUser.Name, updateUser.Email, updateUser.Password, updateUser.PhoneNumber, updateUser.Address, updateUser.UpdatedAt, updateUser.UserId)
	if err != nil {
		return models.User{}, err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return models.User{}, errors.New("data not found")
	}

	return updateUser, nil
}
