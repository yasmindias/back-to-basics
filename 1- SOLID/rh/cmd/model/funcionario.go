package model

import (
	"time"
)

type Funcionario struct {
	DadosPessoais      DadosPessoais `json:"dados_pessoais"`
	DataUltimoReajuste time.Time     `json:"data_ultimo_reajuste"`
}

func (f *Funcionario) AtualizarSalario(aumento float64) {
	f.DadosPessoais.salario += aumento
}

func (f *Funcionario) Promover(novoCargo Cargo) {
	f.DadosPessoais.Cargo = novoCargo
}
