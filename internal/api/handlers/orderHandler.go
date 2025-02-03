package handlers

import (
	"database/sql"
	"encoding/json"
	"kids-shop/internal/domain/models"
	"kids-shop/internal/repository/postgres"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	db *sql.DB
	orderRepo *postgres.OrderRepository
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{
		db: db, 
		orderRepo: postgres.NewOrderRepository(db),
	}
}


func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)

	
	var orders []models.Order
	orders, err := h.orderRepo.GetOrdersByUserID(userID)
	if err != nil {
		http.Error(w, "Error fetching orders", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order.UserID = userID
	err := h.orderRepo.CreateOrder(&order)
	if err != nil {
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.orderRepo.GetOrderById(orderID, userID	)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	if order.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
} 

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	var order models.Order
	if err = json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order.ID = orderID
	if order.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	err = h.orderRepo.UpdateOrder(&order)
	if err != nil {
		http.Error(w, "Error updating order", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}		

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	vars := mux.Vars(r)
	orderID := vars["id"]

	
	_, err := h.db.Exec("DELETE FROM orders WHERE id = $1 AND user_id = $2", orderID, userID)
	
	if err != nil {
		http.Error(w, "Error deleting order", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["order_id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	_, err = h.orderRepo.GetOrderById(orderID, userID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	var orderItem models.OrderItem
	if err = json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	orderItem.OrderID = orderID
	err = h.orderRepo.CreateOrderItem(&orderItem)
	if err != nil {
		http.Error(w, "Error creating order item", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(orderItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}	

func (h *OrderHandler) GetOrderItems(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	_, err = h.orderRepo.GetOrderById(orderID, userID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	var orderItems []models.OrderItem
	orderItems, err = h.orderRepo.GetOrderItemsByOrderID(orderID)
	if err != nil {
		http.Error(w, "Error fetching order items", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(orderItems)
	if err != nil {
		http.Error(w, "Error encoding order items", http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)
	vars := mux.Vars(r)
	orderItemID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order item ID", http.StatusBadRequest)
		return
	}
	orderItem, err := h.orderRepo.GetOrderItemById(orderItemID)
	if err != nil {
		http.Error(w, "Order item not found", http.StatusNotFound)
		return
	}
	_, err = h.orderRepo.GetOrderById(orderItem.OrderID, userID)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = h.orderRepo.DeleteOrderItem(orderItemID)
	if err != nil {
		http.Error(w, "Error deleting order item", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"message": "Order item deleted successfully"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

