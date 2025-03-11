package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//Marshal the payload to json

	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("JSON Parsing error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Println("Error while returning the data", err)
	}
}
