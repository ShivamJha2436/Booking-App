package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
var mu sync.Mutex // Mutex to ensure thread-safe access to shared resources

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets, err := getUserInput()
		if err != nil {
			fmt.Printf("Error: %v. Please try again.\n", err)
			continue
		}

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}
	wg.Wait()
}

// greetUsers prints a welcome message and ticket availability
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// getFirstNames returns a list of first names from the bookings
func getFirstNames() []string {
	mu.Lock()         // Lock the mutex before accessing shared resources
	defer mu.Unlock() // Unlock the mutex after accessing shared resources

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// getUserInput collects and validates user input with error handling
func getUserInput() (string, string, string, uint, error) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	if _, err := fmt.Scan(&firstName); err != nil || firstName == "" {
		return "", "", "", 0, fmt.Errorf("invalid input for first name: %v", err)
	}

	fmt.Println("Enter your last name: ")
	if _, err := fmt.Scan(&lastName); err != nil || lastName == "" {
		return "", "", "", 0, fmt.Errorf("invalid input for last name: %v", err)
	}

	fmt.Println("Enter your email address: ")
	if _, err := fmt.Scan(&email); err != nil || email == "" {
		return "", "", "", 0, fmt.Errorf("invalid input for email: %v", err)
	}

	fmt.Println("Enter number of tickets: ")
	if _, err := fmt.Scan(&userTickets); err != nil || userTickets == 0 {
		return "", "", "", 0, fmt.Errorf("invalid input for number of tickets: %v", err)
	}

	return firstName, lastName, email, userTickets, nil
}

// bookTicket updates the remaining tickets and records the booking
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	mu.Lock()         // Lock the mutex before modifying shared resources
	defer mu.Unlock() // Unlock the mutex after modifying shared resources

	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// sendTicket simulates sending a ticket via email
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // Simulate delay for sending ticket
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
