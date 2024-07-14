package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/sir-george2500/g-server/internal/auth"
	"github.com/sir-george2500/g-server/internal/database"
	"net/http"
	"time"
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

	responWithJson(w, 201, databaseUserToUser(user))
}

// handle respon with user
func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Couldn't create user post %v", err))
		return
	}
	responWithJson(w, 200, databasePostsToPosts(posts))
}
