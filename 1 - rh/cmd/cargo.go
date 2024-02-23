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

func (c Cargo) ProximoCargo() Cargo {
	switch c {
	case Assistente:
		return Analista
	case Analista:
		return Especialista
	case Especialista:
		return Gerente
	}
	return Gerente
}
