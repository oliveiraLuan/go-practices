package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var stringNumber, response string
	var erro error
	var intNumber int

	for response != "N" {
		fmt.Println("Digite um número positivo e maior que 1: ")
		fmt.Scan(&stringNumber)

		intNumber, erro = strconv.Atoi(stringNumber)

		for erro != nil || intNumber <= 0 {
			fmt.Println("Erro, digite um número inteiro e positivo.")
			fmt.Scan(&stringNumber)
			intNumber, erro = strconv.Atoi(stringNumber)
		}

		fmt.Println("O resultado é: ", fat(intNumber))

		fmt.Println("Desejar continuar? [S/N]:")
		fmt.Scan(&response)

		for response != "S" && response != "N" {
			fmt.Println("Caractere inválido, por favor escolha entre sim ou não com as respectivas letras [S] ou [N]:")
			fmt.Scan(&response)
			response = strings.ToUpper(response)
		}
	}
}

func fat(number int) int {
	if number == 0 {
		return 1
	}
	return number * fat(number-1)
}
