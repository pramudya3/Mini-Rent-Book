package services

import (
	"context"
	"rent-book/models"
	"rent-book/repositories"
)

type RentServiceInterface interface {
	NewRent(ctx context.Context, newRent models.NewRent, idToken int) error
	GetRentByLogin(ctx context.Context, idToken int) (models.Rent, error)
	GetAllRent(ctx context.Context) ([]models.Rent, error)
	UpdateRent(ctx context.Context, updateRent models.UpdateRent, idToken int) (models.UpdateRent, error)
}

type RentService struct {
	rentRepository repositories.RentRepositoryInterface
}

func NewRentService(rentRepo repositories.RentRepositoryInterface) RentServiceInterface {
	return &RentService{
		rentRepository: rentRepo,
	}
}

func (rs *RentService) NewRent(ctx context.Context, newRent models.NewRent, idToken int) error {
	err := rs.rentRepository.NewRent(ctx, newRent, idToken)
	return err
}

func (rs *RentService) GetRentByLogin(ctx context.Context, idToken int) (models.Rent, error) {
	rent, err := rs.rentRepository.GetRentByLogin(ctx, idToken)

	rentResponse := models.Rent{
		BookId:     rent.BookId,
		UserId:     rent.UserId,
		Title:      rent.Title,
		Author:     rent.Author,
		BorrowDate: rent.BorrowDate,
		ReturnMax:  rent.ReturnMax,
		ReturnDate: rent.ReturnDate,
	}
	return rentResponse, err
}

func (rs *RentService) GetAllRent(ctx context.Context) ([]models.Rent, error) {
	rents, err := rs.rentRepository.GetAllRent(ctx)
	return rents, err
}

func (rs *RentService) UpdateRent(ctx context.Context, updateRent models.UpdateRent, idToken int) (models.UpdateRent, error) {
	rent, err := rs.rentRepository.UpdateRent(ctx, updateRent, idToken)
	return rent, err
}
