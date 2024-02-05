package main

import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50
	bookings := []string{}

	fmt.Printf("Conference name is: %T, Remaining tickets is %T\n", conferenceName, conferenceTicket)

	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have %d conference ticket and %d remaining tickets!\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	for {
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
		
		remainingTicket = remainingTicket - userTicket
		
		//var bookings [50]string
		//bookings[0] = firstName + " " + lastName //array
		bookings = append(bookings, firstName+" "+lastName)
		
		fmt.Printf("%v tickets are remaining for %v\n", remainingTicket, conferenceName)
		fmt.Printf("Thank you %v %v for booking %v tickets! You will recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
		
		fmt.Printf("These are all our bookings: %v\n", bookings)
	}

	
}
