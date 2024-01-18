package validacao

type Regra func(value []interface{}) error

type Regras map[string]Regra

type Validator struct {
	regras Regras
}

func (v *Validator) Add(key string, regra Regra) {
	if v.regras == nil {
		v.regras = make(map[string]Regra)
	}
	v.regras[key] = regra
}

func (v *Validator) Validar(data map[string][]interface{}) []error {
	var errors []error
	for keyRegra, regra := range v.regras {
		input := data[keyRegra]
		if input != nil {
			if err := regra(data[keyRegra]); err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}
