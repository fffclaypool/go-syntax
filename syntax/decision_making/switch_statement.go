/*
  A switch statement allows a variable to be tested for equality against a list of
values. Each value is called a case, and the variable being switched on is checked for
each switch case. In Go programming, switch statements are of two types −

	Expression Switch
	  In expression switch, a case contains expressions, which is compared
	against the value of the switch expression.

	Type Switch
	  In type switch, a case contain type which is compared against the type
	of a specially annotated switch expression.
*/

package main

import "fmt"

func switchstatement() {
	expressionSwitch()
	typeSwitch()
}

func expressionSwitch() string {
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("Excelent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}

	fmt.Printf("Your frade is %s\n", grade)

	return grade
}

func typeSwitch() interface{} {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x: %T\n", i)
	case int:
		fmt.Printf("x is int\n")
	case float64:
		fmt.Printf("x is float64\n")
	case func(int) float64:
		fmt.Printf("x is func(int)\n")
	case bool, string:
		fmt.Printf("x is bool or string\n")
	default:
		fmt.Printf("don't know the type\n")
	}
	return x
}
