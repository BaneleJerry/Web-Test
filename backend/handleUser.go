package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	passwordHash, err := Hash.GenerateHash([]byte(params.Password), nil)
	if err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// Create new dbUser in the database
	dbUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
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

	// Encode dbUser information as JSON and send response

	respondWithJSON(w, http.StatusOK, databaseUserToUser(dbUser,""))
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decode JSON request body into params struct
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	// Retrieve dbUser from the database
	dbUser, err := cfg.DB.GetUserByUsername(r.Context(), params.Username)
	if err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// If dbUser doesn't exist, return a 400 Bad Request
	if dbUser.ID == uuid.Nil { // Assuming ID is the primary key and zero means not found
		ErrorResponse(w, errors.New("dbUser not found"), http.StatusBadRequest)
		return
	}

	// Compare passwords
	err = Hash.Compare(dbUser.PasswordHash, dbUser.Salt, []byte(params.Password))
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	token, err := createToken(dbUser)
	if err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	User := databaseUserToUser(dbUser, token)
	respondWithJSON(w, http.StatusOK, User)

}

// func (cfg *apiConfig) LogoutHandler(w http.ResponseWriter, r *http.Request) {
// 	session, err := cfg.store.Get(r, "session-id")
// 	if err != nil {
// 		// Handle error retrieving session
// 		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
// 		return
// 	}

// 	session.Options.MaxAge = -1
// 	err = session.Save(r, w)
// 	if err != nil {
// 		// Handle error saving session
// 		http.Error(w, "Error saving session", http.StatusInternalServerError)
// 		return
// 	}

// 	// Write a confirmation message to the response writer
// 	w.Write([]byte("Logout successful"))
// }
