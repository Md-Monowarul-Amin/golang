package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// CapitalizeFirstLetter capitalizes the first letter of a string to Export

// var Wg *sync.WaitGroup
var Wg = sync.WaitGroup{}

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets >= 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}

func SendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#########################")
	Wg.Done()
}
