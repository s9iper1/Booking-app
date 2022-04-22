package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	// fmt.Printf("conferenceName type is %T", conferenceName)
	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		//validation

		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTickets(userTickets, firstName, lastName, email)
			// print firstname
			fmt.Printf("First names for bookings are : %v\n", getFirstNames())

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, come back again next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("The email address is invalid")
			}
			if !isValidUserTickets {
				fmt.Println("Number of the tickets you entered is invalid")
			}
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets from here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		// var firstname = names[0]
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= uint(remainingTickets)

	return isValidName, isValidEmail, isValidUserTickets
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var userTickets uint
	var lastName string
	var email string
	// asking username from user

	fmt.Println("Enter your firstname")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastname")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName, lastName string, email string) {
	remainingTickets = remainingTickets - (userTickets)
	// bookings[] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
