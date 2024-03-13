package service

import "design-patterns-loja/cmd/model"

const (
	AliquotaICMS = 0.1
)

func Calcular(orcamento model.Orcamento) float64 {
	return float64(orcamento.GetValor()) * AliquotaICMS
}
