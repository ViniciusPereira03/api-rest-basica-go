package services

import "api/models"

var pessoas []models.Pessoa

func GetPessoas() []models.Pessoa {
	var pessoa1 models.Pessoa
	var pessoa2 models.Pessoa

	if len(pessoas) == 0 {
		pessoa1.Nome = "Natália Cristiane Lívia Aragão"
		pessoa1.Idade = 68
		pessoa1.CPF = "489.591.511-57"
		pessoa1.RG = "45.833.906-4"
		pessoa1.DataNasc = "11/01/1956"
		pessoa1.Sexo = "Feminino"
		pessoa1.Signo = "Capricórnio"
		pessoa1.Mae = "Carolina Bianca Alice"
		pessoa1.Pai = "Caio Luan Breno Aragão"
		pessoa1.Email = "nataliacristianearagao@cteep.com.br"
		pessoa1.Senha = "o8jbpFRZZi"
		pessoa1.CEP = "89234-015"
		pessoa1.Endereco = "Rua Antônio João de Borba"
		pessoa1.Numero = 771
		pessoa1.Bairro = "Paranaguamirim"
		pessoa1.Cidade = "Joinville"
		pessoa1.Estado = "SC"
		pessoa1.TelefoneFixo = "(47) 3683-1989"
		pessoa1.Celular = "(47) 99211-1246"
		pessoa1.Altura = "1,60"
		pessoa1.Peso = 72
		pessoa1.TipoSanguineo = "A+"
		pessoa1.Cor = "amarelo"
		pessoas = append(pessoas, pessoa1)

		pessoa2.Nome = "Nina Teresinha Nascimento"
		pessoa2.Idade = 61
		pessoa2.CPF = "057.781.444-31"
		pessoa2.RG = "11.786.760-3"
		pessoa2.DataNasc = "26/05/1963"
		pessoa2.Sexo = "Feminino"
		pessoa2.Signo = "Gêmeos"
		pessoa2.Mae = "Marina Rayssa Emanuelly"
		pessoa2.Pai = "Kauê Diogo Nascimento"
		pessoa2.Email = "nina_teresinha_nascimento@studiodeideias.com"
		pessoa2.Senha = "jOLK6ZNDGN"
		pessoa2.CEP = "65082-280"
		pessoa2.Endereco = "Avenida José Sarney"
		pessoa2.Numero = 567
		pessoa2.Bairro = "Vila Ariri"
		pessoa2.Cidade = "São Luís"
		pessoa2.Estado = "MA"
		pessoa2.TelefoneFixo = "(98) 2702-0212"
		pessoa2.Celular = "(98) 98236-8024"
		pessoa2.Altura = "1,83"
		pessoa2.Peso = 52
		pessoa2.TipoSanguineo = "AB+"
		pessoa2.Cor = "azul"
		pessoas = append(pessoas, pessoa2)
	}

	return pessoas
}

func AddPessoa(pessoa models.Pessoa) models.Pessoa {
	pessoas = append(pessoas, pessoa)
	return pessoa
}
