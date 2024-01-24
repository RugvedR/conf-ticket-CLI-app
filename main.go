package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference" //you can also do it like this below

	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("The following are the types of these variables: conferenceName: %T, conferenceTicket: %T\n", conferenceName, conferenceTickets)

	fmt.Println("Welcome to " + conferenceName + " Ticket booking Application")
	fmt.Println("We have total of ", remainingTickets, " tickets remaining!")
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your firstname: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The slice type: %T\n", bookings)
	fmt.Printf("length of slice: %v\n", len(bookings))

	fmt.Printf("User %v %v has booked %v tickets.\nYou will recieve a notification on your email: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("#### %v tickets remaing for %v", remainingTickets, conferenceName)

}
