package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var usersFile = "users.json"

func LoadUsers() ([]User, error) {
	var users []User
	file, err := ioutil.ReadFile(usersFile)
	if err == nil {
		err = json.Unmarshal(file, &users)
	}
	return users, err
}

func SaveUsers(users []User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err == nil {
		err = ioutil.WriteFile(usersFile, data, 0644)
	}
	return err
}

func ManageUsers() {
	for {
		fmt.Println("\nManage Users")
		fmt.Println("1. Add User")
		fmt.Println("2. View Users")
		fmt.Println("3. Delete User")
		fmt.Println("4. Back")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			AddUser()
		case 2:
			ViewUsers()
		case 3:
			DeleteUser()
		case 4:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func AddUser() {
	var user User
	fmt.Print("Enter User ID: ")
	fmt.Scanln(&user.ID)
	fmt.Print("Enter User Name: ")
	fmt.Scanln(&user.Name)

	users, _ := LoadUsers()
	users = append(users, user)
	if err := SaveUsers(users); err == nil {
		fmt.Println("User added successfully!")
	} else {
		fmt.Println("Failed to add user:", err)
	}
}

func ViewUsers() {
	users, err := LoadUsers()
	if err != nil {
		fmt.Println("Failed to load users:", err)
		return
	}

	fmt.Println("\nUsers List:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}
}

func DeleteUser() {
	var id int
	fmt.Print("Enter User ID to delete: ")
	fmt.Scanln(&id)

	users, _ := LoadUsers()
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			SaveUsers(users)
			fmt.Println("User deleted successfully!")
			return
		}
	}
	fmt.Println("User not found.")
}
