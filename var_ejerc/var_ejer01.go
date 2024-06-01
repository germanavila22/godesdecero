package varejerc

import (
	"strconv"
)

func Mifuncion(numero string) (int, string) {
	var text_func string
	numero_entero, mierror := strconv.Atoi(numero)
	if mierror != nil {
		return 0, "hay un error" + mierror.Error()
	}
	if numero_entero > 100 {
		text_func = "es mayor a 100"
		//return numero_entero, "es mayor a 100"
	} else {
		text_func = "es menor a 100"
		//return numero_entero, "es menor a 100"
	}
	return numero_entero, text_func

}
