package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BaneleJerry/Web-Test/internal/database"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type apiConfig struct {
	port string
	DB   *database.Queries
}

func newApiServer(port string, db *sql.DB) *apiConfig {
	return &apiConfig{
		port: port,
		DB:   database.New(db),
	}
}

func (cfg *apiConfig) apiRun() {
	router := mux.NewRouter()

	// Your routing logic goes here-

	// router.HandleFunc("/logout", Chain(cfg.LogoutHandler, Logging(), Method("GET")))

	router.HandleFunc("/login", Chain(cfg.handleLogin, Logging(), Method("POST")))
	router.HandleFunc("/signup", Chain(cfg.handleRegister, Logging(), Method("POST")))
	// r.HandleFunc("/auth/authStatus", Chain(cfg.handleAuthStatus, Logging(), Method("Get")))
	router.HandleFunc("/rr",cfg.withJWTAuth(func(w http.ResponseWriter, r *http.Request) {}))
	//routing End/Up
	corsMux := middlewareCors(router)
	srv := &http.Server{
		Addr:         ":" + cfg.port,
		Handler:      corsMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	fmt.Println("Server Started listening on", port)
	log.Fatal(srv.ListenAndServe())
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
