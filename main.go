package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings = [50]string{}
var bookings = []string{} // Slice (array with dynamic size)

func main() {

	greetingMessage := Hello("John")
	fmt.Printf("Message for greetings.go: %v\n", greetingMessage)
	fmt.Printf("Sum of 2 and 3 is %v\n", Sum(2, 3))

	greetUsers()

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		firstName, lastName, email, userTickets = getUserInput(firstName, lastName, email, userTickets)

		isValidName, isValidEmail, isValidTickets := validUserInput(firstName, lastName, email, userTickets)
		if !isValidName {
			fmt.Println("Please enter a valid name")
			continue
		}
		if !isValidEmail {
			fmt.Println("Please enter a valid email")
			continue
		}

		if isValidTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining\n", remainingTickets)
			continue
		}
		bookTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("These are all our first names: %v\n", firstNames)
		fmt.Printf("These are all our bookings: %v\n", bookings)
		if remainingTickets == 0 {
			fmt.Println(("All tickets are sold out. Come next year"))
			break
		}
	}

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets are %v for %v\n", remainingTickets, conferenceName)

}

func getUserInput(firstName string, lastName string, email string, userTickets uint) (string, string, string, uint) {
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 0 && len(lastName) > 0
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)

		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}
