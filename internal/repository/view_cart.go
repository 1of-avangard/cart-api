package repository

import (
	"cart_api/internal/entity"
	"cart_api/internal/errorsx"
	"context"
	"database/sql"
)

func (r *Repository) GetCart(ctx context.Context, cartID int) (entity.Cart, error) {

	var cart entity.Cart
	err := r.db.GetContext(ctx, &cart, "SELECT id FROM carts WHERE id = $1", cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Cart{}, errorsx.ErrCartNotFound
		}
		return entity.Cart{}, err
	}

	var items []entity.CartItem
	err = r.db.SelectContext(ctx, &items, "SELECT id, cart_id, product, price FROM cart_items WHERE cart_id = $1", cartID)
	if err != nil {
		return entity.Cart{}, err
	}
	cart.Items = items
	return cart, nil

}
