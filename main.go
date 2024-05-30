package main

import (
	"fmt"

	"github.com/godesdecero/godesdecero/variables"
)

func main() {
	//fmt.Println("hola mundo")
	variables.Muestroenteros()
	variables.Restovariables()
	estado, texto := variables.Conviertotexto(15500)
	fmt.Println("el estado", estado)
	fmt.Println("el numero a texto", texto)
}
