package handlers

import (
	"api/models"
	"api/services"
	"encoding/json"
	"net/http"
)

func GetCarros(w http.ResponseWriter, r *http.Request) {
	carros := services.GetCarros()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carros)
}

func AddCarro(w http.ResponseWriter, r *http.Request) {
	var data models.Carro
	json.NewDecoder(r.Body).Decode(&data)
	carro := services.AddCarro(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carro)
}
