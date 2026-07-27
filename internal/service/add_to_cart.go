package service

import (
	"cart_api/internal/entity"
	"cart_api/internal/errorsx"
	"context"
)

func (s *Service) AddCartItemToCart(ctx context.Context, cartID int, dto entity.AddCartItemDTO) (entity.CartItem, error) {
	if dto.Product == "" {
		return entity.CartItem{}, errorsx.ErrProductNotFound
	}

	if dto.Price <= 0 {
		return entity.CartItem{}, errorsx.ErrPriceNotFound
	}
	cart, err := s.repo.GetCart(ctx, cartID)
	if err != nil {
		return entity.CartItem{}, err
	}
	if len(cart.Items) >= 5 {
		return entity.CartItem{}, errorsx.ErrCartFull
	}

	return s.repo.AddCartItem(ctx, cartID, dto)

}
