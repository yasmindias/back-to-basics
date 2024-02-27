package validacao

import (
	"fmt"
	"rh/cmd/model"
	"time"
)

const (
	segundosEmUmMes = 2628288
)

type ValidacaoPeriodicidade struct{}

func (vpr ValidacaoPeriodicidade) Validate(value ...interface{}) error {
	f := value[0].(*model.Funcionario)

	ultimoReajuste := f.DataUltimoReajuste
	tempoPassado := time.Since(ultimoReajuste)
	mesesDesdeUltimoReajuste := tempoPassado.Seconds() / segundosEmUmMes

	if mesesDesdeUltimoReajuste < 6 {
		return fmt.Errorf("intervalo entre reajustes deve ser de no mínimo 6 meses")
	}

	return nil
}
