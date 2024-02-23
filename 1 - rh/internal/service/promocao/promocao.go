package promocao

import (
	"fmt"
	"rh/cmd"
)

func Promover(f *cmd.Funcionario, metaBatida bool) error {
	if f.Cargo == cmd.Gerente {
		return fmt.Errorf("gerentes nao podem ser promovidos")
	}

	if metaBatida {
		novoCargo := f.Cargo.ProximoCargo()
		f.Promover(novoCargo)
	} else {
		return fmt.Errorf("funcionário não bateu a meta")
	}

	return nil
}
