package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Sueldo64 float64
var Fecha time.Time

func Restovariables() {
	Nombre = "Petro"
	Estado = true
	Sueldo = 15.55
	Sueldo64 = 15665.5
	Fecha = time.Now()

	fmt.Println("el nombre es ", Nombre)
	fmt.Println("el estado es ", Estado)
	fmt.Println("el sueldo es ", Sueldo)
	fmt.Println("el sueldo 64 es ", Sueldo64)
	fmt.Println("el fecha es ", Fecha)

}
func Conviertotexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto
}
