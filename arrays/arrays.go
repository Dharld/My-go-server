package main

import "fmt"

// Define a struct for a Person
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

type Describer interface {
	Display() 
}

// Implement the Describer interface
func (p Person) Display() {
    fmt.Printf("Name: %s\nAge: %d\n", p.FullName(), p.Age)
}
// Method to display person's full name
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}



func main() {
    // Create a new Person instance
    person1 := Person {
        FirstName: "John",
        LastName:  "Doe",
        Age: 30,
    }

	person2 := Person {
		FirstName: "Jane",
        LastName:  "Smith",
        Age: 25,
	}

	describers := []Describer{person1, person2}
    // Display the person's details

	for _, d := range describers {
        d.Display()
    }
}