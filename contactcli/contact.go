package main

//Requirements:
// Store contacts (Name + Email) in a slice of structs
// Add a new contact
// List all contacts
// Search contact by name

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name string
	Email string
}

//pointer
func (c *Contact) UpdateEmail(newEmail string) {
	c.Email = newEmail
}

func (c *Contact) UpdateName(newName string) {
	c.Name = newName
}
//method to greet the contact

func (c Contact) Greet() string {
	return fmt.Sprintf("Hello, %s! Your email is %s.", c.Name, c.Email)
}

func  main() {

// Create a slice to store contacts
//contacts := []Contact{}
//map to store contacts by name for quick search

	// Using map for faster lookup
contacts := make(map[string]*Contact)


for {
        fmt.Println("\n1. Add Contact")
        fmt.Println("2. List Contacts")
        fmt.Println("3. Search Contact")
        fmt.Println("4. Update Email")
        fmt.Println("5. Delete Contact")
        fmt.Println("6. Update Name")
        fmt.Println("7. Exit")
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
			name = strings.ToLower(name)

            //contacts = append(contacts, Contact{Name: name, Email: email})
			//contacts[name] = &Contact{Name: name, Email: email}
            //fmt.Println("Contact added!")
			if _, exists := contacts[name]; exists {
				fmt.Println("Contact already exists.")
			} else {
				contacts[name] = &Contact{Name: name, Email: email}
				fmt.Println("Contact added!")
			}
        case 2:
            fmt.Println("Contacts:")
            for _, c := range contacts {
                fmt.Println("-", c.Name, ":", c.Email)
				fmt.Println("Total contacts:", len(contacts))
            }
        case 3:
            var name string
            fmt.Print("Enter name to search: ")
            fmt.Scanln(&name)
			name = strings.ToLower(name)

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
		case 4: // Update Email
			var name, newEmail string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)

			if contact, exists := contacts[name]; exists {
				fmt.Print("Enter new email: ")
				fmt.Scanln(&newEmail)
				contact.UpdateEmail(newEmail)
				fmt.Println("Email updated!")
			} else {
				fmt.Println("Contact not found.")
			}
		case 5: // Delete Contact
			var name string
			var confirm string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)

			if _, exists := contacts[name]; exists {
				fmt.Print("Confirm deletion (yes/no): ")
				fmt.Scanln(&confirm)
				if confirm == "yes" {
					delete(contacts, name)
					fmt.Println("Contact deleted.")
				} else {
					fmt.Println("Deletion cancelled.")
				}
			} else {
				fmt.Println("Contact not found.")
			}

		case 6: // update name
		 var oldName, newName string
		 fmt.Print("Enter current name: ")
		 fmt.Scanln(&oldName)
		 oldName = strings.ToLower(oldName)
		 fmt.Print("Enter new name: ")
		 fmt.Scanln(&newName)

		 if contact, exists := contacts[oldName]; exists {
			 delete(contacts, oldName)
			 contacts[newName] = contact
			 fmt.Println("Contact name updated!")
		 } else {
			 fmt.Println("Contact not found.")
		 }

        case 7:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid option.")
        }
    }

}