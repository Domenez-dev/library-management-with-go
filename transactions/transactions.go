package transactions

import (
	"fmt"
	"library-management/books"
	"library-management/users"
)

func ManageTransactions() {
	for {
		fmt.Println("\nManage Transactions")
		fmt.Println("1. Issue Book")
		fmt.Println("2. Return Book")
		fmt.Println("3. Back")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			IssueBook()
		case 2:
			ReturnBook()
		case 3:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func IssueBook() {
	var bookID, userID int
	fmt.Print("Enter Book ID to issue: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter User ID: ")
	fmt.Scanln(&userID)

	// Load Books and Users
	allBooks, _ := books.LoadBooks() // Renamed variable to avoid conflicts
	allUsers, _ := users.LoadUsers()

	// Validate User
	userFound := false
	for _, user := range allUsers {
		if user.ID == userID {
			userFound = true
			break
		}
	}
	if !userFound {
		fmt.Println("User not found.")
		return
	}

	// Validate and Issue Book
	for i, book := range allBooks {
		if book.ID == bookID {
			if book.IsIssued {
				fmt.Println("Book is already issued.")
				return
			}
			allBooks[i].IsIssued = true
			books.SaveBooks(allBooks) // Correct usage of SaveBooks
			fmt.Printf("Book ID %d issued to User ID %d successfully.\n", bookID, userID)
			return
		}
	}
	fmt.Println("Book not found.")
}

func ReturnBook() {
	var bookID int
	fmt.Print("Enter Book ID to return: ")
	fmt.Scanln(&bookID)

	allBooks, _ := books.LoadBooks() // Renamed variable to avoid conflicts

	// Validate and Return Book
	for i, book := range allBooks {
		if book.ID == bookID {
			if !book.IsIssued {
				fmt.Println("Book is not currently issued.")
				return
			}
			allBooks[i].IsIssued = false
			books.SaveBooks(allBooks) // Correct usage of SaveBooks
			fmt.Printf("Book ID %d returned successfully.\n", bookID)
			return
		}
	}
	fmt.Println("Book not found.")
}
