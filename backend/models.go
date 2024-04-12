package main

import (
	"time"

	"github.com/BaneleJerry/Web-Test/internal/database"
)

type User struct {
	Token     string    `json:"token"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserToUser(user database.User, token string) User {
	return User{
		Token:     token,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
