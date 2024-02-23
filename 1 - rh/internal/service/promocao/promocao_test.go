package promocao_test

import (
	"rh/cmd"
	"testing"
	"time"
)

func Test_PromocaoAssistente(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Assistente,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.Cargo.ProximoCargo())

	want := cmd.Analista
	got := f.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoAnalista(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Analista,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.Cargo.ProximoCargo())

	want := cmd.Especialista
	got := f.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoEspecialista(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Especialista,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.Cargo.ProximoCargo())

	want := cmd.Gerente
	got := f.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}

func Test_PromocaoGerente_Error(t *testing.T) {
	f := cmd.Funcionario{
		Nome:               "John Doe",
		Cpf:                "01234567899",
		Cargo:              cmd.Gerente,
		DataUltimoReajuste: time.Now().AddDate(-1, 0, 0),
	}
	f.Promover(f.Cargo.ProximoCargo())

	want := cmd.Gerente
	got := f.Cargo

	if got != want {
		t.Errorf("got: %v \n wanted %v", got, want)
	}
}
