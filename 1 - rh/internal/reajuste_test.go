package internal_test

import (
	"fmt"
	"rh/cmd"
	"rh/internal"
	"testing"
	"time"
)

func Test_ReajustePercentualSucesso(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Assistente,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	error := internal.ReajustarSalario(&f, 400)
	if error != nil {
		fmt.Printf("error: %v \n", error.Error())
	}
}

func Test_ReajustePercentualErro(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Assistente,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	want := fmt.Errorf("reajuste nao pode ser superior a 40%% do salario")
	got := internal.ReajustarSalario(&f, 500)
	if got.Error() != want.Error() {
		t.Errorf("got: %v \n wanted %v", got, want)
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

	error := internal.ReajustarSalario(&f, 300)
	if error != nil {
		fmt.Printf("error: %v \n", error.Error())
	}
}

func Test_ReajustePeriodicidadeErro(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Assistente,
		DataUltimoReajuste: time.Now().AddDate(0, -3, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	want := fmt.Errorf("intervalo entre reajustes deve ser de no mínimo 6 meses")
	got := internal.ReajustarSalario(&f, 400)
	if got.Error() != want.Error() {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}
