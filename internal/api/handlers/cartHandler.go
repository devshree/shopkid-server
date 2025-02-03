package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"kids-shop/internal/domain/models"
	"kids-shop/internal/repository/postgres"
)

type CartHandler struct {
	db *sql.DB
	cartRepo *postgres.CartRepository
}


func NewCartHandler(db *sql.DB) *CartHandler {
	cr := postgres.NewCartRepository(db)
	return &CartHandler{
		db: db,
		cartRepo: cr,
	}	
}	

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	
	cart, err := h.cartRepo.GetCart(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	var item models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.cartRepo.AddToCart(userID, item.ProductID, item.Quantity, item.Price); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	var item models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.cartRepo.RemoveFromCart(userID, item.ProductID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}		
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	if err := h.cartRepo.ClearCart(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CartHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	var item models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.cartRepo.UpdateCartItem(userID, item.ProductID, item.Quantity, item.Price); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
