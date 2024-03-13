package impostos

type Imposto interface {
	Calcular(value ...interface{}) float64
}
