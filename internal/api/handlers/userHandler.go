package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	db *sql.DB
}

type UserProfile struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	var profile UserProfile
	err := h.db.QueryRow("SELECT id, email, name FROM users WHERE id = $1", userID).
		Scan(&profile.ID, &profile.Email, &profile.Name)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(profile)
}

func (h *UserHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	var profile UserProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := h.db.Exec("UPDATE users SET name = $1 WHERE id = $2",
		profile.Name, userID)
	if err != nil {
		http.Error(w, "Error updating profile", http.StatusInternalServerError)
		return
	}

	h.GetUserProfile(w, r)
} 