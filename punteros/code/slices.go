package main

import "fmt"

func main() {
	slice := []string{"Hi", "How", "are", "you", "?"}

	updateSlice(slice)

	fmt.Println(slice)

}

func updateSlice(s []string) {
	s[2] = "Bye"
}
