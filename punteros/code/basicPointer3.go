package main

import (
	"fmt"
)

type address struct {
	numero int
}

func main() {

	var v int = 19
	var p *int

	p = &v

	fmt.Printf("La variable v es: %d \n", v)

	fmt.Printf("La dirección de memoria de v es: %v \n", &v)

	fmt.Printf("El puntero p referencia a la dirección de memoria: %v \n", p)

	fmt.Printf("Al desrefenciar el puntero p obtengo el valor: %d \n", *p)

}
