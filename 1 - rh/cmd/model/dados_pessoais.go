package model

type DadosPessoais struct {
	Nome    string
	Cpf     string
	Cargo   Cargo
	salario float64
}

func (dp *DadosPessoais) GetSalario() float64 {
	return dp.salario
}
