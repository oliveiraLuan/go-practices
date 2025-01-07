package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nome vuoto")
	}

	message := fmt.Sprintf("Ciao, %v. Benvenuto!", name)
	return message, nil
}
