package services

import "api/models"

var empresas []models.Empresa

func GetEmpresas() []models.Empresa {
	var empresa1 models.Empresa
	var empresa2 models.Empresa

	if len(empresas) == 0 {
		empresa1.Nome = "Rosângela e Vitória Buffet Ltda"
		empresa1.CNPJ = "98.569.066/0001-00"
		empresa1.InscricaoEstadual = "125.760.542.699"
		empresa1.DataAbertura = "08/11/2019"
		empresa1.CEP = "17511-263"
		empresa1.Endereco = "Praça Ricardo Gracindo"
		empresas = append(empresas, empresa1)

		empresa2.Nome = "Márcia e Caroline Doces & Salgados Ltda"
		empresa2.CNPJ = "45.950.188/0001-77"
		empresa2.InscricaoEstadual = "354.713.355.315"
		empresa2.DataAbertura = "23/10/2019"
		empresa2.CEP = "15706-208"
		empresa2.Endereco = "Rua dos Sabiás"
		empresas = append(empresas, empresa2)
	}

	return empresas
}

func AddEmpresa(empresa models.Empresa) models.Empresa {
	empresas = append(empresas, empresa)
	return empresa
}
