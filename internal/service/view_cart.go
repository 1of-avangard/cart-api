package service

import (
	"cart_api/internal/entity"
	"context"
)

func (s *Service) ViewCart(ctx context.Context, cartID int) (entity.Cart, error) {
	return s.repo.GetCart(ctx, cartID)
}

func (s *Service) CalculateCartPrice(ctx context.Context, cartID int) (entity.CartPriceResponse, error) {
	cart, err := s.repo.GetCart(ctx, cartID)
	if err != nil {
		return entity.CartPriceResponse{}, err
	}
	totalPrice := 0.0

	for _, item := range cart.Items {
		totalPrice += item.Price
	}

	if totalPrice > 5000.0 {
		return entity.CartPriceResponse{
			CartId:          cartID,
			TotalPrice:      totalPrice,
			DiscountPercent: 10,
			FinalPrice:      totalPrice * 0.9,
		}, nil
	} else if len(cart.Items) > 3 {
		return entity.CartPriceResponse{
			CartId:          cartID,
			TotalPrice:      totalPrice,
			DiscountPercent: 5,
			FinalPrice:      totalPrice * 0.95,
		}, nil
	}

	return entity.CartPriceResponse{
		CartId:          cartID,
		TotalPrice:      totalPrice,
		DiscountPercent: 0,
		FinalPrice:      totalPrice,
	}, nil
}
