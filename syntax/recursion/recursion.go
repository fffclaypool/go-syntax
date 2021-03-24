/* Recursion
  Recursion is the process of repeating items in a self-similar way. The same
concept applies in programming languages as well. If a program allows to call
a function inside the same function, then it is called a recursive function
call.
*/

package main

import "fmt"

func main() {
	calculatingFactorial()
	fibonacciSeries()
}

func calculatingFactorial() {
	var i int = 15
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func fibonacciSeries() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
	fmt.Println()
}

func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
