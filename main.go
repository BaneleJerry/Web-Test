package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cookie := http.Cookie{
			Name:    "Banele",
			Value:   "BaneletHA",
			Expires: time.Now().Add(time.Minute),
		}
		http.SetCookie(w, &cookie)

		w.Write([]byte("Hello World!"))
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "login.html")
		} else if r.Method == http.MethodPost {
			fmt.Println(r.FormValue("username"))
			fmt.Println(r.FormValue("password"))
			fmt.Println(r.FormValue("remember"))

			cookie := http.Cookie{
				Name:    "Banele",
				Value:   "BaneletHA",
				Expires: time.Now().Add(time.Minute),
			}
			http.SetCookie(w, &cookie)

			w.Write([]byte("Hello World!"))
		}
	})

	corsMux := middlewareCors(mux)

	srv := &http.Server{Addr: port, Handler: corsMux}

	fmt.Println("Server Started listening on", port)
	log.Fatal(srv.ListenAndServe())
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
