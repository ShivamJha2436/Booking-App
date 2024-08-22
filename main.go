package main

import "fmt"

func main() {
	confrenceName := "Go Confrence"
	const confrenceTickets int = 50
	var remainigTickets uint = 50
	var bookings = []string{}

	//fmt.Printf("confrenceTickets is %T, remainingTickets is %T, confrenceName is %T\n", confrenceTickets, remainigTickets, confrenceName)

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainigTickets)
	fmt.Println("Get Your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainigTickets = remainigTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Println("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n", lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainigTickets, confrenceName)

	fmt.Printf("These are all our bookings: %v\n", bookings)
}
