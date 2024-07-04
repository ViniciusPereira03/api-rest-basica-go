package handlers

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	response := "API Rest básica em Go"

	json.NewEncoder(w).Encode(response)
}
