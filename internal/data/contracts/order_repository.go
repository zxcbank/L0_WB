package contracts

import (
	"go-template-microservice-v2/internal/data/entities"

	"github.com/google/uuid"
)

type IOrderRepository interface {
	AddOrder(orderEntity entities.OrderEntity) error
	GetOrder(id uuid.UUID) (*entities.OrderEntity, error)
	GetAllOrder() ([]*entities.OrderEntity, error)
	UpdateOrder(orderEntity entities.OrderEntity) error
}
