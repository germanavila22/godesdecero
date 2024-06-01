package main

import (
	"fmt"
	"runtime"

	varejerc "github.com/godesdecero/godesdecero/var_ejerc"
	"github.com/godesdecero/godesdecero/variables"
)

func main() {
	//fmt.Println("hola mundo")
	variables.Muestroenteros()
	variables.Restovariables()
	estado, texto := variables.Conviertotexto(15500)
	fmt.Println("el estado", estado)
	fmt.Println("el numero a texto", texto)
	if os := runtime.GOOS; os == "Linux," || os == "OS X." {
		fmt.Println("esto no es windows")
	} else {
		fmt.Println("esto  es windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux.":
		fmt.Println("esto es linux")
	case "darwin":
		fmt.Println("esto es darwin")
	default:
		fmt.Printf("%s \n", os)

	}
	numero, texto := varejerc.Mifuncion("99")
	fmt.Println(numero)
	fmt.Println(texto)

}
