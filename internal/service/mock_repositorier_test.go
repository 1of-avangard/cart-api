package service

import (
	"cart_api/internal/entity"
	"context"
)

type mockCartRepository struct {
	getCartFunc     func(ctx context.Context, cartID int) (entity.Cart, error)
	addCartItemFunc func(ctx context.Context, cartID int, dto entity.AddCartItemDTO) (entity.CartItem, error)
}

func (m *mockCartRepository) GetCart(ctx context.Context, cartID int) (entity.Cart, error) {
	return m.getCartFunc(ctx, cartID)
}

func (m *mockCartRepository) AddCartItem(ctx context.Context, cartID int, dto entity.AddCartItemDTO) (entity.CartItem, error) {
	return m.addCartItemFunc(ctx, cartID, dto)
}

func (m *mockCartRepository) CreateCart(ctx context.Context) (int, error) { return 0, nil }
func (m *mockCartRepository) RemoveCartItem(ctx context.Context, cartID int, itemID int) error {
	return nil
}
