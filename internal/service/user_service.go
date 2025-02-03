package service

import (
	"database/sql"
	"kids-shop/internal/repository/postgres"
)

func GetUserRoleFromID(userID int, db *sql.DB) (string, error) {
	userRepo := postgres.NewUserRepository(db)
	return userRepo.GetUserRole(userID)
} 