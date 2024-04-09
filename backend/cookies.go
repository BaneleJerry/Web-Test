package main

import (
	"net/http"
	"time"

	"github.com/BaneleJerry/Web-Test/internal/database"
)
func (cfg *apiConfig) setCookie(w http.ResponseWriter, r *http.Request, user database.User) error {

	apiKeyObj, err := cfg.getApiKey(*r, user)
	if err != nil {
		// If there's an error fetching the API key, try setting a new one
		err := cfg.setApiKey(r, user)
		if err != nil {
			return err
		}
		// Fetch the newly set API key
		apiKeyObj, err = cfg.getApiKey(*r, user)
		if err != nil {
			return err
		}
	}

	cookie := http.Cookie{
		Name:     "session_token" + user.ID.String(),
		Value:    string(apiKeyObj.ApiKeyValue),
		Path:     "/",
		MaxAge:   int(time.Until(time.Now().Add(7 * 24 * time.Hour)).Seconds()),
		HttpOnly: true, // Ensure cookie is only accessible via HTTP
		Secure:   true, // Enable secure cookie (HTTPS only)
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return nil
	
}
