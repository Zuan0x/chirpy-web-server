package main

import (
	"net/http"
	"encoding/json"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		UserID int `json:"user_id"`
	}

	type parameters struct {
		Event string `json:"event"`
		Data Data `json:"data"`
	}


	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	if params.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusOK, "Not a user_created event")
		return
	}

	user, err := cfg.DB.UpgradeUser(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't upgrade user")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
