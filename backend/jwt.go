package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/BaneleJerry/Web-Test/internal/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)



func createToken(user database.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user-id": user.ID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func (cfg *apiConfig) withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT auth middleware")

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			permissionDenied(w)
		}

		// Check if the header has the Bearer authentication scheme
		if !strings.HasPrefix(tokenString, "Bearer ") {
			permissionDenied(w)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}
		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userIDString := claims["user-id"].(string)

		userID, err := uuid.Parse(userIDString)
		if err != nil {
			ErrorResponse(w, errors.New("invalid token-79"), http.StatusInternalServerError)
			return
		}

		user, err := cfg.DB.GetUserByID(r.Context(), userID)
		if err != nil {
			ErrorResponse(w, errors.New("invalid token"), http.StatusForbidden)
			return
		}
		fmt.Println(user.Username)
		handlerFunc(w, r)
	}
}