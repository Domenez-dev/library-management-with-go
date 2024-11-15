package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    IsIssued bool   `json:"is_issued"`
}

var booksFile = "books.json"

func LoadBooks() ([]Book, error) {
    var books []Book
    file, err := ioutil.ReadFile(booksFile)
    if err == nil {
        err = json.Unmarshal(file, &books)
    }
    return books, err
}

func SaveBooks(books []Book) error {
    data, err := json.MarshalIndent(books, "", "  ")
    if err == nil {
        err = ioutil.WriteFile(booksFile, data, 0644)
    }
    return err
}

func ManageBooks() {
    for {
        fmt.Println("\nManage Books")
        fmt.Println("1. Add Book")
        fmt.Println("2. View Books")
        fmt.Println("3. Delete Book")
        fmt.Println("4. Back")
        fmt.Print("Enter your choice: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            AddBook()
        case 2:
            ViewBooks()
        case 3:
            DeleteBook()
        case 4:
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func AddBook() {
    var book Book
    fmt.Print("Enter Book ID: ")
    fmt.Scanln(&book.ID)
    fmt.Print("Enter Book Title: ")
    fmt.Scanln(&book.Title)
    fmt.Print("Enter Book Author: ")
    fmt.Scanln(&book.Author)

    books, _ := LoadBooks()
    books = append(books, book)
    if err := SaveBooks(books); err == nil {
        fmt.Println("Book added successfully!")
    } else {
        fmt.Println("Failed to add book:", err)
    }
}

func ViewBooks() {
    books, err := LoadBooks()
    if err != nil {
        fmt.Println("Failed to load books:", err)
        return
    }

    fmt.Println("\nBooks List:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s, Issued: %v\n", book.ID, book.Title, book.Author, book.IsIssued)
    }
}

func DeleteBook() {
    var id int
    fmt.Print("Enter Book ID to delete: ")
    fmt.Scanln(&id)

    books, _ := LoadBooks()
    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            SaveBooks(books)
            fmt.Println("Book deleted successfully!")
            return
        }
    }
    fmt.Println("Book not found.")
}
