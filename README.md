# Conference Ticket Booking Application

Welcome to the Conference Ticket Booking Application! This Go application allows users to book and cancel tickets for a conference. It includes features for booking tickets, sending confirmation emails, and managing bookings with basic error handling.

## Features

- **Book Tickets:** Users can book tickets for the conference by providing their name, email, and the number of tickets.
- **Cancel Tickets:** Users can cancel their bookings by providing their email and the number of tickets to cancel.
- **Send Confirmation Emails:** Once a ticket is booked, a confirmation email is sent (simulated with a delay in this application).
- **Track Bookings:** The application tracks all bookings and remaining tickets.

## Prerequisites

Ensure you have Go installed on your machine. You can download it from the [Go official website](https://golang.org/dl/).

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/conference-ticket-booking.git
    ```

2. Navigate to the project directory:

    ```bash
    cd conference-ticket-booking
    ```

3. (Optional) Create and switch to a new branch for development:

    ```bash
    git checkout -b feature/new-feature
    ```

## Usage

1. Run the application:

    ```bash
    go run main.go
    ```

2. Follow the prompts to book or cancel tickets:

    - **Book Tickets:** Enter your first name, last name, email, and the number of tickets.
    - **Cancel Tickets:** Enter the email associated with the booking and the number of tickets to cancel.

## Code Overview

- **`main.go`**: Contains the main logic for booking and canceling tickets, including user input handling, ticket booking, and sending confirmation emails.
- **`helper.go`**: Contains utility functions for validating user input.

## Error Handling

The application includes basic error handling for:
- Invalid names (less than 2 characters)
- Invalid email addresses (missing `@`)
- Invalid ticket numbers (not positive or exceeding available tickets)

## Contributing

If you'd like to contribute to this project, please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test them.
4. Submit a pull request with a clear description of your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

If you have any questions or feedback, feel free to reach out:

- **Email:** shivamkumar87148@gmail.com
- **GitHub:** [your-username](https://github.com/your-username)

---

Thank you for using the Conference Ticket Booking Application!
