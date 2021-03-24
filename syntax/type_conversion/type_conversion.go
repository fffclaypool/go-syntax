/* Type Conversion
  Type conversion is a way to convert a variable from one data type to another data
type. For example, if you want to store a long value into a simple integer then you
can type cast long to int. You can convert values from one type to another using the
cast operator. Its syntax is as follows −

	type_name(expression)
*/

package main

import "fmt"

func main() {
	var sum int = 17
	var count int = 15
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("Value of mean: %f\n", mean)
}
