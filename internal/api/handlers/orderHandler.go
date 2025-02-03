package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	db *sql.DB
}

type Order struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Total  float64 `json:"total"`
	Status string  `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Items []OrderItem `json:"items"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{db: db}
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	rows, err := h.db.Query("SELECT id, user_id, total, status FROM orders WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, "Error fetching orders", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.Total, &order.Status); err != nil {
			http.Error(w, "Error scanning orders", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order.UserID = userID
	err := h.db.QueryRow(
		"INSERT INTO orders (user_id, total, status) VALUES ($1, $2, $3) RETURNING id",
		order.UserID, order.Total, "pending",
	).Scan(&order.ID)

	if err != nil {
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	vars := mux.Vars(r)
	orderID := vars["id"]

	var order Order
	err := h.db.QueryRow(
		"SELECT id, user_id, total, status FROM orders WHERE id = $1 AND user_id = $2",
		orderID, userID,
	).Scan(&order.ID, &order.UserID, &order.Total, &order.Status)

	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
} 

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	vars := mux.Vars(r)
	orderID := vars["id"]

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	order.UserID = userID
	err := h.db.QueryRow(
		"UPDATE orders SET total = $1, status = $2 WHERE id = $3 AND user_id = $4",
		order.Total, order.Status, orderID, userID,
	).Scan()
	
	if err != nil {
		http.Error(w, "Error updating order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}		

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	vars := mux.Vars(r)
	orderID := vars["id"]

	
	_, err := h.db.Exec("DELETE FROM orders WHERE id = $1 AND user_id = $2", orderID, userID)
	
	if err != nil {
		http.Error(w, "Error deleting order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
}

func (h *OrderHandler) CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	// userID := r.Context().Value("userID").(int)
	vars := mux.Vars(r)
	orderID := vars["order_id"]

	var orderItem OrderItem
	if err := json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.db.QueryRow(
		"INSERT INTO order_items (order_id, product_id, quantity, price, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		orderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price, 
	).Scan(&orderItem.ID)
	
	if err != nil {
		http.Error(w, "Error creating order item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orderItem)
}	

func (h *OrderHandler) GetOrderItems(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	vars := mux.Vars(r)
	orderID := vars["id"]

	rows, err := h.db.Query("SELECT id, order_id, product_id, quantity, price FROM order_items WHERE order_id = $1 AND user_id = $2", orderID, userID)
	if err != nil {
		http.Error(w, "Error fetching order items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orderItems []OrderItem
	for rows.Next() {
		var orderItem OrderItem
		if err := rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Price); err != nil {
			http.Error(w, "Error scanning order items", http.StatusInternalServerError)
			return
		}
		orderItems = append(orderItems, orderItem)
	}

	json.NewEncoder(w).Encode(orderItems)
}

func (h *OrderHandler) DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	vars := mux.Vars(r)
	orderItemID := vars["id"]

	_, err := h.db.Exec("DELETE FROM order_items WHERE id = $1 AND user_id = $2", orderItemID, userID)
	if err != nil {
		http.Error(w, "Error deleting order item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Order item deleted successfully"})
}

