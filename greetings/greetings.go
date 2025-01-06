package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Ciao, %v. Benvenuto!", name)
	return message
}
