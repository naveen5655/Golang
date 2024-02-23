package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go Conferernce"
	conferenceName := "Go Conferernce"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("conferenceTickets is %T,remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickers and", remainingTickets)
	fmt.Printf("Welcome to %v booking system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Thank you for booking a ticker for the Go Confernece!")
	// fmt.Println(conferenceName)

	bookings := []string{}

	for len(bookings) < 50 && remainingTickets > 0 {

		var firstName string
		var lastName string
		var email string
		var userTickets int
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter user tickets")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// isValidCity := city == "Singapore" || city == "London"

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are:%v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else if userTickets == remainingTickets {
			break
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			fmt.Printf("user input data is invalid")
		}

		// bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole array: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Array type:%T\n", bookings)
		// fmt.Printf("Array length:%v\n", len(bookings))

		// userName = "Tom"
		// userTickets = 2

		city := "London"

		switch city {
		case "New York":
			// execute code for booking new york confernece tickets
		case "Singapore":
			// execute code for booking singapore confernece ticket
		case "London":
			// execute code for booking London confernece tickets
		default:
			fmt.Print("No valid city selected")
		}
	}

}
