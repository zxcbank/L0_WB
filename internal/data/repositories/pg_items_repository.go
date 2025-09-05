package repositories

import (
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go-template-microservice-v2/internal/data/contracts"
	"go-template-microservice-v2/internal/data/entities"
	gormpg "go-template-microservice-v2/pkg/gorm_pg"
)

type PgItemsRepository struct {
	PgGorm *gormpg.PgGorm
}

func NewPgItemsRepository(pgGorm *gormpg.PgGorm) contracts.IItemsRepository {
	return &PgItemsRepository{PgGorm: pgGorm}
}

func (p PgItemsRepository) AddItems(ItemsEntity entities.ItemsEntity) error {
	err := p.PgGorm.DB.Create(ItemsEntity).Error
	if err != nil {
		return errors.Wrap(err, "error in the inserting Items into the database.")
	}

	return nil
}

func (p PgItemsRepository) GetItems(id uuid.UUID) (*entities.ItemsEntity, error) {
	var Items entities.ItemsEntity

	if err := p.PgGorm.DB.First(&Items, id).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the Items with id %s into the database.", id))
	}

	return &Items, nil
}

func (p PgItemsRepository) GetAllItems() ([]*entities.ItemsEntity, error) {
	var Itemss []*entities.ItemsEntity

	if err := p.PgGorm.DB.Find(&Itemss).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the Itemss into the database."))
	}

	return Itemss, nil
}

func (p PgItemsRepository) UpdateItems(ItemsEntity entities.ItemsEntity) error {
	err := p.PgGorm.DB.Save(ItemsEntity).Error
	if err != nil {
		return errors.Wrap(err, "error in the inserting Items into the database.")
	}

	return nil
}
