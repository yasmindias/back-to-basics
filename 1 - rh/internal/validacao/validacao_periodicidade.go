package validacao

import (
	"fmt"
	"time"
)

const (
	segundosEmUmMes = 2628288
)

func ValidarPeriodo(periodoMinimo float64) Regra {
	return func(value []interface{}) error {
		dataUltimoReajuste, ok := value[0].(time.Time)
		if !ok {
			return fmt.Errorf("data invalida")
		}

		ultimoReajuste := dataUltimoReajuste
		tempoPassado := time.Since(ultimoReajuste)
		mesesDesdeUltimoReajuste := tempoPassado.Seconds() / segundosEmUmMes

		if mesesDesdeUltimoReajuste < periodoMinimo {
			return fmt.Errorf("intervalo entre reajustes deve ser de no mínimo %v meses", periodoMinimo)
		}
		return nil
	}
}
