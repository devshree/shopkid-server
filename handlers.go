package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	rows, err := h.db.Query("SELECT id, name, description, price, category, age_range, stock FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Age_Range, &p.Stock)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var p Product
	err = h.db.QueryRow("SELECT id, name, description, price, category, age_range, stock FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Age_Range, &p.Stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.db.QueryRow(
		"INSERT INTO products (name, description, price, category, age_range, stock) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		p.Name, p.Description, p.Price, p.Category, p.Age_Range, p.Stock).Scan(&p.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	var items []CartItem
	rows, err := h.db.Query("SELECT id, product_id, quantity, price FROM cart_items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item CartItem
		err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var item CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.db.QueryRow(
		"INSERT INTO cart_items (product_id, quantity, price) VALUES ($1, $2, $3) RETURNING id",
		item.ProductID, item.Quantity, item.Price).Scan(&item.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (h *Handler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid cart item ID", http.StatusBadRequest)
		return
	}

	_, err = h.db.Exec("DELETE FROM cart_items WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
} 