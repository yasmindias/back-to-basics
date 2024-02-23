package promocao_test

import (
	"rh/cmd/model"
	"testing"
	"time"
)

func Test_PromocaoAssistente(t *testing.T) {
	f := model.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              model.Assistente,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.Cargo.ProximoCargo())

	want := model.Analista
	got := f.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoAnalista(t *testing.T) {
	f := model.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              model.Analista,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.Cargo.ProximoCargo())

	want := model.Especialista
	got := f.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoEspecialista(t *testing.T) {
	f := model.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              model.Especialista,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.Cargo.ProximoCargo())

	want := model.Gerente
	got := f.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoGerente_Error(t *testing.T) {
	f := model.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              model.Gerente,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.Cargo.ProximoCargo())

	want := model.Gerente
	got := f.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}
