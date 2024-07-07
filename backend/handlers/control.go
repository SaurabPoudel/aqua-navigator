package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ControlCommand struct {
	Command string `json:"command"`
}

func ControlBoat(w http.ResponseWriter, r *http.Request) {
	var cmd ControlCommand
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: Implement sending command to Arduino
	// For now, just printing the command
	log.Println("Recieved command:", cmd.Command)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Command received"})
}
