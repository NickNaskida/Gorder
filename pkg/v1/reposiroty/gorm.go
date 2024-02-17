package reposiroty

import (
	"github.com/NickNaskida/Gorder/internal/models"
	interfaces "github.com/NickNaskida/Gorder/pkg/v1"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

// Create a new order in the database
func (repo *Repo) Create(order models.Order) (models.Order, error) {
	err := repo.db.Create(&order).Error
	if err != nil {
		return models.Order{}, err
	}

	return order, err
}

// Get orders from the database
func (repo *Repo) Get(OrderID string) (models.Order, error) {
	var order models.Order
	err := repo.db.Where("OrderId = ?", OrderID).First(&order).Error
	if err != nil {
		return models.Order{}, err
	}

	return order, err
}

// Update an order in the database
func (repo *Repo) Update(order models.Order) (models.Order, error) {
	err := repo.db.Save(&order).Error
	if err != nil {
		return models.Order{}, err
	}

	return order, err
}

// Delete an order from the database
func (repo *Repo) Delete(id string) error {
	var order models.Order
	return repo.db.Where("OrderId = ?", id).Delete(&order).Error
}
