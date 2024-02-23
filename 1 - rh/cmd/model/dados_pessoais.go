package model

type DadosPessoais struct {
	Nome    string  `json:"nome"`
	Cpf     string  `json:"cpf"`
	Cargo   Cargo   `json:"cargo"`
	salario float64 `json:"salario"`
}

func (dp *DadosPessoais) GetSalario() float64 {
	return dp.salario
}

func (dp *DadosPessoais) SetSalario(salario float64) {
	dp.salario = salario
}
