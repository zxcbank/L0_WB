package entities

import (
	"github.com/google/uuid"
)

type PaymentEntity struct {
	Id            uuid.UUID `json:"id" form:"Id"`
	Order_id      uuid.UUID `json:"order_id" form:"Order_id"`
	Transaction   string    `json:"transaction" form:"Transaction"`
	Requested_id  uuid.UUID `json:"requested_id" form:"Requested_id"`
	Currency      string    `json:"currency" form:"Currency"`
	Provider      string    `json:"provider" form:"Provider"`
	Amount        int       `json:"amount" form:"Amount"`
	Payment_dt    int       `json:"payment_dt" form:"Payment_dt"`
	Bank          string    `json:"bank" form:"Bank"`
	Delivery_cost int       `json:"delivery_cost" form:"Delivery_cost"`
	Goods_total   int       `json:"goods_total" form:"Goods_total"`
	Custom_fee    int       `json:"custom_fee" form:"Custom_fee"`
}

func CreatePaymentEntity(order_id uuid.UUID, transaction string, requsted_id uuid.UUID,
	currency string, provider string, amount int,
	payment_dt int, bank string, delivery_cost int,
	goods_total int, custom_fee int) PaymentEntity {
	return PaymentEntity{
		Order_id:      order_id,
		Transaction:   transaction,
		Requested_id:  requsted_id,
		Currency:      currency,
		Provider:      provider,
		Amount:        amount,
		Payment_dt:    payment_dt,
		Bank:          bank,
		Delivery_cost: delivery_cost,
		Goods_total:   goods_total,
		Custom_fee:    custom_fee,
		Id:            uuid.New(),
	}
}
