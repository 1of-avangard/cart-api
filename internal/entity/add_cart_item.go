package entity

type AddCartItemDTO struct {
	Product string  `json:"product"`
	Price   float64 `json:"price"`
}
