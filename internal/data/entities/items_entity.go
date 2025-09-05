package entities

import (
	uuid "github.com/google/uuid"
)

type ItemsEntity struct {
	Id           uuid.UUID `json:"id" form:"id"`
	Order_id     uuid.UUID `json:"order_id" form:"order_id"`
	Chrt_id      uuid.UUID `json:"Chrt_id" form:"Chrt_id"`
	Track_number string    `json:"Track_number" form:"Track_number"`
	Price        int       `json:"Price" form:"Price"`
	Rid          uuid.UUID `json:"Rid" form:"Rid"`
	Name         string    `json:"Name" form:"Name"`
	Sale         int       `json:"Sale" form:"Sale"`
	Size         string    `json:"Size" form:"Size"`
	Total_price  int       `json:"Total_price" form:"Total_price"`
	Nm_id        uuid.UUID `json:"Nm_id" form:"Nm_id"`
	Brand        string    `json:"Brand" form:"Brand"`
	Status       int       `json:"Status" form:"Status"`
}

func CreateItemsEntity(order_id uuid.UUID, chrt_id uuid.UUID,
	track_number string, price int, rid uuid.UUID,
	name string, sale int, size string,
	total_price int, nm_id uuid.UUID, brand string,
	status int) ItemsEntity {
	return ItemsEntity{
		Order_id:     order_id,
		Chrt_id:      chrt_id,
		Track_number: track_number,
		Price:        price,
		Rid:          rid,
		Name:         name,
		Sale:         sale,
		Size:         size,
		Total_price:  total_price,
		Nm_id:        nm_id,
		Brand:        brand,
		Status:       status,
		Id:           uuid.New(),
	}
}
