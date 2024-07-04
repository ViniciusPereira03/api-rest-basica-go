package main

import (
	"api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index).Methods("GET")

	r.HandleFunc("/produtos", handlers.GetProdutos).Methods("GET")
	r.HandleFunc("/produtos", handlers.AddProduto).Methods("POST")

	r.HandleFunc("/carros", handlers.GetCarros).Methods("GET")
	r.HandleFunc("/carros", handlers.AddCarro).Methods("POST")

	r.HandleFunc("/empresas", handlers.GetEmpresas).Methods("GET")
	r.HandleFunc("/empresas", handlers.AddEmpresa).Methods("POST")

	r.HandleFunc("/livros", handlers.GetLivros).Methods("GET")
	r.HandleFunc("/livros", handlers.AddLivro).Methods("POST")

	r.HandleFunc("/pessoas", handlers.GetPessoas).Methods("GET")
	r.HandleFunc("/pessoas", handlers.AddPessoa).Methods("POST")

	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Home Page!"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
