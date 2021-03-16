/* The Select Statement
  The following rules apply to a select statement −

	  You can have any number of case statements within a select. Each case is followed
	by the value to be compared to and a colon.

      The type for a case must be the a communication channel operation.

	  When the channel operation occured the statements following that case will execute.
	No break is needed in the case statement.

	  A select statement can have an optional default case, which must appear at the end
	of the select. The default case can be used for performing a task when none of the
	cases is true. No break is needed in the default case.
*/

package main

import "fmt"

func main() {
	basicOperation()
}

func basicOperation() (int, int) {
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Printf("received %d from c1\n", i1)
	case c2 <- i2:
		fmt.Printf("sent %d to c2\n", i2)
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received %d from c3\n", i3)
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	return i1, i2
}
