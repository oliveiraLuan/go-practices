package main

import "fmt"

func main() {
	a, b := 0, 1

	for a < 1000 {
		fmt.Println(a, "")
		b = a + b
		a = b - a
	}
}
