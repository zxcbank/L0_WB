package repositories

import (
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go-template-microservice-v2/internal/data/contracts"
	"go-template-microservice-v2/internal/data/entities"
	gormpg "go-template-microservice-v2/pkg/gorm_pg"
)

type PgPaymentRepository struct {
	PgGorm *gormpg.PgGorm
}

func NewPgPaymentRepository(pgGorm *gormpg.PgGorm) contracts.IPaymentRepository {
	return &PgPaymentRepository{PgGorm: pgGorm}
}

func (p PgPaymentRepository) AddPayment(PaymentEntity entities.PaymentEntity) error {
	err := p.PgGorm.DB.Create(PaymentEntity).Error
	if err != nil {
		return errors.Wrap(err, "error in the inserting Payment into the database.")
	}

	return nil
}

func (p PgPaymentRepository) GetPayment(id uuid.UUID) (*entities.PaymentEntity, error) {
	var Payment entities.PaymentEntity

	if err := p.PgGorm.DB.First(&Payment, id).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the Payment with id %s into the database.", id))
	}

	return &Payment, nil
}

func (p PgPaymentRepository) GetAllPayment() ([]*entities.PaymentEntity, error) {
	var Payments []*entities.PaymentEntity

	if err := p.PgGorm.DB.Find(&Payments).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the Payments into the database."))
	}

	return Payments, nil
}

func (p PgPaymentRepository) UpdatePayment(PaymentEntity entities.PaymentEntity) error {
	err := p.PgGorm.DB.Save(PaymentEntity).Error
	if err != nil {
		return errors.Wrap(err, "error in the inserting Payment into the database.")
	}

	return nil
}
