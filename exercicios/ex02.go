package main

import (
	"fmt"
	"math"
)

func main() {
	var angulo float64

	fmt.Println("Digite o ângulo desejado: ")
	fmt.Scan(&angulo)

	fmt.Println("Cosseno de", angulo, "é  = ", math.Cos(angulo))
	fmt.Println("Seno de", angulo, "é  = ", math.Sin(angulo))
}
