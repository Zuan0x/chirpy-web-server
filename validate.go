package main

import (
	"net/http"
	"encoding/json"
	"log"
)


func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
type parameters struct {
        Body string `json:"body"`
        error `json:"error"`
    }

	type returnVals struct {
		Valid bool `json:"valid"`
	}

	log.Printf("Validating chirp", r)
    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")		
		return
    }
	const maxChirpLength = 140

	if len(params.Body) == 0 {
		respondWithError(w, http.StatusBadRequest, "Chirp is empty")
		return
	}

	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}	
	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
	})
	return
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
