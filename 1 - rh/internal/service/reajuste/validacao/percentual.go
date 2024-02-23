package validacao

import (
	"fmt"
	"math"
	"rh/cmd/model"
)

type ValidacaoPercentual struct{}

func (vpct ValidacaoPercentual) Validate(value ...interface{}) error {
	f := value[0].(*model.Funcionario)
	aumento := value[1].(float64)
	percentualReajuste := roundFloat(aumento/f.DadosPessoais.GetSalario(), 1)
	if percentualReajuste > 0.4 {
		return fmt.Errorf("reajuste nao pode ser superior a 40%% do salario")
	}

	return nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
