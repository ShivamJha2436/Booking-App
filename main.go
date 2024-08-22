package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Declare WaitGroup
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	greetUsers()

	for {
		fmt.Println("Do you want to book a ticket or cancel an existing booking? (book/cancel):")
		var action string
		fmt.Scan(&action)

		if action == "book" {
			// Get user input for booking
			firstName, lastName, email, userTickets, err := getUserInput()
			if err != nil {
				fmt.Printf("Error: %v. Please try again.\n", err)
				continue
			}

			isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

			if isValidName && isValidEmail && isValidTicketNumber {
				bookTicket(userTickets, firstName, lastName, email)

				// Add to WaitGroup before starting the goroutine
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
					fmt.Println("First name or last name is too short.")
				}
				if !isValidEmail {
					fmt.Println("Email address is invalid.")
				}
				if !isValidTicketNumber {
					fmt.Println("Number of tickets is invalid.")
				}
			}
		} else if action == "cancel" {
			// Get user input for cancellation
			fmt.Println("Enter your email address to cancel the booking: ")
			var email string
			fmt.Scan(&email)

			fmt.Println("Enter number of tickets to cancel: ")
			var userTickets uint
			fmt.Scan(&userTickets)

			err := cancelBooking(email, userTickets)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Invalid action. Please choose 'book' or 'cancel'.")
		}
	}
	// Wait for all goroutines to complete
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint, error) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	// Simple validation
	if len(firstName) < 2 || len(lastName) < 2 {
		return "", "", "", 0, errors.New("first name or last name is too short")
	}
	if !strings.Contains(email, "@") {
		return "", "", "", 0, errors.New("email address is invalid")
	}
	if userTickets <= 0 || userTickets > remainingTickets {
		return "", "", "", 0, errors.New("number of tickets is invalid")
	}

	return firstName, lastName, email, userTickets, nil
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	mu.Lock()         // Lock before modifying shared resources
	defer mu.Unlock() // Unlock after modifying shared resources

	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	defer wg.Done() // Ensure Done is called when the goroutine exits
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
}

func cancelBooking(email string, userTickets uint) error {
	mu.Lock()         // Lock before modifying shared resources
	defer mu.Unlock() // Unlock after modifying shared resources

	for i, booking := range bookings {
		if booking.email == email && booking.numberOfTickets == userTickets {
			remainingTickets += userTickets
			bookings = append(bookings[:i], bookings[i+1:]...) // Remove the booking
			fmt.Printf("Booking for %v %v has been canceled. %v tickets have been refunded.\n", booking.firstName, booking.lastName, userTickets)
			return nil
		}
	}
	return fmt.Errorf("no matching booking found for the email: %v", email)
}
