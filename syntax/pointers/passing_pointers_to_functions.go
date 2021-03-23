/* Passing pointers to functions
  Go programming language allows you to pass a pointer to a function. To do so, simply
declare the function parameter as a pointer type. In the following example, we pass two
pointers to a function and change the value inside the function which reflects back in
the calling function −
*/

package main

import "fmt"

func passingPointersToFunctions() (int, int) {
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)

	return a, b
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
