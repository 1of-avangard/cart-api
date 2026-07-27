package service

import (
	"cart_api/internal/entity"
	"context"
)

type CartService interface {
	CreateCart(ctx context.Context) (entity.Cart, error)
	ViewCart(ctx context.Context, cartID int) (entity.Cart, error)
	CalculateCartPrice(ctx context.Context, cartID int) (entity.CartPriceResponse, error)
	AddCartItemToCart(ctx context.Context, cartID int, dto entity.AddCartItemDTO) (entity.CartItem, error)
	RemoveCartItemFromCart(ctx context.Context, cartID int, itemID int) error
}
