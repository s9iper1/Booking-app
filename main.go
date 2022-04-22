package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	// fmt.Printf("conferenceName type is %T", conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets from here to attend")

	bookings := []string{}

	for {

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

		if userTickets > uint(remainingTickets) {
			fmt.Printf("we only have %v tickects remaining, you cannot book %v tickets\n", remainingTickets, userTickets)
			break
		}

		remainingTickets = remainingTickets - int(userTickets)
		// bookings[] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets)

		// userName = "Bilal"
		// userTickets = 2

		// fmt.Printf("The whole array %v\n", bookings)
		// fmt.Printf("The First value is %v\n", bookings[1])
		// fmt.Printf("The array type %T\n", bookings)
		// fmt.Printf("Array length %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n",
			firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			// var firstname = names[0]
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("First names for bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out, come back again next year")
			break
		}
	}

}
