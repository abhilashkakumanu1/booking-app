package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to '%v' booking application\n", conferenceName)
	fmt.Printf("We have a total of %v of which %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Go have oly one loop construct - for. This can be used for while, do while, for-each, etc
	//  Infinite loop
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// ask user for their name
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - uint(userTickets)
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You wil receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings: %v\n", bookings)
	}

}
