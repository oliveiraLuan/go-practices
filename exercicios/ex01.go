package exercicios

import "fmt"

func main() {
	var altura, base, area float32

	fmt.Print("Digite o valor da base do triângulo:")
	fmt.Scan(&base)

	fmt.Print("Digite o valor da altura do triângulo:")
	fmt.Scan(&altura)

	area = (base * altura) / 2

	fmt.Printf("Área do Triângulo = %v", area)
}
