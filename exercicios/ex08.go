package main

import "fmt"

func main() {
	var number int
	fmt.Println("Digite um número positivo e maior que 1: ")
	fmt.Scan(&number)
	fmt.Print("O resultado é: ", fat(number))
}

func fat(number int) int {
	if number == 0 {
		return 1
	}
	return number * fat(number-1)
}
