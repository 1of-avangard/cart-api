package repository

import (
	"cart_api/internal/entity"
	"context"
)

func (r *Repository) AddCartItem(ctx context.Context, cartID int, dto entity.AddCartItemDTO) (entity.CartItem, error) {

	query := "INSERT INTO cart_items (cart_id, product, price) VALUES ($1, $2, $3) RETURNING id, cart_id, product, price"

	var item entity.CartItem

	err := r.db.GetContext(ctx, &item, query, cartID, dto.Product, dto.Price)
	if err != nil {
		return entity.CartItem{}, err
	}

	return item, nil
}
