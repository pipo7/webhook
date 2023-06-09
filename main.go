package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type WebhookRequest struct {
	Event string `json:"event"`
	Data  struct {
		UserID int `json:"user_id"`
	} `json:"data"`
}

type WebhookResponse struct {
	Membership bool `json:"is_chirpy_red"`
}

func WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// first check whether it has authorization key
		fmt.Printf("http request: %v", r)
		authHeader := r.Header.Get("Authorization")
		fmt.Printf("authoHeader: %v", authHeader)
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		apiString := strings.TrimPrefix(authHeader, "ApiKey ")

		if apiString != apiCfg.APIKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		var req WebhookRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Event != "user.upgraded" {
			w.WriteHeader(http.StatusOK)
		}

		//get the users from db
		user, err := db.UpdateMembership(req.Data.UserID, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		res := WebhookResponse{
			Membership: true,
		}

		json.NewEncoder(w).Encode(res)
	}
}
func (db *DB) UpdateMembership(userID int, membership bool) (User, error) {
	// Load the current JSON data from the database file
	user, err := db.GetUser(userID)
	if err != nil {
		log.Error(err)
	}

	index := user.ID
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user.Membership = membership
	dbStructure.Users[index] = user

	err = db.writeDB(dbStructure)
	if err != nil {
		log.Error(err)
	}

	return user, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	apikey := os.Getenv("APIKey")

	apiCfg := &config.ApiConfig{
		FileserverHits: 0,
		JwtSecret:      jwtSecret,
		APIKey:         apikey,
	}

	WebhookHandler()
}
