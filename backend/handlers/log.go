package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"aqua-nav/models"
)

func LogData(w http.ResponseWriter, r *http.Request) {
	var data models.LogEntry
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: Save the data to a database
	// For now just printing the LogData

	log.Println("recieved data", data)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Log saved"})
}
