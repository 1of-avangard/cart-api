package handlers

import (
	"cart_api/internal/entity"
	"cart_api/internal/errorsx"
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) AddCartItemToCart(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("cartID")
	cartID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid cart id", http.StatusBadRequest)
		return
	}
	var dto entity.AddCartItemDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	item, err := s.service.AddCartItemToCart(r.Context(), cartID, dto)
	if err != nil {
		if err == errorsx.ErrProductNotFound || err == errorsx.ErrPriceNotFound || err == errorsx.ErrCartFull {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)

}
