package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

// var bookings = [50]string{}

func main() {
	// conferenceName := "Go Conference"
	// const conferenceTickets int = 50
	// var remainingTickets uint = 50
	// // var bookings = [50]string{}
	// var bookings []string

	// fmt.Printf("conferenceTickets is %T, remremainingTickets is %T , conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUser()

	// fmt.Printf("Welcome to %v booking application\n", conferenceName)

	// for remainingTickets > 0 && len(bookings) < 50 {
	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// fmt.Println(&remainingTickets)

		// userName = "Tom"

		// isValidCity := city == "Singapore" || city == "London"
		// isInValidCity := city != "Singapore" && city != "London"

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// bookings[0] = firstName + " " + lastName

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("slice type: %T\n", bookings)
			// fmt.Printf("slice size: %v\n", len(bookings))

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// var noTickerRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				//	end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("slice type: %T\n", bookings)
			// fmt.Printf("slice size: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation on email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// var noTickerRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				//	end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
			// }else if userTickets == remainingTickets {
			//do something else

		} else {
			if !isValidName {
				fmt.Println("first name or last name your entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

			// fmt.Println("Your input data is invalid.")
			// break
		}
	}

	// city := "London"
	// switch city {
	// case "New York":
	// 	//

	// case "Singapore", "Berlin":
	// case "London", "Berlin":
	// 	//
	// default:
	// 	fmt.Println("No valid value selected.")
	// }
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation on email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
