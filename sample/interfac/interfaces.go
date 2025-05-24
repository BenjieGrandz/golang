package main

import (
	"fmt"
	"math"
)

// shapes interface
type shapes interface{
	area() float64
	circumference() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods 
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumference() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return math.Pi * (c.radius * 2)
}

func printShapeInfo(s shapes){
	fmt.Println("Area is: ", s.area())
	fmt.Println("Circumference is: ", s.circumference())
}

func main() {
	shapes := []shapes{
		square{length: 15.2},
		circle{radius: 15.2},
		circle{radius: 15.2},
		square{length: 4.9},
	}

	for i, v := range shapes {
		printShapeInfo(v)
		fmt.Printf("---%d \n", i+1)
	}
}
