package main

import (
	"fmt"
	"library-management/books"
	"library-management/transactions"
	"library-management/users"
	"os"
)

func main() {
	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Manage Books")
		fmt.Println("2. Manage Users")
		fmt.Println("3. Issue/Return Books")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			books.ManageBooks()
		case 2:
			users.ManageUsers()
		case 3:
			transactions.ManageTransactions()
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
