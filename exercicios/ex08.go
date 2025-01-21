package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stringNumber string
	var erro error
	var intNumber int
	fmt.Println("Digite um número positivo e maior que 1: ")
	fmt.Scan(&stringNumber)

	intNumber, erro = strconv.Atoi(stringNumber)

	for erro != nil || intNumber <= 0 {
		fmt.Println("Erro, digite um número inteiro e positivo.")
		fmt.Scan(&stringNumber)
		intNumber, erro = strconv.Atoi(stringNumber)
	}

	fmt.Print("O resultado é: ", fat(intNumber))
}

func fat(number int) int {
	if number == 0 {
		return 1
	}
	return number * fat(number-1)
}
