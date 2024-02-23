package model

import (
	"time"
)

type Funcionario struct {
	DadosPessoais      DadosPessoais
	DataUltimoReajuste time.Time
}

func (f *Funcionario) AtualizarSalario(aumento float64) {
	f.DadosPessoais.salario += aumento
}

func (f *Funcionario) Promover(novoCargo Cargo) {
	f.DadosPessoais.Cargo = novoCargo
}
