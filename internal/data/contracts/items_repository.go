package contracts

import (
	"go-template-microservice-v2/internal/data/entities"

	uuid "github.com/satori/go.uuid"
)

type IItemsRepository interface {
	AddItems(itemsEntity entities.ItemsEntity) error
	GetItems(id uuid.UUID) (*entities.ItemsEntity, error)
	GetAllItems() ([]*entities.ItemsEntity, error)
	UpdateItems(itemsEntity entities.ItemsEntity) error
}
