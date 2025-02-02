package main

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"` // "clothes" or "toys"
	Age_Range   string  `json:"age_range"`
	Stock       int     `json:"stock"`
}

type CartItem struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
} 	
