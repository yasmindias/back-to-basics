package cmd

type Cargo int64

const (
	Assistente Cargo = iota
	Analista
	Especialista
	Gerente
)

func (c Cargo) String() string {
	switch c {
	case Assistente:
		return "assistente"
	case Analista:
		return "analista"
	case Especialista:
		return "especialista"
	case Gerente:
		return "gerente"
	}
	return "unknown"
}
