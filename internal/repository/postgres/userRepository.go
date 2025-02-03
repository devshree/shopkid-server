package postgres

import (
	"database/sql"

	"kids-shop/internal/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password, role FROM users WHERE email = $1", email)
	
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Password, user.Role)
	return err
}


func (r *UserRepository) UpdateUser(user *models.User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2, password = $3, role = $4 WHERE id = $5", user.Name, user.Email, user.Password, user.Role, user.ID)
	return err
}

func (r *UserRepository) GetUserRole(userID int) (string, error) {
	var role string
	err := r.db.QueryRow("SELECT role FROM users WHERE id = $1", userID).Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}


