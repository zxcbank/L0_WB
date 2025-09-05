package repositories

import (
	"fmt"
	"go-template-microservice-v2/internal/data/contracts"
	"go-template-microservice-v2/internal/data/entities"
	gormpg "go-template-microservice-v2/pkg/gorm_pg"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type PgOrderRepository struct {
	PgGorm *gormpg.PgGorm
}

func NewPgOrderRepository(pgGorm *gormpg.PgGorm) contracts.IOrderRepository {
	return &PgOrderRepository{PgGorm: pgGorm}
}

func (p PgOrderRepository) AddOrder(OrderEntity entities.OrderEntity) error {
	err := p.PgGorm.DB.Create(OrderEntity).Error
	if err != nil {
		return errors.Wrap(err, "error in the inserting Order into the database.")
	}

	return nil
}

func (p PgOrderRepository) GetOrder(id uuid.UUID) (*entities.OrderEntity, error) {
	var Order entities.OrderEntity

	if err := p.PgGorm.DB.First(&Order, id).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the Order with id %s into the database.", id))
	}

	return &Order, nil
}

func (p PgOrderRepository) GetAllOrder() ([]*entities.OrderEntity, error) {
	var Orders []*entities.OrderEntity

	if err := p.PgGorm.DB.Find(&Orders).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the Orders into the database."))
	}

	return Orders, nil
}

func (p PgOrderRepository) UpdateOrder(OrderEntity entities.OrderEntity) error {
	err := p.PgGorm.DB.Save(OrderEntity).Error
	if err != nil {
		return errors.Wrap(err, "error in the inserting Order into the database.")
	}

	return nil
}
