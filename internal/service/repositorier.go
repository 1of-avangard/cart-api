package service

import (
	"cart_api/internal/entity"
	"context"
)

type CartRepository interface {
	CreateCart(ctx context.Context) (int, error)
	GetCart(ctx context.Context, cartID int) (entity.Cart, error)
	AddCartItem(ctx context.Context, cartID int, dto entity.AddCartItemDTO) (entity.CartItem, error)
	RemoveCartItem(ctx context.Context, cartID int, itemID int) error
}
