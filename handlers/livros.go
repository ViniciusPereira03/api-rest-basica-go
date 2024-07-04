package handlers

import (
	"api/models"
	"api/services"
	"encoding/json"
	"net/http"
)

func GetLivros(w http.ResponseWriter, r *http.Request) {
	livros := services.GetLivros()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livros)
}

func AddLivro(w http.ResponseWriter, r *http.Request) {
	var data models.Livro
	json.NewDecoder(r.Body).Decode(&data)
	livro := services.AddLivro(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livro)
}
