/* Arrays
  Go programming language provides a data structure called the array, which
can store a fixed-size sequential collection of elements of the same type.
An array is used to store a collection of data, but it is often more useful
to think of an array as a collection of variables of the same type.
*/

package main

func arrays() [10]int {
	var n [10]int
	var i int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	return n
}
