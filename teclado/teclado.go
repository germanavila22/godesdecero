package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func Ingresanumero() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ingresa primer numero")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Numero invalido" + err.Error())
		}
	}

	fmt.Println("ingresa primer numero2")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Numero invalido" + err.Error())
		}
	}

	fmt.Println("leyenda")
	if scanner.Scan() {
		leyenda = scanner.Text()
		if err != nil {
			panic("Numero invalido" + err.Error())
		}
	}
	fmt.Println(leyenda, numero1*numero2)
}
