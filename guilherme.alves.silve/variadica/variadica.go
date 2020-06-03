package main

import "fmt"

func main() {
	fmt.Println(soma(1))
	fmt.Println(soma(1, 2))
	fmt.Println(soma(1, 2, 3))
	fmt.Println(soma(1, 2, 3, 4))
}

func soma(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}

	return resultado
}
