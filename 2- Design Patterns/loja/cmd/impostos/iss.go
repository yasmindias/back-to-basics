package impostos

import "design-patterns-loja/cmd/model"

type ISS struct{}

func (iss ISS) Calcular(value ...interface{}) float64 {
	orcamento := value[0].(model.Orcamento)

	return float64(orcamento.GetValor()) * 0.06
}
