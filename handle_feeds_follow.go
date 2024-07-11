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
func (apiCfg *apiConfig) handleCreateFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramenters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	param := paramenters{}
	err := decoder.Decode(&param)

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    param.FeedID,
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Couldn't create a follow the feed%v", err))
		return
	}

	responWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}
