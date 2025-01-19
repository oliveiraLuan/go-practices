package main

import "fmt"

func main() {
	var numero1, numero2 int
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&numero1)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&numero2)

	for numero1 > numero2 {
		fmt.Println("O número 1 deve ser maior que o número 2")
		fmt.Println("Digite o primeiro número:")
		fmt.Scan(&numero1)
		fmt.Println("Digite o segundo número:")
		fmt.Scan(&numero2)
	}
}
