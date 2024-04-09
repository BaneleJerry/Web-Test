package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/BaneleJerry/Web-Test/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleRegister(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	// Decode JSON request body into params struct
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// Generate password hash
	passwordHash, err := cfg.Hash.GenerateHash([]byte(params.Password), nil)
	if err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// Create new user in the database
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:           uuid.New(),
		Username:     params.Username,
		PasswordHash: passwordHash.Hash,
		Salt:         passwordHash.Salt,
		Email:        params.Email,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		ErrorResponse(w, err, http.StatusConflict)
		return
	}

	// Encode user information as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	session, err := cfg.store.Get(r, "session-id")
	if err != nil {
		ErrorResponse(w, err, http.StatusForbidden)
	}

	var params struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decode JSON request body into params struct
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// Retrieve user from the database
	user, err := cfg.DB.GetUserByUsername(r.Context(), params.Username)
	if err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// If user doesn't exist, return a 400 Bad Request
	if user.ID == uuid.Nil { // Assuming ID is the primary key and zero means not found
		ErrorResponse(w, errors.New("user not found"), http.StatusBadRequest)
		return
	}

	// Compare passwords
	err = cfg.Hash.Compare(user.PasswordHash, user.Salt, []byte(params.Password))
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// Set HTTP status code and redirect to homepage
	session.Values["authenticated"] = true
	session.Save(r, w)

	// cfg.setCookie(w, r, user)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
