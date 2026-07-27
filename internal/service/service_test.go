package service

import (
	"cart_api/internal/entity"
	"cart_api/internal/errorsx"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_AddCartItemToCart_CartFull(t *testing.T) {
	mockRepo := &mockCartRepository{
		getCartFunc: func(ctx context.Context, cartID int) (entity.Cart, error) {
			return entity.Cart{
				ID: cartID,
				Items: []entity.CartItem{
					{ID: 1, Product: "Item 1"},
					{ID: 2, Product: "Item 2"},
					{ID: 3, Product: "Item 3"},
					{ID: 4, Product: "Item 4"},
					{ID: 5, Product: "Item 5"},
				},
			}, nil
		},
	}
	svc := NewService(mockRepo)

	_, err := svc.AddCartItemToCart(context.Background(), 1, entity.AddCartItemDTO{
		Product: "Item 6",
		Price:   100.0,
	})

	assert.Error(t, err)

	assert.ErrorIs(t, err, errorsx.ErrCartFull)
}

// Тест 2: Успешное добавление (В корзине есть место)
func TestService_AddCartItemToCart_Success(t *testing.T) {
	mockRepo := &mockCartRepository{
		getCartFunc: func(ctx context.Context, cartID int) (entity.Cart, error) {
			return entity.Cart{ID: cartID, Items: []entity.CartItem{}}, nil
		},
		addCartItemFunc: func(ctx context.Context, cartID int, dto entity.AddCartItemDTO) (entity.CartItem, error) {
			return entity.CartItem{
				ID:      42,
				CartId:  cartID,
				Product: dto.Product,
				Price:   dto.Price,
			}, nil
		},
	}
	svc := NewService(mockRepo)

	expectedDTO := entity.AddCartItemDTO{Product: "Shoes", Price: 2500.50}

	result, err := svc.AddCartItemToCart(context.Background(), 1, expectedDTO)

	assert.NoError(t, err)
	assert.Equal(t, 42, result.ID)
	assert.Equal(t, expectedDTO.Product, result.Product)
	assert.Equal(t, expectedDTO.Price, result.Price)
}
