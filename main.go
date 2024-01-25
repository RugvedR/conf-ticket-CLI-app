package main

import (
	"fmt"
	"strings"
)

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
	for {

		// ask user for their name
		fmt.Println("Enter your firstname: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets you want: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("User %v %v has booked %v tickets.\nYou will recieve a notification on your email: %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("#### %v tickets remaing for %v", remainingTickets, conferenceName)

			var firstNames []string
			for _, booking := range bookings {
				var names = strings.Fields(booking) //splits the string with white space as separator, and returns a slice with split elements
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets available. Please enter a valid choice. [you entered %v tickets]\n", remainingTickets, userTickets)
		}

	}

}
