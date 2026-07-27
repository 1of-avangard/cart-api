package entity

type CartItem struct {
	ID      int     `json:"id" db:"id"`
	CartId  int     `json:"cart_id" db:"cart_id"`
	Product string  `json:"product" db:"product"`
	Price   float64 `json:"price" db:"price"`
}
