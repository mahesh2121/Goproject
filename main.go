package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var reminigtickets = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "Tickets and", reminigtickets, "are still available")
	fmt.Println("Get your tickets here to attend")

}