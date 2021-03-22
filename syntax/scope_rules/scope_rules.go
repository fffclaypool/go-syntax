/* Scope Rules
  A scope in any programming is a region of the program where a defined variable can
exist and beyond that the variable cannot be accessed. There are three places where
variables can be declared in Go programming language −

	# Inside a function or a block (local variables)
	# Outside of all functions (global variables)
	# In the definition of function parameters (formal parameters)

  Let us find out what are local and global variables and what are formal parameters.
*/

package main

func scoperules() {
	localVariables()
	globalVariables()
	formalParameters()
}

func localVariables() int {
	var a, b, c int

	a = 10
	b = 20
	c = a + b

	return c
}

var g int
var h int = 20

func globalVariables() (int, int) {
	var a, b int

	a = 10
	b = 20
	g = a + b
	var h int = 10

	return g, h
}

func formalParameters() int {
	var a int = 10
	var b int = 20
	var c int = 0

	c = sum(a, b)
	return c
}

func sum(a, b int) int {
	return a + b
}
