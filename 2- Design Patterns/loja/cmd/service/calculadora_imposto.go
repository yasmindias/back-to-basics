package service

import (
	"design-patterns-loja/cmd/impostos"
	"design-patterns-loja/cmd/model"
)

func Calcular(orcamento model.Orcamento, imposto impostos.Imposto) float64 {
	return imposto.Calcular(orcamento)
}
