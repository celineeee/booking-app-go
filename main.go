package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Package Level Variables
// Defined at the top outside all functions

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	// arrays in go are fixed size
	// var bookings [50]string

	// Now to create a slice: an abstruction of an array (vairable length)
	// var bookings []string
	// bookings := []string{}

	// We changed the datatype of userData to maps
	// Need to create a slice of maps
	// We need to mention the initial size. Slices are dynamic
	var bookings = make([]map[string]string, 0)

	// Create a UserData struct
	// var userData = UserData {
	// 	firstName: firstName,
	// 	lastName: lastName,
	// 	email: email,
	// 	numberOfTickets: conferenceTickets
	// }

	// there's only for loops in go
	// inifinite loop: for{}
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		// & this is a pointer (or a special variable )
		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)

		fmt.Printf("Please enter your last name\n")
		fmt.Scan(&lastName)

		fmt.Printf("Please enter your email\n")
		fmt.Scan(&email)

		fmt.Printf("Please enter number of tickets")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			fmt.Printf("We only have %v tickets remaining. ", remainingTickets)
			continue
		} else if userTickets == remainingTickets {
			// do something
		} else {
			fmt.Printf("Your input is invalid, try again")
		}

		remainingTickets = remainingTickets - userTickets
		// create a map for a user, and we cannot mix datatypes
		var userData = make(map[string]string)

		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		// bookings[0] = firstName + " " + lastName

		// dynamic arrays with slice
		bookings = append(bookings, userData)

		fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTickets)
		fmt.Printf("%v tickets remaining\n", remainingTickets)

		// for each loop
		// we want only the first names
		firstNames := []string{}
		// Range iterates over elements for different data stuctures (so not only arrays and slices)
		// Blank identifier: use _ to ignore a varaible you don't want to use
		for _, booking := range bookings {
			// use the string.fields to split the string with white space as seperator
			// var names = strings.Fields(booking)
			// firstNames = append(firstNames, names[0])
			firstNames = append(firstNames, booking["firstName"])

		}

		fmt.Printf("The first names of the books are : %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked out. ")
			break
		}
		// This blocks until the WaitGroup counter is 0
		wg.Wait()
		// This keyword starts a new goroutine
		// It handles new thread initialization and deletion
		// The main thread does not wait for any goroutine.
		// To fix this, create a waitGroup
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// city := "London"

		// switch city {
		// case "New York":
		// 	// start
		// case "Singapore", "Berlin":
		// 	//
		// default:
		// 	fmt.Printf("No valid city specified")
		// }

	}

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// This only print it to the screen
	// fmt.Printf("%v tickets for %v %v", userTicket, firstName, lastName)
	// Use this to create a string instead
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket: %v to email address %v\n", ticket, email)
	fmt.Println("#########################")
	// Decrement the counter of the waitGroup
	wg.Done()
}
