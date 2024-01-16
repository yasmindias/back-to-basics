package internal

import (
	"errors"
	"math"
	"rh/cmd"
	"time"
)

func ReajustarSalario(f *cmd.Funcionario, aumento float64) error {
	percentualReajuste := roundFloat(aumento/f.GetSalario(), 1)
	if percentualReajuste > 0.4 {
		return errors.New("reajuste nao pode ser superior a 40% do salario")
	}

	f.AtualizarSalario(aumento)
	f.DataUltimoReajuste = time.Now()

	return nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
