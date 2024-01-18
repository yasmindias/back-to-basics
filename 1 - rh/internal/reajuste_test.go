package internal_test

import (
	"fmt"
	"rh/cmd"
	"rh/internal"
	"testing"
	"time"
)

func Test_ReajustePercentualErro(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Assistente,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	errors := internal.ReajustarSalario(&f, 500)
	if len(errors) > 0 {
		t.Errorf("got %v errors \n", len(errors))
		for _, err := range errors {
			fmt.Printf("error: %v \n", err)
		}
	}
}

func Test_ReajustePeriodicidadeSucesso(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Assistente,
		DataUltimoReajuste: time.Now().AddDate(0, -3, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	errors := internal.ReajustarSalario(&f, 400)
	if len(errors) > 0 {
		t.Errorf("got %v errors \n", len(errors))
		for _, err := range errors {
			fmt.Printf("error: %v \n", err)
		}
	}
}
