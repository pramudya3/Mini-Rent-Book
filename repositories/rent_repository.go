package repositories

import (
	"context"
	"database/sql"
	"errors"
	"rent-book/models"
	"time"
)

type RentRepositoryInterface interface {
	NewRent(ctx context.Context, rentBook models.NewRent, idToken int) error
	GetRentByLogin(ctx context.Context, rentId int) (models.Rent, error)
	GetAllRent(ctx context.Context) ([]models.Rent, error)
	UpdateRent(ctx context.Context, updateRent models.UpdateRent, idToken int) (models.UpdateRent, error)
}

type RentRepository struct {
	mysql *sql.DB
}

func NewRentRepository(db *sql.DB) *RentRepository {
	return &RentRepository{
		mysql: db,
	}
}

func (rr *RentRepository) NewRent(ctx context.Context, rentBook models.NewRent, idToken int) error {
	query := "INSERT INTO rent(bookId, borrow_date, return_date, userId) VALUES (?, ?, ?, ?)"

	borrowDate := time.Now()
	timeBorrow := 72 * time.Hour
	returnMax := borrowDate.Add(timeBorrow)

	_, err := rr.mysql.ExecContext(ctx, query, rentBook.BookId, borrowDate, returnMax, idToken)
	if err != nil {
		return err
	}
	return nil
}

func (rr *RentRepository) GetRentByLogin(ctx context.Context, idToken int) (models.Rent, error) {
	var rent models.Rent
	query := "SELECT userId, bookId, title, author, borrow_date, return_max WHERE userId = ?"

	err := rr.mysql.QueryRowContext(ctx, query, idToken).Scan(&rent.UserId, &rent.BookId, &rent.Title, &rent.Author, &rent.BorrowDate, &rent.ReturnMax, idToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Rent{}, errors.New("Data not found")
		}
		return models.Rent{}, err
	}
	return rent, nil
}

func (rr *RentRepository) GetAllRent(ctx context.Context) ([]models.Rent, error) {
	query := "SELECT bookId, UserId, borrow_date, return_date from rent"

	rows, err := rr.mysql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rents []models.Rent
	for rows.Next() {
		var rent models.Rent
		err := rows.Scan(&rent.UserId, &rent.BookId, &rent.BorrowDate, &rent.ReturnDate)
		if err != nil {
			return nil, err
		}
		rents = append(rents, rent)
	}
	return rents, nil
}

func (rr *RentRepository) UpdateRent(ctx context.Context, updateRent models.UpdateRent, idToken int) (models.UpdateRent, error) {
	query := "UPDATE rent SET return_date = ? WHERE userId = ?"

	result, err := rr.mysql.ExecContext(ctx, query, updateRent.ReturnDate, idToken)
	if err != nil {
		return models.UpdateRent{}, err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return models.UpdateRent{}, errors.New("data not found")
	}

	return updateRent, nil
}
