package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference" //you can also do it like this below
var remainingTickets uint = 50
var bookings []string = []string{}

func main() {
	greetUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Println("Welcome to " + conferenceName + " Ticket booking Application")
	fmt.Println("We have total of ", remainingTickets, " tickets remaining!")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		var names = strings.Fields(booking) //splits the string with white space as separator, and returns a slice with split elements
		firstNames = append(firstNames, names[0])
	}
	// fmt.Printf("These are all our bookings: %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("User %v %v has booked %v tickets.\nYou will recieve a notification on your email: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("#### %v tickets remaing for %v", remainingTickets, conferenceName)

}
