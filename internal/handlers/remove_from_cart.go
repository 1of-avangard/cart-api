package handlers

import (
	"cart_api/internal/errorsx"
	"net/http"
	"strconv"
)

func (s *Server) RemoveCartItemFromCart(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("cartID")
	itStr := r.PathValue("itemID")
	cartID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid cart id", http.StatusBadRequest)
		return
	}
	itemID, err := strconv.Atoi(itStr)
	if err != nil {
		http.Error(w, "invalid item id", http.StatusBadRequest)
		return
	}
	err = s.service.RemoveCartItemFromCart(r.Context(), cartID, itemID)
	if err != nil {
		if errorsx.ErrItemNotFound == err {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{}"))
}
