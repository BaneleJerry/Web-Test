package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/gob"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	port, dbURL string
	db          *sql.DB
	Hash        *Argon2idHash
)

func init() {
	Hash = NewArgon2idHash(1, 32, 64*1024, 32, 256)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading ENV:", err)
	}

	port = os.Getenv("PORT")
	dbURL = os.Getenv("DB_URL")
	if port == "" || dbURL == "" {
		log.Fatal("Please Check environment variables")
	}

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	gob.Register(User{})
}

func main() {
	server := newApiServer(port, db)
	server.apiRun()
}

func randomSecret(length uint32) ([]byte, error) {

	secret := make([]byte, length)
	_, err := rand.Read(secret)

	if err != nil {

		return nil, err

	}

	return secret, nil

}
