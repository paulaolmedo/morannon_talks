package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
	Ejemplo de código que implementa la interface de la librería Standard fmt.Stringer
	podemos ver la declaración de la interface en el siguiente link
	https://pkg.go.dev/fmt#Stringer
*/

// Declaramos un Book type que satisface la fmt.Stringer interface
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declaramos Count type que satisface fmt.Stringer interface.
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Declaramos la funcion WriteLog() que toma un objeto que satisface la fmt.Stringer interface
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	// Initialize Book object
	book := Book{"La comunidad del anillo", "J. R. R. Tolkien"}
	WriteLog(book)

	// Initialize a Count object
	count := Count(3)
	WriteLog(count)
}
