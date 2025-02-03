package postgres

import (
	"database/sql"

	"kids-shop/internal/domain/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateLogin(login *models.LoginHistory) error {
	_, err := r.db.Exec("INSERT INTO logins (user_id, status) VALUES ($1, $2)", login.UserId, login.Status)
	return err
}

func (r *AuthRepository) GetLoginsByUserID(userID int) ([]models.LoginHistory, error) {
	rows, err := r.db.Query("SELECT id, user_id, status, created_at FROM logins WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var logins []models.LoginHistory
	for rows.Next() {
		var login models.LoginHistory
		err := rows.Scan(&login.ID, &login.UserId, &login.Status, &login.CreatedAt)
		if err != nil {
			return nil, err
		}
		logins = append(logins, login)
	}
	return logins, nil
}


