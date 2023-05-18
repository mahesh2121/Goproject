package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var reminigtickets = 50
	bookings := [50]string{}

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "Tickets and", reminigtickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email addresses: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		reminigtickets = reminigtickets - int(userTickets)
		bookings[0] = firstName + "" + lastName

		fmt.Printf("the whole arry %v\n", bookings)
		fmt.Printf("the first arry %v\n", bookings[0])
		fmt.Printf("the arry type %T\n", bookings)
		fmt.Printf("the whole arry %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", reminigtickets, conferenceName)
		fmt.Printf("These are all our bookings %v\n", bookings)

	}
}
