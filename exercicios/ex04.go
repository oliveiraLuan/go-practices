package main

import "fmt"

func main() {
	var number1, number2 float32
	var operator string

	fmt.Println("Digite o número 1:")
	fmt.Scan(&number1)

	fmt.Println("Digite o número 2:")
	fmt.Scan(&number2)

	fmt.Println("Escolha um operador matemático: (-,+,*,/)")
	fmt.Scan(&operator)

	if isAnInvalidOperator(operator) {
		fmt.Println("Erro ao ler caractere, por favor digite um operador válido: (-,+,*,/)")
		fmt.Scan(&operator)
	}

	fmt.Println("Resultado da operação: ", calc(number1, number2, operator))
}

func calc(numero1 float32, numero2 float32, operador string) float32 {
	switch operador {
	case "*":
		return numero1 * numero2
	case "/":
		return numero1 / numero2
	case "+":
		return numero1 + numero2
	default:
		return numero1 - numero2
	}
}

func isAnInvalidOperator(operator string) bool {
	return operator != "+" && operator != "-" && operator != "*" && operator != "/"
}
