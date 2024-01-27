package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference" //you can also do it like this below
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

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
		firstNames = append(firstNames, booking["firstName"])
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

	//create a map for a user
	// var userData = map[string]string
	userData := make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("User %v %v has booked %v tickets.\nYou will recieve a notification on your email: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("#### %v tickets remaing for %v", remainingTickets, conferenceName)

}
