package main

//Requirements:
// Store contacts (Name + Email) in a slice of structs
// Add a new contact
// List all contacts
// Search contact by name

import "fmt"

type Contact struct {
	Name string
	Email string
}

//method to greet the contact

func (c Contact) Greet() string {
	return fmt.Sprintf("Hello, %s! Your email is %s.", c.Name, c.Email)
}

func  main() {

// Create a slice to store contacts
contacts := []Contact{}

for {
        fmt.Println("\n1. Add Contact")
        fmt.Println("2. List Contacts")
        fmt.Println("3. Search Contact")
        fmt.Println("4. Exit")
        fmt.Print("Choose option: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            var name, email string
            fmt.Print("Enter name: ")
            fmt.Scanln(&name)
            fmt.Print("Enter email: ")
            fmt.Scanln(&email)

            contacts = append(contacts, Contact{Name: name, Email: email})
            fmt.Println("Contact added!")
        case 2:
            fmt.Println("Contacts:")
            for _, c := range contacts {
                fmt.Println("-", c.Name, ":", c.Email)
            }
        case 3:
            var name string
            fmt.Print("Enter name to search: ")
            fmt.Scanln(&name)

            found := false
            for _, c := range contacts {
                if c.Name == name {
                    fmt.Println("Found:", c.Name, "->", c.Email)
                    found = true
                    break
                }
            }
            if !found {
                fmt.Println("Contact not found.")
            }
        case 4:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid option.")
        }
    }

}