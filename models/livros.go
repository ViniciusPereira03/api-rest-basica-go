package models

type Livro struct {
	Titulo    string `json:"titulo"`
	Categoria string `json:"categoria"`
	ISBN      string `json:"isbn"`
	Ano       string `json:"ano"`
	Paginas   string `json:"paginas"`
	Preco     int    `json:"preco"`
	Estoque   string `json:"estoque"`
}
