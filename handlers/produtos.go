package handlers

import (
	"api/models"
	"api/services"
	"encoding/json"
	"net/http"
)

func GetProdutos(w http.ResponseWriter, r *http.Request) {
	produtos := services.GetProdutos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func AddProduto(w http.ResponseWriter, r *http.Request) {
	var data models.Produto
	json.NewDecoder(r.Body).Decode(&data)
	produto := services.AddProduto(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}
