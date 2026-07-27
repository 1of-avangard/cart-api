package handlers

import (
	"cart_api/internal/errorsx"
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) ViewCart(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("cartID")
	cartID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid cart id", http.StatusBadRequest)
		return
	}
	cart, err := s.service.ViewCart(r.Context(), cartID)
	if err != nil {
		if err == errorsx.ErrCartNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (s *Server) CalculateCartPrice(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("cartID")
	cartID, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid cart id", http.StatusBadRequest)
		return
	}
	res, err := s.service.CalculateCartPrice(r.Context(), cartID)

	if err != nil {
		if err == errorsx.ErrCartNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
