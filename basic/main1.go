package main

import (
	"fmt"
	"log"
	"strconv"
)

func printSum() {
	var aStr, bStr string
	fmt.Println("Enter two numbers to add")
	fmt.Scan(&aStr)
	fmt.Scan(&bStr)

	a, errA := strconv.ParseUint(aStr, 10, 64)
	b, errB := strconv.ParseUint(bStr, 10, 64)
	if errA != nil || errB != nil {
		fmt.Println("Invalid input. Please enter valid positive numbers.")
		return
	}
	fmt.Printf("a: %v, b: %v\n", a, b)

	sumValue, error := Sum(uint(a), uint(b))

	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("Sum of %v and %v is %v\n", a, b, sumValue)
}

func main() {
	fmt.Println("Hello, World!")
	// printSum()

	message, error := Hello("John")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("Formatted message: %v\n", message)

	names := []string{"John", "Jane", "Doe"}
	messages, error := Hellos(names)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Messages: %v\n", messages)
}
