/* function closure
  Go programming language supports anonymous functions which can acts as function
closures. Anonymous functions are used when we want to define a function inline
without passing any name to it. In our example, we created a function getSequence()
which returns another function. The purpose of this function is to close over a
variable i of upper function to form a closure.
*/

package main

import (
	"fmt"
)

func getSequence() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}
}

func functionclosure() {
	basicFunctionclosure()
}

func basicFunctionclosure() {
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()

	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
