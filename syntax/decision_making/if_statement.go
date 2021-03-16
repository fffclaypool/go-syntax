/* if statement
  An if statement consists of a boolean expression followed by one or more statements.
If the boolean expression evaluates to true, then the block of code inside the if statement
is executed. If boolean expression evaluates to false, then the first set of code after
the end of the if statement (after the closing curly brace) is executed.
*/

package main

import "fmt"

func ifstatement() int {
	var a int = 10

	if a < 20 {
		fmt.Printf("a is less than 20\n")
	}
	fmt.Printf("value of a is: %d\n", a)
	return a
}
