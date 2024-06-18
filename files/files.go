package files

import (
	//	"bufio"
	"bufio"
	"fmt"

	//"go/scanner"

	//"io/ioutil"

	//"go/printer"
	//	"ioutil"
	"os"

	"github.com/godesdecero/godesdecero/ejercicio02"
)

var filename string = "./files/txt/tabla.txt"

func Grabatabla() {
	var texto string = ejercicio02.Damenumero()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("error al crear archivoi" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func Sumatabla() {
	var texto string = ejercicio02.Damenumero()
	if !Append(filename, texto) {
		fmt.Println("error al concatenar")
	}
}
func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error durante el append")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("erro al escribir")
		return false
	}
	arch.Close()
	return true
}

func Leoarchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("error al leer archivo")
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
