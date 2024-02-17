package usecase

import (
	"errors"
	"github.com/NickNaskida/Gorder/internal/models"
	interfaces "github.com/NickNaskida/Gorder/pkg/v1"
	"gorm.io/gorm"
)

type UseCase struct {
	repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UseCaseInterface {
	return &UseCase{repo}
}

// Create creates a new order in the database from supplied order argument
func (uc *UseCase) Create(order models.Order) (models.Order, error) {
	if _, err := uc.repo.Get(order.OrderId); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Order{}, errors.New("no such user with the id supplied")
	}
	return uc.repo.Create(order)
}

// Get retrieves an order from the database using the supplied id
func (uc *UseCase) Get(id string) (models.Order, error) {
	var order models.Order
	var err error

	if order, err = uc.repo.Get(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Order{}, errors.New("no such order with the id supplied")
		}
		return models.Order{}, err
	}

	return order, nil
}

// Update updates an order in the database using the supplied order argument
func (uc *UseCase) Update(updateOrder models.Order) (models.Order, error) {
	var updatedOrder models.Order
	var err error

	if _, err = uc.repo.Get(updateOrder.OrderId); err != nil {
		return models.Order{}, errors.New("no such order with the id supplied")
	}

	updatedOrder, err = uc.repo.Update(updateOrder)
	if err != nil {
		return models.Order{}, err
	}

	return updatedOrder, nil
}

// Delete deletes an order from the database using the supplied id
func (uc *UseCase) Delete(id string) error {
	var err error

	if _, err = uc.repo.Get(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no such order with the id supplied")
		}
		return err
	}

	err = uc.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
