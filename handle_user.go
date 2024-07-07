package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sir-george2500/g-server/internal/database"
)

// handle respon with user
func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramenters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	param := paramenters{}
	err := decoder.Decode(&param)

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Couldn't create user %v", err))
		return
	}

	responWithJson(w, 200, databaseUserToUser(user))
}
