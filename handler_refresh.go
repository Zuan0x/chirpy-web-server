package main

import (
	"net/http"
	"github.com/Zuan0x/chirpy-web-server/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	isRevoked, err := cfg.DB.IsTokenRevoked(token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get token")
		return
	}

	if isRevoked {
		respondWithError(w, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	accessToken, err := auth.RefreshJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't refresh token")
		return
	}

	respondWithJSON(w, http.StatusOK, response{Token: accessToken})
}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshtoken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	err = cfg.DB.RevokeToken(refreshtoken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke token")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
