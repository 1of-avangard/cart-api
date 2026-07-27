package service

import (
	"cart_api/internal/entity"
	"context"
)

func (s *Service) CreateCart(ctx context.Context) (entity.Cart, error) {
	id, err := s.repo.CreateCart(ctx)
	if err != nil {
		return entity.Cart{}, err
	}
	return entity.Cart{ID: id, Items: []entity.CartItem{}}, nil

}
