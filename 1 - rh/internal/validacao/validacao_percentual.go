package validacao

import (
	"errors"
	"fmt"
	"math"
)

func ValidarPercentual(percentualMaximo float64) Regra {
	return func(value []interface{}) error {
		salario, ok := value[0].(float64)
		if !ok {
			return fmt.Errorf("salario nao é um numero")
		}
		aumento, ok := value[1].(float64)
		if !ok {
			return fmt.Errorf("aumento nao é um numero")
		}

		percentualReajuste := roundFloat(aumento/salario, 1)
		if percentualReajuste > percentualMaximo {
			return errors.New("reajuste nao pode ser superior a 40% do salario")
		}

		return nil
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
