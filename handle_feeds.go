package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/sir-george2500/g-server/internal/auth"
	"github.com/sir-george2500/g-server/internal/database"
)

// handle respon with user
func (apiCfg *apiConfig) handleCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramenters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	param := paramenters{}
	err := decoder.Decode(&param)

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
		Url:       param.URL,
		UserID:    user.ID,
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Couldn't create feed%v", err))
		return
	}

	responWithJson(w, 201, databaseFeedToFeed(feed))
}

// handle respon with user
func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Could not get feed%v", err))
		return
	}

	responWithJson(w, 200, databaseFeedToFeeds(feeds))
}
