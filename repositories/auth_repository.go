package repositories

import (
	"context"
	"database/sql"
	"errors"
	"rent-book/models"
)

type AuthRepositoryInterface interface {
	Login(ctx context.Context, email string) (models.LoginDataResponse, error)
}

type AuthRepository struct {
	mysql *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		mysql: db,
	}
}

func (ar *AuthRepository) Login(ctx context.Context, email string) (models.LoginDataResponse, error) {
	query := "SELECT userId, name, email, password FROM users WHERE email = ?"

	var user models.LoginDataResponse
	err := ar.mysql.QueryRowContext(ctx, query, email).Scan(&user.UserId, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.LoginDataResponse{}, errors.New("data not found")
		}
		return models.LoginDataResponse{}, err
	}
	return user, err
}
