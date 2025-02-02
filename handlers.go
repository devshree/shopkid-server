package main

import (
	"database/sql"
	"encoding/json"
	"log"
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
	rows, err := h.db.Query("SELECT id, name, description, price, category, age_range, stock, image, created_at FROM products")
	if err != nil {
		log.Println("Error querying products:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Age_Range, &p.Stock, &p.Image, &p.CreatedAt)
		if err != nil {
			log.Println("Error scanning product:", err)
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
		log.Println("Error converting product ID:", err)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var p Product
	err = h.db.QueryRow("SELECT id, name, description, price, category, age_range, stock, image, created_at FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Age_Range, &p.Stock, &p.Image, &p.CreatedAt)
	if err != nil {
		log.Println("Error querying product:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Println("Error decoding product:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.db.QueryRow(
		"INSERT INTO products (name, description, price, category, age_range, stock, image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		p.Name, p.Description, p.Price, p.Category, p.Age_Range, p.Stock, p.Image).Scan(&p.ID)
	if err != nil {
		log.Println("Error creating product:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	var items []CartItem
	rows, err := h.db.Query("SELECT id, product_id, quantity, price FROM cart_items")
	if err != nil {
		log.Println("Error querying cart items:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item CartItem
		err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			log.Println("Error scanning cart item:", err)
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
		log.Println("Error decoding cart item:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.db.QueryRow(
		"INSERT INTO cart_items (product_id, quantity, price) VALUES ($1, $2, $3) RETURNING id",
		item.ProductID, item.Quantity, item.Price).Scan(&item.ID)
	if err != nil {
		log.Println("Error adding to cart:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (h *Handler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Error converting cart item ID:", err)
		http.Error(w, "Invalid cart item ID", http.StatusBadRequest)
		return
	}

	_, err = h.db.Exec("DELETE FROM cart_items WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting cart item:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.db.Exec(
		"UPDATE products SET name=$1, description=$2, price=$3, category=$4, age_range=$5, stock=$6, image=$7, updated_at=CURRENT_TIMESTAMP WHERE id=$8",
		p.Name, p.Description, p.Price, p.Category, p.Age_Range, p.Stock, p.Image, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	result, err := h.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
} 