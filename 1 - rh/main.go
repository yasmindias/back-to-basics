package main

import (
	"encoding/json"
	"net/http"
	"rh/cmd/model"
	internal "rh/internal/service/reajuste"
	"rh/internal/service/reajuste/validacao"
)

func main() {
	http.HandleFunc("/reajuste", reajuste)

	http.ListenAndServe(":8090", nil)
}

func reajuste(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		validacoes := []validacao.ValidacaoReajuste{
			validacao.ValidacaoPercentual{},
			validacao.ValidacaoPeriodicidade{},
		}

		body := model.Request{}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		f := parseFuncionario(body)

		var resp []byte
		err := internal.ReajustarSalario(&f, body.ValorReajuste, validacoes)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			resp, _ = json.Marshal("Salario Reajustado")
			w.Write(resp)
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func parseFuncionario(body model.Request) model.Funcionario {
	f := model.Funcionario{
		DadosPessoais: model.DadosPessoais{
			Nome:  body.Nome,
			Cpf:   body.Cpf,
			Cargo: body.Cargo,
		},
		DataUltimoReajuste: body.DataUltimoReajuste,
	}
	f.DadosPessoais.SetSalario(body.Salario)

	return f
}
