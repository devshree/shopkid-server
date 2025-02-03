package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// GetUserProfile handles GET requests to fetch the user's profile
func (h *Handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// You'll need to implement the logic to:
	// 1. Get user ID from the session/token
	// 2. Fetch user profile from database
	// 3. Return the profile as JSON
}

// UpdateUserProfile handles PUT requests to update the user's profile
func (h *Handler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	// You'll need to implement the logic to:
	// 1. Get user ID from the session/token
	// 2. Decode the request body into a profile struct
	// 3. Validate the input
	// 4. Update the profile in database
	// 5. Return success response
}

// Auth handlers
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	err := h.db.QueryRow(
		"SELECT id, email, password, role FROM users WHERE email = $1",
		req.Email,
	).Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// TODO: Generate JWT token here
	// For now, just return the user without password
	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert user
	err = h.db.QueryRow(
		"INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at",
		user.Name, user.Email, string(hashedPassword), "buyer",
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = "" // Don't send password back
	json.NewEncoder(w).Encode(user)
}

// Cart handlers
func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	// TODO: Get user_id from JWT token
	userID := 1 // Temporary hardcoded value

	rows, err := h.db.Query(`
		SELECT ci.id, ci.user_id, ci.product_id, ci.quantity, ci.price, ci.created_at,
			   p.name, p.description, p.image, p.category, p.age_range, p.stock
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		WHERE ci.user_id = $1
	`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []CartItem
	for rows.Next() {
		var item CartItem
		item.Product = &Product{}
		err := rows.Scan(
			&item.ID, &item.UserID, &item.ProductID, &item.Quantity, &item.Price, &item.CreatedAt,
			&item.Product.Name, &item.Product.Description, &item.Product.Image,
			&item.Product.Category, &item.Product.Age_Range, &item.Product.Stock,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}

// Order handlers
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// TODO: Get user_id from JWT token
	userID := 1 // Temporary hardcoded value

	tx, err := h.db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Get cart items
	rows, err := tx.Query("SELECT product_id, quantity, price FROM cart_items WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var totalAmount float64
	var orderItems []OrderItem
	for rows.Next() {
		var item OrderItem
		err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		totalAmount += item.Price * float64(item.Quantity)
		orderItems = append(orderItems, item)
	}

	// Create order
	var order Order
	err = tx.QueryRow(
		"INSERT INTO orders (user_id, total_amount) VALUES ($1, $2) RETURNING id, created_at",
		userID, totalAmount,
	).Scan(&order.ID, &order.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create order items
	for _, item := range orderItems {
		_, err = tx.Exec(
			"INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
			order.ID, item.ProductID, item.Quantity, item.Price,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Clear cart
	_, err = tx.Exec("DELETE FROM cart_items WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order.Items = orderItems
	json.NewEncoder(w).Encode(order)
} 