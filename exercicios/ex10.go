package main

import "fmt"

func main() {
	var numeros [10]int
	var media, soma int

	fmt.Println("Digite 10 números: ")

	for i := 0; i < len(numeros); i++ {
		fmt.Println("Número: ", i+1)
		fmt.Scan(&numeros[i])
	}

	for i := 0; i < len(numeros); i++ {
		soma = soma + numeros[i]
	}

	media = soma / len(numeros)

	fmt.Println("A soma de todos os números é: ", soma)
	fmt.Println("A média de todos os números é: ", media)
}
