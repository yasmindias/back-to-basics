package model

import "time"

type Request struct {
	Nome               string    `json:"nome"`
	Cpf                string    `json:"cpf"`
	Cargo              Cargo     `json:"cargo"`
	Salario            float64   `json:"salario"`
	DataUltimoReajuste time.Time `json:"data_ultimo_reajuste"`
	ValorReajuste      float64   `json:"valor_reajuste"`
}
