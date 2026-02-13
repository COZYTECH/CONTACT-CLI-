package main

//Requirements:
// Store contacts (Name + Email) in a slice of structs
// Add a new contact
// List all contacts
// Search contact by name

// Challenge:

// Convert the map values to a slice, sort them alphabetically by name, then print.

// Hint:

// Use sort package

// Use sort.Slice

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
//save func to disk or contact list to file 

func saveContactsToFile(contacts map[string]*Contact, filename string) error {
	data, err := json.MarshalIndent(contacts, "", "  ") // pretty print
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644) // write to file

}

func loadContactsFromFile(filename string) (map[string]*Contact, error) {
	contacts := make(map[string]*Contact)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
        return contacts, nil // file doesn’t exist yet
    }

    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &contacts)
    if err != nil {
        return nil, err
    }

    return contacts, nil
}


func  main() {

// Create a slice to store contacts
//contacts := []Contact{}
//map to store contacts by name for quick search

	// Using map for faster lookup
contacts := make(map[string]*Contact)
// Load contacts from file if it exists
const filename = "C:\\Users\\HP\\Desktop\\conatctcli\\contactcli\\contacts.json"

contacts, err := loadContactsFromFile(filename)
if err != nil {
    fmt.Println("Error loading contacts:", err)
    return
}


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
				//fmt.Println("Contact added!")
				err := saveContactsToFile(contacts, filename)
				if err != nil {
					fmt.Println("Error saving contacts:", err)
				} else {
					fmt.Println("Contacts saved to", filename)
				}
			}
        case 2:
			//convert map to slice for sorting
			var contactSlice [] *Contact
			for _, c := range contacts {
				contactSlice = append(contactSlice, c)
			}
			//sort the slice by name
			sort.Slice(contactSlice, func(i, j int) bool {
				return strings.ToLower(contactSlice[i].Name) < strings.ToLower(contactSlice[j].Name)
			})
            fmt.Println("Contacts:")
            for _, c := range contactSlice {
                fmt.Println("-", c.Name, ":", c.Email)
				fmt.Println("Total contacts:", len(contactSlice))
            }
        case 3:
            var name string
            fmt.Print("Enter name to search: ")
            fmt.Scanln(&name)
			name = strings.ToLower(name)

            found := false
            // for _, c := range contacts {
            //     if c.Name == name {
            //         fmt.Println("Found:", c.Name, "->", c.Email)
            //         found = true
            //         break
            //     }
            // }
			name = strings.ToLower(name)
			if contact, exists := contacts[name]; exists {
				fmt.Println("Found:", contact.Name, "->", contact.Email)
			} else {
				fmt.Println("Contact not found.")
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
			 newName = strings.ToLower(newName)
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