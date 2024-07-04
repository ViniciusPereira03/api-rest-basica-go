package handlers

import (
	"api/models"
	"api/services"
	"encoding/json"
	"net/http"
)

func GetPessoas(w http.ResponseWriter, r *http.Request) {
	pessoas := services.GetPessoas()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pessoas)
}

func AddPessoa(w http.ResponseWriter, r *http.Request) {
	var data models.Pessoa
	json.NewDecoder(r.Body).Decode(&data)
	pessoa := services.AddPessoa(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pessoa)
}
