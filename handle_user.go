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

type CreateUser struct {
	Name string `json:"name"`
}

// handleCreateUser creates a new user
// @Summary Create User
// @Description Create a new user with the given parameters
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUser true "User Parameters"
// @Success 201 {object} CreateUser
// @Failure 400 {string} string "Error message"
// @Router /users [post]
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

// @Summary Get User
// @Description Retrieve details of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Failure 400 {string} string "Error message"
// @Router /users [get]
// @Security ApiKeyAuth
func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responWithJson(w, 200, databaseUserToUser(user))
}

// @Summary Get Posts for User
// @Description Retrieve posts associated with the authenticated user
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {array} Post
// @Failure 400 {string} string "Error message"
// @Router /posts [get]
// @Security ApiKeyAuth
func (apiCfg *apiConfig) handlGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Couldn't retrieve user posts: %v", err))
		return
	}
	responWithJson(w, http.StatusOK, databasePostsToPosts(posts))
}
