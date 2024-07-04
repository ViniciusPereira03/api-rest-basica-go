package handlers

import (
	"api/models"
	"api/services"
	"encoding/json"
	"net/http"
)

func GetEmpresas(w http.ResponseWriter, r *http.Request) {
	empresas := services.GetEmpresas()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empresas)
}

func AddEmpresa(w http.ResponseWriter, r *http.Request) {
	var data models.Empresa
	json.NewDecoder(r.Body).Decode(&data)
	empresa := services.AddEmpresa(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empresa)
}
