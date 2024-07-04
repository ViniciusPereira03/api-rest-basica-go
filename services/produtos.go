package services

import "api/models"

var produtos []models.Produto

func GetProdutos() []models.Produto {
	var produto1 models.Produto
	var produto2 models.Produto

	if len(produtos) == 0 {
		produto1.ID = 1
		produto1.Name = "Avental"
		produto1.Description = "Jaleco de alta qualidade fabricado para atender aos clientes mais exigentes"
		produto1.Price = "R$ 999,99"
		produto1.Available = true
		produtos = append(produtos, produto1)

		produto2.ID = 2
		produto2.Name = "Touca"
		produto2.Description = "Jaleco de alta qualidade fabricado para atender aos clientes mais exigentes"
		produto2.Price = "R$ 999,99"
		produto2.Available = true
		produtos = append(produtos, produto2)
	}

	return produtos
}

func AddProduto(product models.Produto) models.Produto {
	produtos = append(produtos, product)
	return product
}
