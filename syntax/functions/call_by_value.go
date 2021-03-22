/* Call by value
  The call by value method of passing arguments to a function copies the actual
value of an argument into the formal parameter of the function. In this case,
changes made to the parameter inside the function have no effect on the argument.

  By default, Go programming language uses call by value method to pass arguments.
In general, this means that code within a function cannot alter the arguments
used to call the function.
*/

package main

import "fmt"

func callbyvalue() {
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a: %d\n", a)
	fmt.Printf("Before swap, value of b: %d\n", b)

	callbyvalueswap(a, b)

	fmt.Printf("After swap, value of a: %d\n", a)
	fmt.Printf("After swap, value of b: %d\n", b)
}

func callbyvalueswap(x, y int) int {
	var temp int

	temp = x
	x = y
	y = temp

	return temp
}
