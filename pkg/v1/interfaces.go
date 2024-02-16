package v1

import "github.com/NickNaskida/Gorder/internal/models"

type RepoInterface interface {
	Create(models.Order) (models.Order, error)

	Get(id string) (models.Order, error)

	Update(models.Order) error

	Delete(id string) error
}

type UseCaseInterface interface {
	Create(order models.Order) (models.Order, error)

	Get(id string) (models.Order, error)

	Update(models.Order) error

	Delete(id string) error
}
