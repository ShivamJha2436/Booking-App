package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// conferenceTickets defines the total number of tickets available
const conferenceTickets int = 50

// Global variables used across the application
var conferenceName = "Go Conference"
var remainingTickets uint = 50     // Remaining tickets that can be booked
var bookings = make([]UserData, 0) // Slice to store booking information
var mutex = &sync.Mutex{}          // Mutex to ensure thread-safe access to shared resources

// UserData struct represents the information of a user booking tickets
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// wg is used to wait for goroutines to finish their execution before exiting the program
var wg = sync.WaitGroup{}

func main() {

	greetUsers() // Greet the user and provide initial information about the conference

	// Main loop to handle booking until tickets are sold out
	for {
		// Get user input for booking
		firstName, lastName, email, userTickets := getUserInput()

		// Validate the user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// If all validations pass, proceed with the booking
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email) // Book the tickets

			// Use a goroutine to send the ticket asynchronously
			wg.Add(1) // Increment the WaitGroup counter
			go sendTicket(userTickets, firstName, lastName, email)

			// Display the first names of all bookings
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Check if tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break // Exit the loop as no more tickets are available
			}
		} else {
			// Display appropriate error messages if validation fails
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain a valid format.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}

	wg.Wait() // Wait for all goroutines to finish before exiting
}

// greetUsers prints a welcome message and information about the conference
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// getFirstNames returns a slice of first names from the bookings
func getFirstNames() []string {
	firstNames := []string{}
	mutex.Lock() // Lock the mutex to ensure thread-safe access to the bookings slice
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	mutex.Unlock() // Unlock the mutex after accessing the bookings
	return firstNames
}

// getUserInput prompts the user for their details and returns the input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Get user details through standard input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	// Trim whitespace from the input to avoid errors
	return strings.TrimSpace(firstName), strings.TrimSpace(lastName), strings.TrimSpace(email), userTickets
}

// validateUserInput checks if the user's input is valid
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2                  // Validate that the names have at least 2 characters
	isValidEmail := strings.Contains(email, "@") && len(email) >= 5           // Validate that the email contains "@" and has at least 5 characters
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // Validate that the number of tickets is positive and does not exceed available tickets
	return isValidName, isValidEmail, isValidTicketNumber
}

// bookTicket processes the booking by reducing the number of available tickets and storing the booking information
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	mutex.Lock() // Lock the mutex to ensure thread-safe updates to shared resources
	remainingTickets -= userTickets

	// Create a new booking entry
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Add the new booking to the bookings slice
	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	mutex.Unlock() // Unlock the mutex after updating shared resources
}

// sendTicket simulates the process of sending a ticket via email
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // Simulate a delay in sending the email
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() // Decrement the WaitGroup counter to indicate that this goroutine is done
}

// cancelBooking allows a user to cancel their booking by providing their email address
func cancelBooking(email string) error {
	mutex.Lock()         // Lock the mutex to ensure thread-safe access to the bookings slice
	defer mutex.Unlock() // Ensure the mutex is unlocked even if an error occurs

	// Find the booking with the given email and remove it
	for i, booking := range bookings {
		if booking.email == email {
			remainingTickets += booking.numberOfTickets        // Restore the tickets to the available pool
			bookings = append(bookings[:i], bookings[i+1:]...) // Remove the booking from the slice
			fmt.Printf("Booking for %v %v has been canceled. %v tickets have been refunded.\n", booking.firstName, booking.lastName, booking.numberOfTickets)
			return nil // Return nil to indicate successful cancellation
		}
	}
	return errors.New("booking not found") // Return an error if no booking was found for the given email
}
