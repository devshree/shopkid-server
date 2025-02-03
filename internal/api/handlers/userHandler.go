package handlers

import (
	"database/sql"
	"encoding/json"
	"kids-shop/internal/domain/models"
	"kids-shop/internal/repository/postgres"
	"net/http"
)

type UserHandler struct {
	db *sql.DB	
	userRepo *postgres.UserRepository
}

func NewUserHandler(db *sql.DB) *UserHandler {
		return &UserHandler{
		db: db,
		userRepo: postgres.NewUserRepository(db),
	}
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)

	var profile models.User
	err := h.db.QueryRow("SELECT id, email, name FROM users WHERE id = $1", userID).
		Scan(&profile.ID, &profile.Email, &profile.Name)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		http.Error(w, "Error encoding profile", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(models.UserIDKey).(int)

	var profile models.User
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if profile.ID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := h.userRepo.UpdateUser(&profile)
	if err != nil {
		http.Error(w, "Error updating profile", http.StatusInternalServerError)
		return
	}

	h.GetUserProfile(w, r)
} 