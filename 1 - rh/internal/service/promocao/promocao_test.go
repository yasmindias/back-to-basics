package promocao_test

import (
	"rh/cmd/model"
	"testing"
	"time"
)

func Test_PromocaoAssistente(t *testing.T) {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  "John Doe",
			Cpf:   "01234567899",
			Cargo: model.Assistente,
		},
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.DadosPessoais.Cargo.ProximoCargo())

	want := model.Analista
	got := f.DadosPessoais.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoAnalista(t *testing.T) {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  "John Doe",
			Cpf:   "01234567899",
			Cargo: model.Assistente,
		},
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.DadosPessoais.Cargo.ProximoCargo())

	want := model.Especialista
	got := f.DadosPessoais.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoEspecialista(t *testing.T) {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  "John Doe",
			Cpf:   "01234567899",
			Cargo: model.Assistente,
		},
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.DadosPessoais.Cargo.ProximoCargo())

	want := model.Gerente
	got := f.DadosPessoais.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoGerente_Error(t *testing.T) {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  "John Doe",
			Cpf:   "01234567899",
			Cargo: model.Assistente,
		},
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.DadosPessoais.Cargo.ProximoCargo())

	want := model.Gerente
	got := f.DadosPessoais.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}
