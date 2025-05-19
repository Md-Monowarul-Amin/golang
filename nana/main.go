package main

import (
	"booking-app/helper"
	"fmt"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Print prints the given message to the console.
func main() {

	fmt.Printf("conferenceTickets is of type %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//Booking Infos
			bookTicket(remainingTickets, userTickets, conferenceName, firstName, lastName, email)
			// Call function Print First names
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("First name or last name you entered is too short. Please try again.\n")
			}
			if !isValidEmail {
				fmt.Printf("Email address you entered doesn't contain @ sign. Please try again.\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid. Please try again.\n")
			}
			fmt.Printf("Your input data is invalid. Please try again.\n")
			continue
		}
	}
	// var bookings = [50]string{"Nana", "Nicole", "Peter"}

	city := "London"

	switch city {
	case "London", "Berlin":
		fmt.Println("You are in London")
	case "New York":
		fmt.Println("You are in New York")
	case "Singapore", "Hong Kong":
		fmt.Println("You are in Paris")
	case "Mexico City":
		fmt.Println("You are in Tokyo")
	default:
		fmt.Println("No valid city selected")
	}
}

func greetUser() {
	fmt.Printf("Hello, welcome to the %v! booking application \n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	var firstNames = []string{}

	for index, booking := range bookings {
		firstName := booking.firstName
		firstNames = append(firstNames, firstName)
		fmt.Printf("First name of booking %v is %v\n", index+1, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
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

	fmt.Printf("User %v %v booked %v tickets fir now\n", firstName, lastName, userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(remainingTickets uint, userTickets int, conferenceName string, firstName string, lastName string, email string) {
	fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

	// Create a map for a user
	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(userTickets),
	}

	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, userData)
	// bookings[0] = firstName + " " + lastName

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice Type: %T\n", bookings)
	// // fmt.Print("Slice Type: %T\n", bookings[0:5])
	// fmt.Printf("Slice Length: %v\n", len(bookings))

	fmt.Printf("List of all bookings: %v\n", bookings)

	fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
