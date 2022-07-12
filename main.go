package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to '%v' booking application\n", conferenceName)
	fmt.Printf("We have a total of %v of which %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Arrays in Go have fixed size, same data type
	var bookings [50]string

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
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole bookings array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You wil receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
