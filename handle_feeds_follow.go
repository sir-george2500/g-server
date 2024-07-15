package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	_ "github.com/sir-george2500/g-server/internal/auth"
	"github.com/sir-george2500/g-server/internal/database"
)

type paramenters struct {
	FeedID uuid.UUID `json:"feed_id"`
}

// handleCreateFeedsFollow creates a new feed follow
// @Summary Create Feed Follow
// @Description Create a new feed follow for the authenticated user
// @Tags feed_follows
// @Accept  json
// @Produce  json
// @Param feed body paramenters true "Feed Follow Parameters"
// @Success 201 {object} database.FeedsFollow
// @Failure 400 {string} string "Error message"
// @Router /feed_follows [post]
// @Security ApiKeyAuth
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

// handleGetFeedsFollow retrieves all feed follows for the authenticated user
// @Summary Get Feed Follows
// @Description Retrieve all feed follows for the authenticated user
// @Tags feed_follows
// @Accept  json
// @Produce  json
// @Success 200 {array} database.FeedsFollow
// @Failure 400 {string} string "Error message"
// @Router /feed_follows [get]
// @Security ApiKeyAuth
func (apiCfg *apiConfig) handleGetFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Couldn't get feed follow feed%v", err))
		return
	}

	responWithJson(w, 201, databaseFeedFollowsToFeedFollows(feedFollow))
}

// handleDeleteFeedFollow deletes a feed follow
// @Summary Delete Feed Follow
// @Description Delete a feed follow for the authenticated user
// @Tags feed_follows
// @Accept  json
// @Produce  json
// @Param feedFollowID path string true "Feed Follow ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string "Error message"
// @Router /feed_follows/{feedFollowID} [delete]
// @Security ApiKeyAuth
func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Cann't parse feed to UUID%v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Can't delete feed follow%v", err))
		return
	}
	responWithJson(w, 200, struct{}{})
}
