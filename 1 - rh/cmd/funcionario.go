package cmd

import (
	"time"
)

type Funcionario struct {
	Nome               string
	Cpf                string
	Cargo              Cargo
	salario            float64
	DataUltimoReajuste time.Time
}

func (f *Funcionario) GetSalario() float64 {
	return f.salario
}

func (f *Funcionario) AtualizarSalario(aumento float64) {
	f.salario += aumento
}

func (f *Funcionario) Promover(novoCargo Cargo) {
	f.Cargo = novoCargo
}
