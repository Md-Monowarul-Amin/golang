package main

import (
	"fmt"
	"strings"
)

// Print prints the given message to the console.
func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is of type %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// Ask user for their name
		fmt.Print("Enter your First name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your Last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets)

		remainingTickets = remainingTickets - uint(userTickets)
		bookings = append(bookings, firstName+" "+lastName)
		// bookings[0] = firstName + " " + lastName

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice Type: %T\n", bookings)
		// // fmt.Print("Slice Type: %T\n", bookings[0:5])
		// fmt.Printf("Slice Length: %v\n", len(bookings))

		fmt.Printf("These are all our bookings: %v\n", bookings)

		fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		var firstNames = []string{}

		for index, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
			fmt.Printf("First name of booking %v is %v\n", index+1, names[0])
		}
		fmt.Printf("First names of bookings are: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}

	}
	// var bookings = [50]string{"Nana", "Nicole", "Peter"}

}
