/* Constants
  Constants refer to fixed values that the program may not alter during its execution.
These fixed values are also called literals. Constants can be of any of the basic data
types like an integer constant, a floating constant, a character constant, or a string
literal. There are also enumeration constants as well. Constants are treated just like
regular variables except that their values cannot be modified after their definition.
*/

package main

import "fmt"

func constants() int {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area: %d\n", area)

	return area
}
