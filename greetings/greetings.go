package greetings

import "fmt"

func Hello(name string) string{
	message := Sprintf("Ciao, %v. Benvenuto!", name)
	return message
}