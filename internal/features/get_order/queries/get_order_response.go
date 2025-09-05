package queries

import (
	"time"

	uuid "github.com/google/uuid"
)

type GetOrderResponse struct {
	Order GetOrderResponseItem `json:"items,omitempty"`
}

type GetOrderResponseItem struct {
	Id                 uuid.UUID `json:"Id"`
	Track_number       string    `json:"track_number"`
	Entry              string    `json:"entry"`
	Locale             string    `json:"locale"`
	Internal_signature string    `json:"internal_signature"`
	Custromer_id       uuid.UUID `json:"custromer_id"`
	Delivery_service   string    `json:"delivery_service"`
	Shardkey           string    `json:"shardkey"`
	Sm_id              uuid.UUID `json:"sm_id"`
	Date_created       time.Time `json:"date_created"`
	Oof_shard          string    `json:"oof_shard"`
}
