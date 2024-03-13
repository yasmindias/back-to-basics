package service

import "design-patterns-loja/cmd/model"

const (
	AliquotaICMS = 0.1
	AliquotaISS  = 0.06
)

func Calcular(orcamento model.Orcamento, tipoImposto model.Imposto) float64 {
	switch tipoImposto {
	case model.ICMS:
		return float64(orcamento.GetValor()) * AliquotaICMS
	case model.ISS:
		return float64(orcamento.GetValor()) * AliquotaISS
	}
	return 0
}
