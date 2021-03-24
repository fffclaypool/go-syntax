/* Error Handling
  Go programming provides a pretty simple error handling framework with inbuilt
error interface type of the following declaration −

	type error interface {
		Error() string
	}
*/

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
