package models

type Empresa struct {
	Nome              string `json:"nome"`
	CNPJ              string `json:"cnpj"`
	InscricaoEstadual string `json:"inscricao_estadual"`
	DataAbertura      string `json:"data_abertura"`
	CEP               string `json:"cep"`
	Endereco          string `json:"endereco"`
}
