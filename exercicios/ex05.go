package main

import "fmt"

func main() {
	var qtd int
	var serie int

	fmt.Println("Quantos valores serão exibidos?")
	fmt.Scan(&qtd)
	fmt.Println("Termos da série: ")

	for x := 1; x <= qtd; i++ {
		serie = x*x + 1
		fmt.Println(serie)
	}
}
