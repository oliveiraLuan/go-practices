package main

import "fmt"

func main() {
	var nota, resultado float32

	for i := 1; i < 5; i++ {
		fmt.Println("Digite uma nota referente ao bimestre ", i)
		fmt.Scan(&nota)
		resultado = resultado + nota
	}
	fmt.Println("A média dos bimestres foi de: ", calculaMedia(resultado))
}

func calculaMedia(resultado float32) float32 {
	return resultado / 4
}
