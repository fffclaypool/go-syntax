/* goto Statement
  A goto statement in Go programming language provides an unconditional jump from
the goto to a labeled statement in the same function.

  Note − Use of goto statement is highly discouraged in any programming language
because it becomes difficult to trace the control flow of a program, making the
program difficult to understand and hard to modify. Any program that uses a goto
can be rewritten using some other construct.

  The syntax for a goto statement in Go is as follows −

	goto label;
	..
	.
	label: statement;

  Here, label can be any plain text except Go keyword and it can be set anywhere
in the Go program above or below to goto statement.
*/

package main

import "fmt"

func gotostatement() {
	gotostatementBasicOperation()
}

func gotostatementBasicOperation() int {
	var a int = 10

LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
	return a
}
