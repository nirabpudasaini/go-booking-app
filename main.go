package main

import "fmt"

func main() {
	var confrenceName = "Go confrence"
	const confrenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", confrenceName, "booking application")
	fmt.Println("Total tickets:", confrenceTickets, "and", "Remaining Tickets:", remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
