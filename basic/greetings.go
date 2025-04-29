package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())  // Uncomment this line to see the test fail

	return message, nil
}

func Sum(a uint, b uint) (uint, error) {
	fmt.Printf("a: %v, b: %v\n", a, b)
	if a <= 0 || b <= 0 {
		log.Fatal("Please enter a positive number")
		return 0, fmt.Errorf("Please enter a positive number")
	}

	return (a + b), nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
