/* if...else statement
  An if statement can be followed by an optional else statement, which executes when
the boolean expression is false. If the boolean expression evaluates to true, then the
if block of code will be executed, otherwise else block of code is executed.
*/

package main

import "fmt"

func ifelsestatement() {
	basicOperationX()
	basicOperationY()
}

func basicOperationX() int {
	var a int = 100

	if a < 20 {
		fmt.Printf("a is less than 20\n")
	} else {
		fmt.Printf("a is not less than 20\n")
	}
	fmt.Printf("value of a is: %d\n", a)
	return a
}

func basicOperationY() int {
	var a int = 100

	if a == 10 {
		fmt.Printf("Value of a is 10\n")
	} else if a == 20 {
		fmt.Printf("Value of a is 20\n")
	} else if a == 30 {
		fmt.Printf("Value of a is 30\n")
	} else {
		fmt.Printf("None of the values is matching\n")
	}
	fmt.Printf("Excat value of a is: %d\n", a)
	return a
}
