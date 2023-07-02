package main

import (
	"fmt"
)

type Dimension struct {
	length int
	width  int
	height int
}

// func (d Dimension) Area() int {
// 	d.length * d.width
// }

// Multiple return values
func dimensions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = length * width * height

	return
}

func main() {
	area, volume := dimensions(10, 5, 5)
	fmt.Printf("Area : %d, Volume : %d", area, volume)
}
