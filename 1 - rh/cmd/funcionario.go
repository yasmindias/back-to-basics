package cmd

import (
	"errors"
	"math"
	"time"
)

type Funcionario struct {
	Nome               string
	Cpf                string
	Cargo              Cargo
	Salario            float64
	DataUltimoReajuste time.Time
}

func (f *Funcionario) ReajustarSalario(aumento float64) error {
	percentualReajuste := roundFloat(aumento/f.Salario, 1)
	if percentualReajuste > 0.4 {
		return errors.New("reajuste nao pode ser superior a 40% do salario")
	}

	f.Salario += aumento
	f.DataUltimoReajuste = time.Now()

	return nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
