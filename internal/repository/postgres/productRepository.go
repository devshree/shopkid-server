package postgres

import (
	"database/sql"
	"kids-shop/internal/domain/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	rows, err := r.db.Query(`
		SELECT id, name, description, price, category, 
			   age_range, stock, image, created_at 
		FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.Price,
			&p.Category, &p.Age_Range, &p.Stock,
			&p.Image, &p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	var p models.Product
	err := r.db.QueryRow(`
		SELECT id, name, description, price, category,
			   age_range, stock, image, created_at
		FROM products WHERE id = $1
	`, id).Scan(
		&p.ID, &p.Name, &p.Description, &p.Price,
		&p.Category, &p.Age_Range, &p.Stock,
		&p.Image, &p.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Create(p *models.Product) error {
	return r.db.QueryRow(`
		INSERT INTO products (name, description, price, category,
			age_range, stock, image)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`,
		p.Name, p.Description, p.Price, p.Category,
		p.Age_Range, p.Stock, p.Image,
	).Scan(&p.ID, &p.CreatedAt)
}

func (r *ProductRepository) Update(p *models.Product) error {
	result, err := r.db.Exec(`
		UPDATE products 
		SET name = $1, description = $2, price = $3,
			category = $4, age_range = $5, stock = $6,
			image = $7
		WHERE id = $8
	`,
		p.Name, p.Description, p.Price, p.Category,
		p.Age_Range, p.Stock, p.Image, p.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *ProductRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// ... implement other methods 