package main

import "fmt"

func main() {
	var numeros [10]int
	fmt.Println("Digite 10 números: ")

	for i := 0; i < len(numeros); i++ {
		fmt.Println("Numero: ", i+1)
		fmt.Scan(&numeros[i])
	}
	fmt.Println("Lista normal: ", numeros)

	fmt.Println("Lista invertida: ")
	for i := len(numeros) - 1; i >= 0; i-- {
		fmt.Print(numeros[i])
	}
}
