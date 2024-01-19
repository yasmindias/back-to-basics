package internal

import (
	"fmt"
	"rh/cmd"
	"rh/internal/validacao"
	"strings"
)

func validarDadosReajuste(f *cmd.Funcionario, aumento float64) []error {
	validacoes := []validacao.ValidacaoReajuste{
		validacao.ValidacaoPercentual{},
		validacao.ValidacaoPeriodicidade{},
	}

	var errors []error
	for _, validacao := range validacoes {
		if err := validacao.Validate(f, aumento); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func ReajustarSalario(f *cmd.Funcionario, aumento float64) error {
	errors := validarDadosReajuste(f, aumento)
	if len(errors) > 0 {
		var strErr string
		for _, err := range errors {
			strErr += err.Error() + "\n"
		}
		return fmt.Errorf(strings.Trim(strErr, "\n"))
	}

	f.AtualizarSalario(aumento)

	return nil
}
