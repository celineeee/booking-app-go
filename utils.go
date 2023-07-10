package main

import "fmt"

// Capitlize the first character of the function makes it visible to all packages
func GreetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	// Or we could do this (only apply to variables, and you can't define its type)
	// conferenceName := "Go Conference"

	fmt.Printf("conferenceName is %T, remainingTickets is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "tickets remaining")
	fmt.Println("Get your tickets here to attend")

}
