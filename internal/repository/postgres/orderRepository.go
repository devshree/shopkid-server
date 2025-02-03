package postgres

import (
	"database/sql"

	"kids-shop/internal/domain/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (user_id, total_amount, status) VALUES ($1, $2, $3)", order.UserID, order.TotalAmount, order.Status)
	return err
}

func (r *OrderRepository) GetOrdersByUserID(userID int) ([]models.Order, error) {
	rows, err := r.db.Query("SELECT id, user_id, total_amount, status, created_at, updated_at FROM orders WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderById(id int) (*models.Order, error) {
	row := r.db.QueryRow("SELECT id, user_id, total_amount, status, created_at, updated_at FROM orders WHERE id = $1", id)

	var order models.Order
	err := row.Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	_, err := r.db.Exec("UPDATE orders SET user_id = $1, total_amount = $2, status = $3, created_at = $4, updated_at = $5 WHERE id = $6", order.UserID, order.TotalAmount, order.Status, order.CreatedAt, order.UpdatedAt, order.ID)
	return err
}	

func (r *OrderRepository) DeleteOrder(id int) error {
	_, err := r.db.Exec("DELETE FROM orders WHERE id = $1", id)
	return err
}

func (r *OrderRepository) CreateOrderItem(orderItem *models.OrderItem) error {
	_, err := r.db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}

func (r *OrderRepository) GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error) {
	rows, err := r.db.Query("SELECT id, order_id, product_id, quantity, price, created_at, updated_at FROM order_items WHERE order_id = $1", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderItems []models.OrderItem
	for rows.Next() {
		var orderItem models.OrderItem
		err := rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Price, &orderItem.CreatedAt, &orderItem.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}


func (r *OrderRepository) DeleteOrderItem(id int) error {
	_, err := r.db.Exec("DELETE FROM order_items WHERE id = $1", id)
	return err
}


