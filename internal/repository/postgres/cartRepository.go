package postgres

import (
	"database/sql"

	"kids-shop/internal/domain/models"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) GetCart(userID int) ([]models.CartItem, error) {
	rows, err := r.db.Query(`
		SELECT ci.id, ci.product_id, ci.quantity, ci.price, p.name, p.image
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		WHERE ci.user_id = $1
	`, userID)
		if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []models.CartItem
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.Price, &item.Product.Name, &item.Product.Image); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, item)
	}
	return cartItems, nil
}

func (r *CartRepository) AddToCart(userID int, productID int, quantity int, price float64) error {
	_, err := r.db.Exec(`
		INSERT INTO cart_items (user_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
	`, userID, productID, quantity, price)
	if err != nil {
		return err
	}
	return nil
}

func (r *CartRepository) RemoveFromCart(userID int, cartItemID int) error {
	_, err := r.db.Exec(`
		DELETE FROM cart_items WHERE user_id = $1 AND id = $2
	`, userID, cartItemID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CartRepository) ClearCart(userID int) error {
	_, err := r.db.Exec(`
		DELETE FROM cart_items WHERE user_id = $1
	`, userID)
	return err
}

func (r *CartRepository) UpdateCartItem(userID int, cartItemID int, quantity int) error {
	_, err := r.db.Exec(`
		UPDATE cart_items SET quantity = $1 WHERE user_id = $2 AND id = $3
	`, quantity, userID, cartItemID)
	return err
}

func (r *CartRepository) GetCartTotal(userID int) (float64, error) {
	rows, err := r.db.Query(`
		SELECT SUM(ci.quantity * ci.price) FROM cart_items ci WHERE ci.user_id = $1
	`, userID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var total float64
	if rows.Next() {
		if err := rows.Scan(&total); err != nil {
			return 0, err
		}
	}
	return total, nil
}		

func (r *CartRepository) GetCartCount(userID int) (int, error) {
	rows, err := r.db.Query(`
		SELECT COUNT(*) FROM cart_items WHERE user_id = $1
	`, userID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (r *CartRepository) GetCartItems(userID int) ([]models.CartItem, error) {
	rows, err := r.db.Query(`
		SELECT ci.id, ci.product_id, ci.quantity, ci.price, p.name, p.image
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		WHERE ci.user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []models.CartItem
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.Price, &item.Product.Name, &item.Product.Image); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, item)
	}
	return cartItems, nil
}
