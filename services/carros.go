package services

import "api/models"

var carros []models.Carro

func GetCarros() []models.Carro {
	var carro1 models.Carro
	var carro2 models.Carro

	if len(carros) == 0 {
		carro1.Marca = "Chevrolet"
		carro1.Modelo = "Agile LT 1.4 MPFI 8V flex"
		carro1.Valor = 21500
		carro1.UnidadesEmEstoque = 5
		carro1.Ano = 2010
		carro1.Categoria = "carro"
		carros = append(carros, carro1)

		carro2.Marca = "Dodge"
		carro2.Modelo = "Dakota Durango SLT 5.2 4x4 V8"
		carro2.Valor = 28470
		carro2.UnidadesEmEstoque = 2
		carro2.Ano = 1998
		carro2.Categoria = "carro"
		carros = append(carros, carro2)
	}

	return carros
}

func AddCarro(carro models.Carro) models.Carro {
	carros = append(carros, carro)
	return carro
}
