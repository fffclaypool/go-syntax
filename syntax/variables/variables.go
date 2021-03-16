/* Variables
  A variable is nothing but a name given to a storage area that the programs can manipulate.
Each variable in Go has a specific type, which determines the size and layout of the
variable's memory, the range of values that can be stored within that memory, and the set
of operations that can be applied to the variable.

  The name of a variable can be composed of letters, digits, and the underscore character.
It must begin with either a letter or an underscore. Upper and lowercase letters are
distinct because Go is case-sensitive. Based on the basic types explained in the previous
chapter, there will be the following basic variable types −

	1.byte
	Typically a single octet(one byte). This is an byte type.

	2.int
	The most natural size of integer for the machine.

	3.float32
	A single-precision floating point value.

  A variable definition tells the compiler where and how much storage to create for the
variable. A variable definition specifies a data type and contains a list of one or more
variables of that type as follows −

	var variable_list optional_data_type;

  Here, optional_data_type is a valid Go data type including byte, int, float32, complex64,
boolean or any user-defined object, etc., and variable_list may consist of one or more
identifier names separated by commas. Some valid declarations are shown here −

    var i, j, k int;
    var c, ch byte;
    var f, salary float32;
    d = 42;

  The statement “var i, j, k;” declares and defines the variables i, j and k; which
instructs the compiler to create variables named i, j, and k of type int.

  Variables can be initialized (assigned an initial value) in their declaration. The type
of variable is automatically judged by the compiler based on the value passed to it. The
initializer consists of an equal sign followed by a constant expression as follows −

    variable_name = value;

  For example,

	d = 3, f = 5;    // declaration of d and f. Here d and f are int

  For definition without an initializer: variables with static storage duration are
implicitly initialized with nil (all bytes have the value 0); the initial value of all
other variables is zero value of their data type.
*/

package main

import "fmt"

func variables() {
	staticTypeDeclaration()
	dynamicTypeDeclaration()
	mixedVariableDeclaration()
}

func staticTypeDeclaration() float64 {
	/*
		  A static type variable declaration provides assurance to the compiler that there
		is one variable available with the given type and name so that the compiler can
		proceed for further compilation without requiring the complete detail of the variable.
		A variable declaration has its meaning at the time of compilation only, the compiler
		needs the actual variable declaration at the time of linking of the program.
	*/

	var x float64
	x = 20.0

	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	return x
}

func dynamicTypeDeclaration() (float64, int) {
	/*
		  A dynamic type variable declaration requires the compiler to interpret the type of
		the variable based on the value passed to it. The compiler does not require a variable
		to have type statically as a necessary requirement.

		  Try the following example, where the variables have been declared without any type.
		  Notice, in case of type inference, we initialized the variable y with := operator,
		whereas x is initialized using = operator.
	*/

	var x float64 = 20.0
	y := 42

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	return x, y
}

func mixedVariableDeclaration() (int, int, string) {
	// Variables of different types can be declared in one go using type inference.
	var a, b, c = 3, 4, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	return a, b, c
}
