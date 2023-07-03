package main

import (
	"fmt"
)

type Dimension struct {
	length int
	width  int
	height int
}

type Person struct {
	name    string
	address string
	age     int
	married bool
}

func (p Person) Introduce() string {
	ismarried := "Married"

	if !p.married {
		ismarried = "Not Married"
	}

	return fmt.Sprintf("Hello, my name is %s and I'm from %s. I'm %d years old and I'm %s.", p.name, p.address, p.age, ismarried)
}

// Custom type int
type MyInt int

// Receiver type is Dimension and here named "d", and Area() is name of the method
func (d Dimension) Area() int {
	return d.length * d.width
}

// Receiver type is Dimension and here named "d", and Volume() is name of the method
func (d Dimension) Volume() int {
	return d.length * d.width * d.height
}

// func dimension(length, width, height int) (area int, volume int) {
// 	area = length * width
// 	volume = length * width * height

// 	return
// }

// Using normal function it'll be like this
// area, volume := dimension(10, 5, 6)

// While using Structs and Receiver methods it's shown in main function below

func main() {
	d := Dimension{10, 5, 6}
	fmt.Printf("Area = %d, Volume %d\n", d.Area(), d.Volume())

	person1 := Person{name: "Bakti", address: "East Java", age: 20, married: false}
	person2 := Person{name: "John Doe", address: "West Java", age: 30, married: true}

	fmt.Println(person1.Introduce())
	fmt.Println(person2.Introduce())
}

// More https://www.bogotobogo.com/GoLang/GoLang_Structs.php
