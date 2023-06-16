package main

import (
	"fmt"
	"strings"
)

type User struct {
	firstname string
	lastname  string
	email     string
	tickets   uint
}

var confName = "GoConf"

const confTickets = 50

var remainingTickets uint = confTickets
var bookings = make([]User, 0)

func getUserInput(firstName, lastName, email *string) bool {
	fmt.Printf("Enter first name: ")
	fmt.Scan(firstName)

	fmt.Printf("Enter last name: ")
	fmt.Scan(lastName)
	if len(*firstName) < 2 || len(*lastName) < 2 {
		fmt.Println("Names most have more than 2 characters")
		return true
	}

	fmt.Printf("Enter email: ")
	fmt.Scan(email)
	if !strings.Contains(*email, "@") {
		fmt.Println("Email in invalid")
		return true
	}

	return false
}

func bookTicket(user User) {
	remainingTickets -= user.tickets
	bookings = append(bookings, user)
	// firstNames := []string{}

	// equivalent of forEach loop from JS. range operator returns a key/value pair
	// for _, u := range bookings {
	// 	firstNames = append(firstNames, u.firstname)
	// }
	fmt.Printf("Your just bought %v tickets %v\n", user.tickets, user.firstname)
	fmt.Println(bookings)
	fmt.Printf("Confirmation will be sent to your email at: %v\n", user.email)
}

func main() {
	fmt.Printf("Welcome to %v Booking app!\n", confName)
	fmt.Printf("We have %v available at the moment\n", remainingTickets)
	fmt.Println("Please book a ticket to attend.")

	for remainingTickets > 0 {
		var firstName, lastName, email string
		var userTickets uint

		err := getUserInput(&firstName, &lastName, &email)
		if err {
			continue
		}

		fmt.Printf("How many tickets do you want? ")
		fmt.Scan(&userTickets)
		if userTickets < 1 || userTickets > remainingTickets {
			fmt.Printf("You can't buy more than %v tickets. Please try again\n", remainingTickets)
			continue
		}

		var userData = User{
			firstname: firstName,
			lastname:  lastName,
			email:     email,
			tickets:   userTickets,
		}

		bookTicket(userData)
	}
	fmt.Println("Our conference is fully booked. See you next year")
}
