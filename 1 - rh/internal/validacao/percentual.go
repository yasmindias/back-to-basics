package validacao

import (
	"errors"
	"math"
	"rh/cmd"
)

type ValidacaoPercentual struct{}

func (vpct ValidacaoPercentual) Validate(value ...interface{}) error {
	f := value[0].(*cmd.Funcionario)
	aumento := value[1].(float64)

	percentualReajuste := roundFloat(aumento/f.GetSalario(), 1)
	if percentualReajuste > 0.4 {
		return errors.New("reajuste nao pode ser superior a 40%% do salario")
	}

	return nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
