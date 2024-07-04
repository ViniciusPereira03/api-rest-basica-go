package services

import "api/models"

var livros []models.Livro

func GetLivros() []models.Livro {
	var livro1 models.Livro
	var livro2 models.Livro

	if len(livros) == 0 {
		livro1.Titulo = "Primeiros passos com a linguagem Rust"
		livro1.Categoria = "Programação"
		livro1.ISBN = "978-85-7522-683-4"
		livro1.Ano = "2018"
		livro1.Paginas = "312"
		livro1.Preco = 69
		livro1.Estoque = "Disponível também em formato ebook"
		livros = append(livros, livro1)

		livro2.Titulo = "Programação em Baixo Nível"
		livro2.Categoria = "Hardware &amp; Robótica"
		livro2.ISBN = "978-85-7522-667-4"
		livro2.Ano = "2018"
		livro2.Paginas = "576"
		livro2.Preco = 120
		livro2.Estoque = "Disponível também em formato ebook"
		livros = append(livros, livro2)
	}

	return livros
}

func AddLivro(livro models.Livro) models.Livro {
	livros = append(livros, livro)
	return livro
}
