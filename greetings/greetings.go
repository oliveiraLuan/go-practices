package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nome vuoto")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Ciao, %v. Benvenuto!",
		"È bello vederti %v",
		"Saluto, %v! Ben incontrato",
	}
	return formats[rand.Intn(len(formats))]
}
