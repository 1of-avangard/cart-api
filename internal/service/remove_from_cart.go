package service

import "context"

func (s *Service) RemoveCartItemFromCart(ctx context.Context, cartID int, itemID int) error {
	// 1. Проверяем, существует ли сама корзина
	_, err := s.repo.GetCart(ctx, cartID)
	if err != nil {
		return err
	}

	return s.repo.RemoveCartItem(ctx, cartID, itemID)
}
