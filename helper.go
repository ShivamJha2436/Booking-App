package Booking_App

import (
	"strings"
)

// validateUserInput checks if the user input is valid
func validateUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= main_go.remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
