/* function as values
  Go programming language provides the flexibility to create functions on the fly
and use them as values. In the following example, we've initialized a variable with
a function definition. Purpose of this function variable is just to use inbuilt
math.sqrt() function.
*/

package main

import "math"

func functionasvalues() {
	basicFunctionasvalues()
}

func basicFunctionasvalues() float64 {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	return getSquareRoot(9)
}
