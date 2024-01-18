package internal

import (
	"rh/cmd"
	"rh/internal/validacao"
)

func ReajustarSalario(f *cmd.Funcionario, aumento float64) []error {
	data := map[string][]interface{}{
		"percentual":    {f.GetSalario(), aumento},
		"periodicidade": {f.DataUltimoReajuste},
	}

	validator := criarValidatorComRegras()
	errors := validator.Validar(data)
	return errors
}

func criarValidatorComRegras() validacao.Validator {
	validator := validacao.Validator{}
	validator.Add("percentual", validacao.ValidarPercentual(0.4))
	validator.Add("periodicidade", validacao.ValidarPeriodo(6))

	return validator
}
