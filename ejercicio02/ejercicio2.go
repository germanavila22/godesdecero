package ejercicio02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Damenumero() string {
	var numero1 int
	var err error
	var texto string
	fmt.Println("ingresa numero")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("numero incorrecto", err)
				continue
			} else {
				fmt.Println("vamos a imprimir")
				for i := 1; i <= 10; i++ {
					//fmt.Println(numero1, " x ", i, " = ", i*numero1)
					//fmt.Printf("%d x %d =%d \n", numero1, i, i*numero1)
					texto += fmt.Sprintf("%d x %d =%d \n", numero1, i, i*numero1) ///sprintf te devuelve una cadena de texto formateada
				}
				texto += fmt.Sprintf(" \n") ///sprintf te devuelve una cadena de texto formateada
				break

			}
		}
	}
	return texto
}
