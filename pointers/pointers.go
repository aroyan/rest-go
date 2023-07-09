package main

import "fmt"

type Dimension struct {
	length int
	width  int
	height int
}

// If you using pointers, the value will change for the entire app
func (d *Dimension) Area() int {
	d.height = 8
	return d.length * d.width
}

// When this is not, I guess
func (d Dimension) Volume() int {
	d.height = 6
	return d.length * d.width * d.height
}

func main() {
	x, y := 10, 5

	n := &x

	fmt.Println(*n) // Return 10

	*n = 20

	fmt.Println(*n) // Return 20
	fmt.Println(x)  // Return 20

	t := &y

	fmt.Println(*t) // Return 5

	*t = *t + 25

	fmt.Println(*t) // Return 30

	d := Dimension{10, 5, 6}

	fmt.Println(d.Area())
	fmt.Println(d)

	fmt.Println(d.Volume())
	fmt.Println(d)

}
