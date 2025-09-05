package contracts

import (
	"go-template-microservice-v2/internal/data/entities"

	uuid "github.com/satori/go.uuid"
)

type IPaymentRepository interface {
	AddPayment(paymentEntity entities.PaymentEntity) error
	GetPayment(id uuid.UUID) (*entities.PaymentEntity, error)
	GetAllPayment() ([]*entities.PaymentEntity, error)
	UpdatePayment(paymmentEntity entities.PaymentEntity) error
}
