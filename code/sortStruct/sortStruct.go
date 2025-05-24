// Try sorting a slice of Person by Age.
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func insertionSort(n []Person) []Person {
	for i:=0; i<len(n); i++ {
		key:= n[i].Age
		j := i-1

		for j >= 0 && n[j].Age > key {
			n[j+1].Age = n[j].Age
			j--
		}

		n[j+1].Age = key
	}

	return n

}

func main(){
	customerBase := []Person{
		{Name: "Benjie", Age: 30},
		{Name: "Nyaigoti", Age: 23},
		{Name: "Mary", Age: 70},
		{Name: "Shantelle", Age: 8},
	}

	fmt.Println(insertionSort(customerBase))

}
