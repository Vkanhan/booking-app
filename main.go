package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go conference"

const conferenceTicket = 50

var remainingTicket uint = 50

var bookings = make([]map[string]string, 0)

func main() {

	// Call greetings function
	greetUsers()

	for {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTicket, firstName, lastName, email)

			// Call printFunction
			firstNames := getfirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("Our conference is booked out! Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email you entered doesnt have @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number you entered is not valid")
			}

		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %d conference ticket and %d remaining tickets!\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}

func getfirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstNames"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	// login page

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of ticketes: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func bookTickets(userTicket uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTicket

	// Create a map for the user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets! You will recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTicket, conferenceName)

}
