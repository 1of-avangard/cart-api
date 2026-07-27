package handlers

import (
	"cart_api/internal/entity"
	"context"
	"net/http"
)

type CartService interface {
	CreateCart(ctx context.Context) (entity.Cart, error)
	ViewCart(ctx context.Context, cartID int) (entity.Cart, error)
	CalculateCartPrice(ctx context.Context, cartID int) (entity.CartPriceResponse, error)
	AddCartItemToCart(ctx context.Context, cartID int, dto entity.AddCartItemDTO) (entity.CartItem, error)
	RemoveCartItemFromCart(ctx context.Context, cartID int, itemID int) error
}

type Server struct {
	service CartService
}

func NewServer(service CartService) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /carts", s.CreateCart)
	mux.HandleFunc("GET /carts/{cartID}", s.ViewCart)
	mux.HandleFunc("GET /carts/{cartID}/price", s.CalculateCartPrice)
	mux.HandleFunc("POST /carts/{cartID}/items", s.AddCartItemToCart)
	mux.HandleFunc("DELETE /carts/{cartID}/items/{itemID}", s.RemoveCartItemFromCart)
	return mux
}
