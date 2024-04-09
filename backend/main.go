package main

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BaneleJerry/Web-Test/internal/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gopkg.in/boj/redistore.v1"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB    *database.Queries
	Hash  *Argon2idHash
	store *redistore.RediStore
}

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading ENV")
	}

	r := mux.NewRouter()

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	store, err := redistore.NewRediStore(10, "tcp", "6379", "", []byte(os.Getenv("SESSION_SECRET")))
	if port == "" || dbURL == "" || store == nil {
		panic("Please Check enviroment variables")
	}
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	defer db.Close()

	cfg := apiConfig{
		DB:    database.New(db),
		Hash:  NewArgon2idHash(1, 32, 64*1024, 32, 256),
		store: store,
	}
	r.HandleFunc("/healthcheck", Chain(cfg.HealthcheckHandler, Logging(), Method("GET")))
	r.HandleFunc("/", Chain(handleRoot, Logging(), Method("GET")))
	r.HandleFunc("/login", Chain(cfg.handleLogin, Logging(), Method("POST")))
	r.HandleFunc("/register", Chain(cfg.handleRegister, Logging(), Method("POST")))
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static"))))
	corsMux := middlewareCors(r)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      corsMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Println("Server Started listening on", port)
	log.Fatal(srv.ListenAndServe())
}

func handleRoot(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		ErrorResponse(w, errors.New("page Not found"), http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "../frontend/index.html")
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func randomSecret(length uint32) ([]byte, error) {

	secret := make([]byte, length)
	_, err := rand.Read(secret)

	if err != nil {

		return nil, err

	}

	return secret, nil

}
