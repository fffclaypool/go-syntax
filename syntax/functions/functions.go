/* Functions
  A function is a group of statements that together perform a task. Every Go
program has at least one function, which is main(). You can divide your code
into separate functions. How you divide your code among different functions is
up to you, but logically, the division should be such that each function
performs a specific task.

  The general form of a function definition in Go programming language is as
follows −

	func function_name( [parameter list] ) [return_types]
	{
	  body of the function
	}
*/

package main

import "fmt"

func functions() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("Max value is: %d\n", ret)

	x, y := functionsswap("Mahesh", "Kumar")
	fmt.Println(x, y)
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func functionsswap(x, y string) (string, string) {
	return y, x
}
