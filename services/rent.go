package services

import (
	"context"
	"rent-book/models"
	"rent-book/repositories"
	"time"
)

type RentServiceInterface interface {
	NewRent(ctx context.Context, newRent models.NewRent, idToken int) error
	GetAllRent(ctx context.Context) ([]models.Rent, error)
	UpdateRent(ctx context.Context, updateRent models.UpdateRent, rentId int, idToken int) (models.UpdateRent, error)
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

func (rs *RentService) GetAllRent(ctx context.Context) ([]models.Rent, error) {
	rents, err := rs.rentRepository.GetAllRent(ctx)
	return rents, err
}

func (rs *RentService) UpdateRent(ctx context.Context, updateRent models.UpdateRent, rentId int, idToken int) (models.UpdateRent, error) {
	_, err := rs.rentRepository.UpdateRent(ctx, updateRent, rentId, idToken)

	responseUpdate := models.UpdateRent{
		ReturnDate: time.Now().Local(),
	}
	return responseUpdate, err
}
