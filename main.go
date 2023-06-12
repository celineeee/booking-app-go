package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// Or we could do this (only apply to variables, and you can't define its type)
	// conferenceName := "Go Conference"

	fmt.Printf("conferenceName is %T, remainingTickets is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "tickets remaining")
	fmt.Println("Get your tickets here to attend")

	// arrays in go are fixed size
	var bookings [50]string

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

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("The length of the array: %v\n", len(bookings))

	fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining\n", remainingTickets)

}
