package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/BaneleJerry/Web-Test/internal/database"
	"github.com/google/uuid"
)

func generateAPIKey() (string, error) {
	// Generate a random byte slice to be used as the API key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	// Encode the byte slice using base64 encoding
	apiKey := base64.StdEncoding.EncodeToString(key)

	return apiKey, nil
}

func GenerateAPIKeyHash(apiKey string) (string, []byte, error) {
	// Generate a random salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", nil, err
	}

	// Append salt to the API key
	apiKeyWithSalt := []byte(apiKey)
	apiKeyWithSalt = append(apiKeyWithSalt, salt...)

	// Hash the API key with salt using SHA-256
	hash := sha256.Sum256(apiKeyWithSalt)

	// Encode the hash to base64
	encodedHash := base64.URLEncoding.EncodeToString(hash[:])

	return encodedHash, salt, nil
}

func (cfg *apiConfig) setApiKey(r *http.Request, user database.User) error {

	apiKey, err := generateAPIKey()
	if err != nil {
		return err
	}

	// Hash the API key
	hashedApiKey, salt, err := GenerateAPIKeyHash(apiKey)
	if err != nil {
		return err
	}

	// Calculate expiration time (90 days from now)
	expirationTime := time.Now().Add(90 * 24 * time.Hour)

	// Insert the API key into the database
	_, err = cfg.DB.InsertAPIKey(r.Context(), database.InsertAPIKeyParams{
		ID:             uuid.New(),
		UserID:         user.ID,
		ApiKeyValue:    []byte(hashedApiKey),
		Salt:           salt,
		ExpirationTime: expirationTime,
		CreatedAt:      time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (cfg *apiConfig) getApiKey(r http.Request, user database.User) (database.ApiKey, error) {

	apikey, err := cfg.DB.SelectAPIKeysByUserID(r.Context(), user.ID)
	if err != nil {
		return database.ApiKey{}, err
	}

	return apikey, nil
}
