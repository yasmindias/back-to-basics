package internal_test

import (
	"fmt"
	"rh/cmd/model"
	internal "rh/internal/service/reajuste"
	"rh/internal/service/reajuste/validacao"
	"testing"
	"time"
)

var validacoes = []validacao.ValidacaoReajuste{
	validacao.ValidacaoPercentual{},
	validacao.ValidacaoPeriodicidade{},
}

func Test_ReajustePercentualSucesso(t *testing.T) {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  "John Doe",
			Cpf:   "01234567899",
			Cargo: model.Assistente,
		},
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	error := internal.ReajustarSalario(&f, 400, validacoes)
	if error != nil {
		fmt.Printf("error: %v \n", error.Error())
	}
}

func Test_ReajustePercentualErro(t *testing.T) {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  "John Doe",
			Cpf:   "01234567899",
			Cargo: model.Assistente,
		},
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	want := fmt.Errorf("reajuste nao pode ser superior a 40%% do salario")
	got := internal.ReajustarSalario(&f, 500, validacoes)
	if got.Error() != want.Error() {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_ReajustePeriodicidadeSucesso(t *testing.T) {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  "John Doe",
			Cpf:   "01234567899",
			Cargo: model.Assistente,
		},
		DataUltimoReajuste: time.Now().AddDate(0, -3, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	error := internal.ReajustarSalario(&f, 300, validacoes)
	if error != nil {
		fmt.Printf("error: %v \n", error.Error())
	}
}

func Test_ReajustePeriodicidadeErro(t *testing.T) {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  "John Doe",
			Cpf:   "01234567899",
			Cargo: model.Assistente,
		},
		DataUltimoReajuste: time.Now().AddDate(0, -3, 0),
	}
	f.AtualizarSalario(1000) //valor inicial do salario

	want := fmt.Errorf("intervalo entre reajustes deve ser de no mínimo 6 meses")
	got := internal.ReajustarSalario(&f, 400, validacoes)
	if got.Error() != want.Error() {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}
