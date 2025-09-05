package entities

import (
	"time"

	"github.com/google/uuid"
)

type OrderEntity struct {
	Id                 uuid.UUID `json:"id" form:"id"`
	Track_number       string    `json:"track_number" form:"track_number"`
	Entry              string    `json:"entry" form:"entry"`
	Locale             string    `json:"locale" form:"locale"`
	Internal_signature string    `json:"internal_signature" form:"internal_signature"`
	Custromer_id       uuid.UUID `json:"customer_id" form:"customer_id"`
	Delivery_service   string    `json:"delivery_service" form:"delivery_service"`
	Shardkey           string    `json:"shardkey" form:"shardkey"`
	Sm_id              uuid.UUID `json:"sm_id" form:"sm_id"`
	Date_created       time.Time `json:"date_created" form:"date_created"`
	Oof_shard          string    `json:"oof_shard" form:"oof_shard"`
}

func CreateOrderEntity(track_number string, entry string,
	locale string, internal_signature string, customer_id uuid.UUID,
	delivery_service string, shardkey string, sm_id uuid.UUID,
	date_created time.Time, oof_shard string) OrderEntity {

	return OrderEntity{
		Track_number:       track_number,
		Entry:              entry,
		Locale:             locale,
		Internal_signature: internal_signature,
		Custromer_id:       customer_id,
		Delivery_service:   delivery_service,
		Shardkey:           shardkey,
		Sm_id:              sm_id,
		Date_created:       date_created,
		Oof_shard:          oof_shard,
		Id:                 uuid.New(),
	}
}
