package main

import "fmt"

// Person represents a person with a name and age.
type Person struct {
    Name string
    Age  int
}

// Greet prints a greeting from the person.
func (p Person) Greet() {
    fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    people := []Person{
        {Name: "Alice", Age: 30},
        {Name: "Bob", Age: 25},
        {Name: "Charlie", Age: 40},
    }
		var person Person 
	person.Age = 30
	person.Name = "Benjie"

	people = append(people, person)


    for _, person := range people {
        person.Greet()
    }


}
