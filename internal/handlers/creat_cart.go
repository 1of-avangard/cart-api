package handlers

import (
	"encoding/json"
	"net/http"
)

func (s *Server) CreateCart(w http.ResponseWriter, r *http.Request) {
	cart, err := s.service.CreateCart(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cart)

}
