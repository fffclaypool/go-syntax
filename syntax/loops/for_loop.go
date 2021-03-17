/* for Loop
  A for loop is a repetition control structure. It allows you to write a loop
that needs to execute a specific number of times.
*/

package main

import "fmt"

func forloop() {
	forloopBasicOperation()
}

func forloopBasicOperation() (int, int) {
	// Pattern A
	for n := 0; n < 10; n++ {
		fmt.Printf("value of n: %d\n", n)
	}

	var a int
	var b int = 15
	numbers := [6]int{1, 2, 3, 5}

	// Pattern B
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}

	// Pattern C
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}

	return a, b
}
