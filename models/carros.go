package models

type Carro struct {
	Marca             string `json:"marca"`
	Modelo            string `json:"modelo"`
	Ano               int    `json:"ano"`
	Valor             int    `json:"valor"`
	UnidadesEmEstoque int    `json:"unidadesEmEstoque"`
	Categoria         string `json:"categoria"`
}
