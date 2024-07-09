package main

import (
	"fmt"
	"net/http"

	"github.com/sir-george2500/g-server/internal/auth"
	"github.com/sir-george2500/g-server/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responWithError(w, 403, fmt.Sprintf("Auth Header %v", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responWithError(w, 403, fmt.Sprintf("Couldn't get user %v", err))
			return
		}

		handler(w, r, user)
	}
}
