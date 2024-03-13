package model

type Orcamento struct {
	valor int64
}

func (o *Orcamento) GetValor() int64 {
	return o.valor
}

func (o *Orcamento) SetValor(novoValor int64) {
	o.valor = novoValor
}
