package validacao

type ValidacaoReajuste interface {
	Validate(value ...interface{}) error
}
