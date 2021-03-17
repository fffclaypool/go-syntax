/* continue Statement
  The continue statement in Go programming language works somewhat like a break statement.
Instead of forcing termination, a continue statement forces the next iteration of the loop
to take place, skipping any code in between.

  In case of the for loop, continue statement causes the conditional test and increment
portions of the loop to execute.
*/

package main

import "fmt"

func continuestatement() {
	continuestatementBasicOperation()
}

func continuestatementBasicOperation() int {
	var a int = 10

	for a < 20 {
		if a == 15 {
			a = a + 1
			continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
	return a
}
