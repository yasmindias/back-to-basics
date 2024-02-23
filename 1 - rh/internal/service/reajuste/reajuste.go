package internal

import (
	"rh/cmd/model"
	"rh/internal/service/reajuste/validacao"
	"time"
)

func validarDadosReajuste(f *model.Funcionario, aumento float64, validacoes []validacao.ValidacaoReajuste) []error {
	var errors []error
	for _, validacao := range validacoes {
		if err := validacao.Validate(f, aumento); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func ReajustarSalario(f *model.Funcionario, aumento float64, validacoes []validacao.ValidacaoReajuste) error {
	errors := validarDadosReajuste(f, aumento, validacoes)
	if len(errors) > 0 {
		return errors[len(errors)-1]
	}

	f.AtualizarSalario(aumento)
	f.DataUltimoReajuste = time.Now()

	return nil
}
