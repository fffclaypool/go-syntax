/* Method
  Go programming language supports special types of functions called methods. In
method declaration syntax, a "receiver" is present to represent the container of
the function. This receiver can be used to call a function using "." operator.
*/

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func method() {
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())
}
