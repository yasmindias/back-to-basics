package service

import (
	"design-patterns-loja/cmd/model"
	"testing"
)

func TestCalcular(t *testing.T) {
	type args struct {
		orcamento   model.Orcamento
		tipoImposto model.Imposto
	}

	orcamento := model.Orcamento{}
	orcamento.SetValor(100)

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ICMS",
			args: args{orcamento, model.ICMS},
			want: 10,
		},
		{
			name: "ISS",
			args: args{orcamento, model.ISS},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calcular(tt.args.orcamento, tt.args.tipoImposto); got != tt.want {
				t.Errorf("Calcular() = %v, want %v", got, tt.want)
			}
		})
	}
}
