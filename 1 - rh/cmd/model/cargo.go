package model

import (
	"encoding/json"
	"strings"
)

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

func (c Cargo) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Cargo) UnmarshalJSON(data []byte) (err error) {
	var cargo string
	if err := json.Unmarshal(data, &cargo); err != nil {
		return err
	}
	if *c, err = ParseCargo(cargo); err != nil {
		return err
	}
	return nil
}

func ParseCargo(cargo string) (Cargo, error) {
	cargo = strings.TrimSpace(strings.ToLower(cargo))
	switch cargo {
	case "assistente":
		return Assistente, nil
	case "analista":
		return Analista, nil
	case "especialista":
		return Especialista, nil
	case "gerente":
		return Gerente, nil
	}
	return Assistente, nil
}
