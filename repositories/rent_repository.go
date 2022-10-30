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
	GetRentByLogin(ctx context.Context, idToken int) (models.Rent, error)
	GetAllRent(ctx context.Context) ([]models.Rent, error)
	UpdateRent(ctx context.Context, updateRent models.UpdateRent, rentId int, idToken int) (models.UpdateRent, error)
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
	query := "INSERT INTO rents(bookId, borrow_date, return_max, userId) VALUES (?, ?, ?, ?)"

	borrowDate := time.Now().Local()
	timeBorrow := 168 * time.Hour
	returnMax := borrowDate.Add(timeBorrow)

	_, err := rr.mysql.ExecContext(ctx, query, rentBook.BookId, borrowDate, returnMax, idToken)
	if err != nil {
		return err
	}
	return nil
}

func (rr *RentRepository) GetRentByLogin(ctx context.Context, idToken int) (models.Rent, error) {
	var rent models.Rent
	query := "SELECT rentId, rents.bookId, rents.borrow_date, rents.return_max, rents.return_date FROM rents WHERE userId = ?"

	err := rr.mysql.QueryRowContext(ctx, query, idToken).Scan(&rent.RentId, &rent.BookId, &rent.BorrowDate, &rent.ReturnMax, &rent.ReturnDate, idToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Rent{}, errors.New("Data not found")
		}
		return models.Rent{}, err
	}
	return rent, nil
}

func (rr *RentRepository) GetAllRent(ctx context.Context) ([]models.Rent, error) {
	query := "SELECT rentId, rents.userId, rents.bookId, rents.borrow_date, rents.return_max FROM rents"

	rows, err := rr.mysql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rents []models.Rent
	for rows.Next() {
		var rent models.Rent
		err := rows.Scan(&rent.RentId, &rent.UserId, &rent.BookId, &rent.BorrowDate, &rent.ReturnMax)
		if err != nil {
			return nil, err
		}
		rents = append(rents, rent)
	}
	return rents, nil
}

func (rr *RentRepository) UpdateRent(ctx context.Context, updateRent models.UpdateRent, rentId int, idToken int) (models.UpdateRent, error) {
	query := "UPDATE rents SET return_date = current_timestamp where rentId = ? AND userId =?"

	result, err := rr.mysql.ExecContext(ctx, query, rentId, idToken)
	if err != nil {
		return models.UpdateRent{}, err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return models.UpdateRent{}, errors.New("data not found")
	}

	return updateRent, nil
}
