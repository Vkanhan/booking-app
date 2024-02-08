package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50
	bookings := []string{}

	// Call greetings function
	greetUsers(conferenceName, conferenceTicket, remainingTicket)

	for {
		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remainingTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTicket = remainingTicket - userTicket
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("%v tickets are remaining for %v\n", remainingTicket, conferenceName)
			fmt.Printf("Thank you %v %v for booking %v tickets! You will recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)

			// Call printFunction
			firstNames := getfirstNames(bookings)
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

func greetUsers(confName string, confTicket int, remTicket uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have %d conference ticket and %d remaining tickets!\n", confTicket, remTicket)
	fmt.Println("Get your tickets here to attend")
}

func getfirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

	return isValidName, isValidEmail, isValidTicketNumber
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
