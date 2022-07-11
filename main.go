package main

import "fmt"

func main() {
	// Type inference when you initialize a variable
	conferenceName := "Go conference"
	const conferenceTickets = 50   // can't define constants with syntactic sugar
	var remainingTickets uint = 50 // protects us from mistakenly set it to negative value

	//  fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to '%v' booking application\n", conferenceName)
	fmt.Printf("We have a total of %v of which %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	// ask user for their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
