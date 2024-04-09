package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (cfg *apiConfig) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {

	session, err := cfg.store.Get(r, "session-id")
	if err != nil {
		ErrorResponse(w, err, http.StatusForbidden)
	}

	if session.Values["authenticated"] != nil &&
		session.Values["authenticated"] != false {
		w.WriteHeader(http.StatusOK)
		message := fmt.Sprintf("Server is alive and ready to accept requests - server time: %s", time.Now().Format(time.RFC3339))
		w.Write([]byte(message))
		return
	}
	fmt.Println(session.ID)
	ErrorResponse(w, errors.New("Forbidden"), http.StatusForbidden)
}
