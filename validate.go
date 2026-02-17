package main

import (
	"log"
	"net/http"
	"encoding/json"
)



func handlerValidate (w http.ResponseWriter, r *http.Request) {
	
	type validateJson struct {
		Body string `json:body`
	}

	decoder := json.NewDecoder(r.Body)
	params := validateJson{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		respondWithError(w, 500, "Error decoding parameters")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp too long")
		return 
	}

	filteredChirp := filterWords(params.Body)


	respondWithJSON(w, 200, filteredChirp)
}

func respondWithError(w http.ResponseWriter, code int, errorDesc string) {

	type returnErrors struct{
 		Error string `json:"error"`
	}

	respBody := returnErrors{
		Error: errorDesc,
	}

	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	type returnVals struct {
		Cleaned_Body interface{} `json:"cleaned_body"`

	}

	respBody := returnVals{
		Cleaned_Body: payload,
	}

	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
